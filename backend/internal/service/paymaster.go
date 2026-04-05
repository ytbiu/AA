package service

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"time"

	"aa-bsc-7702-demo/backend/internal/config"
	appcrypto "aa-bsc-7702-demo/backend/internal/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	gcrypto "github.com/ethereum/go-ethereum/crypto"
)

type QuoteRequest struct {
	Sender               common.Address
	GasToken             common.Address
	Nonce                *big.Int
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}

type SponsorRequest struct{ QuoteRequest }

type QuoteResponse struct {
	GasToken       common.Address
	MaxCost        *big.Int
	TokenAmount    *big.Int
	ValidUntil     uint64
	QuoteHash      common.Hash
	QuoteSignature []byte
}

type SponsorResponse struct {
	QuoteResponse
	PaymasterAndData []byte
}

type PaymasterService struct {
	cfg      *config.Config
	signer   *appcrypto.Signer
	quoteABI abi.Arguments
}

func NewPaymasterService(cfg *config.Config, signer *appcrypto.Signer) (*PaymasterService, error) {
	quoteABI, err := buildQuoteABI()
	if err != nil {
		return nil, err
	}
	return &PaymasterService{cfg: cfg, signer: signer, quoteABI: quoteABI}, nil
}

func (s *PaymasterService) Quote(req QuoteRequest) (*QuoteResponse, error) {
	tokenPerNative, ok := s.cfg.TokenPerNativeByToken[req.GasToken]
	if !ok {
		return nil, fmt.Errorf("unsupported gasToken: %s", req.GasToken.Hex())
	}

	maxCost := s.maxCost(req.CallGasLimit, req.VerificationGasLimit, req.PreVerificationGas, req.MaxFeePerGas)
	if maxCost.Sign() <= 0 {
		return nil, fmt.Errorf("invalid maxCost")
	}

	tokenAmount := s.nativeToToken(maxCost, tokenPerNative)
	tokenAmount = tokenAmount.Mul(tokenAmount, big.NewInt(int64(10_000+s.cfg.SettlementMarkupBps)))
	tokenAmount = tokenAmount.Div(tokenAmount, big.NewInt(10_000))
	if tokenAmount.Sign() == 0 {
		tokenAmount = big.NewInt(1)
	}

	validUntil := uint64(time.Now().Unix() + s.cfg.QuoteValidSeconds)
	quoteHash, err := s.quoteHash(req, tokenAmount, validUntil)
	if err != nil {
		return nil, err
	}

	quoteSig, err := s.signer.SignEthMessageHash(quoteHash.Bytes())
	if err != nil {
		return nil, fmt.Errorf("sign quote hash: %w", err)
	}

	return &QuoteResponse{
		GasToken:       req.GasToken,
		MaxCost:        maxCost,
		TokenAmount:    tokenAmount,
		ValidUntil:     validUntil,
		QuoteHash:      quoteHash,
		QuoteSignature: quoteSig,
	}, nil
}

func (s *PaymasterService) Sponsor(req SponsorRequest) (*SponsorResponse, error) {
	quote, err := s.Quote(req.QuoteRequest)
	if err != nil {
		return nil, err
	}

	paymasterAndData := buildPaymasterAndData(
		s.cfg.PaymasterAddress,
		quote.GasToken,
		quote.ValidUntil,
		quote.TokenAmount,
		quote.QuoteSignature,
	)

	return &SponsorResponse{
		QuoteResponse:    *quote,
		PaymasterAndData: paymasterAndData,
	}, nil
}

func (s *PaymasterService) quoteHash(req QuoteRequest, tokenAmount *big.Int, validUntil uint64) (common.Hash, error) {
	encoded, err := s.quoteABI.Pack(
		s.cfg.ChainID,
		s.cfg.PaymasterAddress,
		req.Sender,
		req.Nonce,
		gcrypto.Keccak256Hash(req.CallData),
		req.CallGasLimit,
		req.VerificationGasLimit,
		req.PreVerificationGas,
		req.MaxFeePerGas,
		req.MaxPriorityFeePerGas,
		req.GasToken,
		tokenAmount,
		new(big.Int).SetUint64(validUntil),
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("encode quote payload: %w", err)
	}
	return gcrypto.Keccak256Hash(encoded), nil
}

func (s *PaymasterService) maxCost(
	callGasLimit *big.Int,
	verificationGasLimit *big.Int,
	preVerificationGas *big.Int,
	maxFeePerGas *big.Int,
) *big.Int {
	// EntryPoint requiredPrefund 在使用 paymaster 时会按更高倍数计入 verificationGas。
	// 如果这里只按 1 倍估算，容易在 validatePaymasterUserOp 阶段出现 AA33（报价不足）。
	verificationCost := new(big.Int).Mul(verificationGasLimit, big.NewInt(3))
	totalGas := new(big.Int).Add(callGasLimit, verificationCost)
	totalGas.Add(totalGas, preVerificationGas)
	return new(big.Int).Mul(totalGas, maxFeePerGas)
}

func (s *PaymasterService) nativeToToken(nativeWei *big.Int, tokenPerNative *big.Int) *big.Int {
	n := new(big.Int).Mul(nativeWei, tokenPerNative)
	return n.Div(n, big.NewInt(1e18))
}

func buildQuoteABI() (abi.Arguments, error) {
	newType := func(t string) (abi.Type, error) {
		return abi.NewType(t, "", nil)
	}

	uint256T, err := newType("uint256")
	if err != nil {
		return nil, err
	}
	addressT, err := newType("address")
	if err != nil {
		return nil, err
	}
	bytes32T, err := newType("bytes32")
	if err != nil {
		return nil, err
	}
	uint48T, err := newType("uint48")
	if err != nil {
		return nil, err
	}

	return abi.Arguments{
		{Type: uint256T},
		{Type: addressT},
		{Type: addressT},
		{Type: uint256T},
		{Type: bytes32T},
		{Type: uint256T},
		{Type: uint256T},
		{Type: uint256T},
		{Type: uint256T},
		{Type: uint256T},
		{Type: addressT},
		{Type: uint256T},
		{Type: uint48T},
	}, nil
}

func buildPaymasterAndData(
	paymaster common.Address,
	gasToken common.Address,
	validUntil uint64,
	tokenAmount *big.Int,
	quoteSig []byte,
) []byte {
	out := make([]byte, 0, 20+20+6+32+65)
	out = append(out, paymaster.Bytes()...)
	out = append(out, gasToken.Bytes()...)

	validUntilBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(validUntilBytes, validUntil)
	out = append(out, validUntilBytes[2:]...)

	out = append(out, leftPad32(tokenAmount.Bytes())...)
	out = append(out, quoteSig...)
	return out
}

func leftPad32(b []byte) []byte {
	if len(b) >= 32 {
		return b[len(b)-32:]
	}
	out := make([]byte, 32)
	copy(out[32-len(b):], b)
	return out
}
