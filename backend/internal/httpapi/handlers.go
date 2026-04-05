package httpapi

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"aa-bsc-7702-demo/backend/internal/config"
	"aa-bsc-7702-demo/backend/internal/service"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	gcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
)

type Server struct {
	cfg       *config.Config
	paymaster *service.PaymasterService
}

func NewServer(cfg *config.Config, paymaster *service.PaymasterService) *Server {
	return &Server{cfg: cfg, paymaster: paymaster}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.handleHealthz)
	mux.HandleFunc("/api/v1/paymaster/quote", s.methodOnly(http.MethodPost, s.handleQuote))
	mux.HandleFunc("/api/v1/paymaster/sponsor", s.methodOnly(http.MethodPost, s.handleSponsor))
	mux.HandleFunc("/api/v1/faucet/mint", s.methodOnly(http.MethodPost, s.handleFaucetMint))
	mux.HandleFunc("/api/v1/userop/send", s.methodOnly(http.MethodPost, s.handleSendUserOp))
	mux.HandleFunc("/api/v1/userop/receipt", s.methodOnly(http.MethodGet, s.handleGetUserOpReceipt))
	mux.HandleFunc("/api/v1/7702/upgrade", s.methodOnly(http.MethodPost, s.handleUpgrade7702ByBackend))
	return withCORS(mux)
}

func (s *Server) handleHealthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"paymaster": s.cfg.PaymasterAddress.Hex(),
		"bundler":   s.cfg.BundlerRPCURL,
	})
}

type quotePayload struct {
	Sender               string `json:"sender"`
	GasToken             string `json:"gasToken"`
	Nonce                string `json:"nonce"`
	CallData             string `json:"callData"`
	CallGasLimit         string `json:"callGasLimit"`
	VerificationGasLimit string `json:"verificationGasLimit"`
	PreVerificationGas   string `json:"preVerificationGas"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
}

type sponsorPayload struct{ quotePayload }

func (s *Server) handleQuote(w http.ResponseWriter, r *http.Request) {
	var payload quotePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid json: %w", err))
		return
	}
	quoteReq, err := parseQuoteRequest(payload)
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}
	resp, err := s.paymaster.Quote(quoteReq)
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"gasToken":       resp.GasToken.Hex(),
		"maxCost":        resp.MaxCost.String(),
		"tokenAmount":    resp.TokenAmount.String(),
		"validUntil":     resp.ValidUntil,
		"quoteHash":      resp.QuoteHash.Hex(),
		"quoteSignature": "0x" + common.Bytes2Hex(resp.QuoteSignature),
	})
}

func (s *Server) handleSponsor(w http.ResponseWriter, r *http.Request) {
	var payload sponsorPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid json: %w", err))
		return
	}

	quoteReq, err := parseQuoteRequest(payload.quotePayload)
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	resp, err := s.paymaster.Sponsor(service.SponsorRequest{
		QuoteRequest: quoteReq,
	})
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"gasToken":         resp.GasToken.Hex(),
		"maxCost":          resp.MaxCost.String(),
		"tokenAmount":      resp.TokenAmount.String(),
		"validUntil":       resp.ValidUntil,
		"quoteHash":        resp.QuoteHash.Hex(),
		"quoteSignature":   "0x" + common.Bytes2Hex(resp.QuoteSignature),
		"paymasterAndData": "0x" + common.Bytes2Hex(resp.PaymasterAndData),
	})
}

func parseQuoteRequest(payload quotePayload) (service.QuoteRequest, error) {
	if !common.IsHexAddress(payload.Sender) {
		return service.QuoteRequest{}, fmt.Errorf("invalid sender")
	}
	if !common.IsHexAddress(payload.GasToken) {
		return service.QuoteRequest{}, fmt.Errorf("invalid gasToken")
	}

	nonce, err := parseBigInt(payload.Nonce, "nonce")
	if err != nil {
		return service.QuoteRequest{}, err
	}
	callData, err := parseHex(payload.CallData, "callData")
	if err != nil {
		return service.QuoteRequest{}, err
	}
	callGasLimit, err := parseBigInt(payload.CallGasLimit, "callGasLimit")
	if err != nil {
		return service.QuoteRequest{}, err
	}
	verificationGasLimit, err := parseBigInt(payload.VerificationGasLimit, "verificationGasLimit")
	if err != nil {
		return service.QuoteRequest{}, err
	}
	preVerificationGas, err := parseBigInt(payload.PreVerificationGas, "preVerificationGas")
	if err != nil {
		return service.QuoteRequest{}, err
	}
	maxFeePerGas, err := parseBigInt(payload.MaxFeePerGas, "maxFeePerGas")
	if err != nil {
		return service.QuoteRequest{}, err
	}
	maxPriorityFeePerGas, err := parseBigInt(payload.MaxPriorityFeePerGas, "maxPriorityFeePerGas")
	if err != nil {
		return service.QuoteRequest{}, err
	}
	return service.QuoteRequest{
		Sender:               common.HexToAddress(payload.Sender),
		GasToken:             common.HexToAddress(payload.GasToken),
		Nonce:                nonce,
		CallData:             callData,
		CallGasLimit:         callGasLimit,
		VerificationGasLimit: verificationGasLimit,
		PreVerificationGas:   preVerificationGas,
		MaxFeePerGas:         maxFeePerGas,
		MaxPriorityFeePerGas: maxPriorityFeePerGas,
	}, nil
}

type faucetMintPayload struct {
	Token  string `json:"token"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func (s *Server) handleFaucetMint(w http.ResponseWriter, r *http.Request) {
	var payload faucetMintPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid json: %w", err))
		return
	}

	if !common.IsHexAddress(payload.Token) {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid token"))
		return
	}
	if !common.IsHexAddress(payload.To) {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid to"))
		return
	}
	amount, err := parseBigInt(payload.Amount, "amount")
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}
	if amount.Sign() <= 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("amount must be > 0"))
		return
	}

	relayerKey, relayerAddr, err := parsePrivateKey(s.cfg.RelayerPrivateKey)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("invalid RELAYER_PRIVATE_KEY: %w", err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, s.cfg.RPCURL)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("connect rpc failed: %w", err))
		return
	}
	defer client.Close()

	tokenAddr := common.HexToAddress(payload.Token)
	toAddr := common.HexToAddress(payload.To)
	nonce, err := client.PendingNonceAt(ctx, relayerAddr)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("get relayer nonce failed: %w", err))
		return
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("suggest gas price failed: %w", err))
		return
	}

	faucetABI, err := abi.JSON(strings.NewReader(`[{"type":"function","name":"faucetMint","stateMutability":"nonpayable","inputs":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"outputs":[]}]`))
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("build faucet abi failed: %w", err))
		return
	}
	calldata, err := faucetABI.Pack("faucetMint", toAddr, amount)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("pack faucet calldata failed: %w", err))
		return
	}

	callMsg := ethereum.CallMsg{
		From:     relayerAddr,
		To:       &tokenAddr,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     calldata,
	}
	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("estimate gas failed: %w", err))
		return
	}
	gasLimit = (gasLimit * 120) / 100

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddr,
		Value:    big.NewInt(0),
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     calldata,
	})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(s.cfg.ChainID), relayerKey)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("sign tx failed: %w", err))
		return
	}
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("send tx failed: %w", err))
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"token":  tokenAddr.Hex(),
		"to":     toAddr.Hex(),
		"amount": amount.String(),
		"txHash": signedTx.Hash().Hex(),
	})
}

type sendUserOpPayload struct {
	EntryPoint    string         `json:"entryPoint"`
	UserOperation map[string]any `json:"userOperation"`
}

type upgrade7702Payload struct {
	Owner string `json:"owner"`
}

func (s *Server) handleUpgrade7702ByBackend(w http.ResponseWriter, r *http.Request) {
	var payload upgrade7702Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil && err != io.EOF {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid json: %w", err))
		return
	}

	ownerKey, ownerAddr, err := parsePrivateKey(s.cfg.RelayerPrivateKey)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("invalid RELAYER_PRIVATE_KEY: %w", err))
		return
	}

	requestedOwner := strings.TrimSpace(payload.Owner)
	if requestedOwner != "" {
		if !common.IsHexAddress(requestedOwner) {
			writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid owner"))
			return
		}
		if common.HexToAddress(requestedOwner) != ownerAddr {
			writeErr(w, http.StatusBadRequest, fmt.Errorf("owner mismatch: 当前后端私钥地址为 %s，请在 MetaMask 使用同一地址", ownerAddr.Hex()))
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, s.cfg.RPCURL)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("connect rpc failed: %w", err))
		return
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("read chain id failed: %w", err))
		return
	}
	if chainID.Cmp(s.cfg.ChainID) != 0 {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("chain id mismatch: rpc=%s env=%s", chainID.String(), s.cfg.ChainID.String()))
		return
	}

	txNonce, err := client.PendingNonceAt(ctx, ownerAddr)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("get owner nonce failed: %w", err))
		return
	}
	authNonce := txNonce + 1

	auth, err := types.SignSetCode(ownerKey, types.SetCodeAuthorization{
		ChainID: *uint256.MustFromBig(s.cfg.ChainID),
		Address: s.cfg.DelegateLogicAddress,
		Nonce:   authNonce,
	})
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("sign setcode authorization failed: %w", err))
		return
	}

	tipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil || tipCap.Sign() <= 0 {
		gasPrice, gpErr := client.SuggestGasPrice(ctx)
		if gpErr != nil {
			writeErr(w, http.StatusBadGateway, fmt.Errorf("suggest gas failed: %w", gpErr))
			return
		}
		tipCap = gasPrice
	}
	feeCap := new(big.Int).Mul(tipCap, big.NewInt(2))

	callMsg := ethereum.CallMsg{
		From:              ownerAddr,
		To:                &ownerAddr,
		GasTipCap:         tipCap,
		GasFeeCap:         feeCap,
		Value:             big.NewInt(0),
		Data:              nil,
		AuthorizationList: []types.SetCodeAuthorization{auth},
	}
	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("estimate 7702 gas failed: %w", err))
		return
	}
	gasLimit = (gasLimit * 120) / 100

	txData := &types.SetCodeTx{
		ChainID:   uint256.MustFromBig(s.cfg.ChainID),
		Nonce:     txNonce,
		GasTipCap: uint256.MustFromBig(tipCap),
		GasFeeCap: uint256.MustFromBig(feeCap),
		Gas:       gasLimit,
		To:        ownerAddr,
		Value:     uint256.NewInt(0),
		Data:      nil,
		AuthList:  []types.SetCodeAuthorization{auth},
	}

	signedTx, err := types.SignNewTx(ownerKey, types.LatestSignerForChainID(s.cfg.ChainID), txData)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("sign setcode tx failed: %w", err))
		return
	}
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("send setcode tx failed: %w", err))
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"owner":       ownerAddr.Hex(),
		"delegateTo":  s.cfg.DelegateLogicAddress.Hex(),
		"txHash":      signedTx.Hash().Hex(),
		"txNonce":     txNonce,
		"authNonce":   authNonce,
		"txType":      "setCode(7702)",
		"description": "后端私钥已发起 EIP-7702 升级交易",
	})
}

type entryPointUserOperation struct {
	Sender               common.Address `abi:"sender"`
	Nonce                *big.Int       `abi:"nonce"`
	InitCode             []byte         `abi:"initCode"`
	CallData             []byte         `abi:"callData"`
	CallGasLimit         *big.Int       `abi:"callGasLimit"`
	VerificationGasLimit *big.Int       `abi:"verificationGasLimit"`
	PreVerificationGas   *big.Int       `abi:"preVerificationGas"`
	MaxFeePerGas         *big.Int       `abi:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int       `abi:"maxPriorityFeePerGas"`
	PaymasterAndData     []byte         `abi:"paymasterAndData"`
	Signature            []byte         `abi:"signature"`
}

func (s *Server) handleSendUserOp(w http.ResponseWriter, r *http.Request) {
	var payload sendUserOpPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid json: %w", err))
		return
	}
	if !common.IsHexAddress(payload.EntryPoint) {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid entryPoint"))
		return
	}
	if err := s.validatePaymasterInUserOp(payload.UserOperation); err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	if s.cfg.BundlerRPCURL == "" {
		s.handleSendUserOpDirect(w, payload)
		return
	}

	rpcReq := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_sendUserOperation",
		"params":  []any{payload.UserOperation, payload.EntryPoint},
	}
	body, _ := json.Marshal(rpcReq)

	resp, err := http.Post(s.cfg.BundlerRPCURL, "application/json", bytes.NewReader(body))
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("call bundler rpc failed: %w", err))
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("read bundler rpc response failed: %w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBody)
}

func (s *Server) validatePaymasterInUserOp(userOp map[string]any) error {
	raw, ok := userOp["paymasterAndData"]
	if !ok {
		return fmt.Errorf("missing paymasterAndData")
	}
	pad, ok := raw.(string)
	if !ok {
		return fmt.Errorf("paymasterAndData must be string")
	}
	pad = strings.TrimPrefix(strings.TrimPrefix(strings.TrimSpace(pad), "0x"), "0X")
	if len(pad) < 40 {
		return fmt.Errorf("paymasterAndData too short")
	}
	detected := common.HexToAddress("0x" + pad[:40])
	if detected != s.cfg.PaymasterAddress {
		return fmt.Errorf("paymaster mismatch: payload=%s backend=%s", detected.Hex(), s.cfg.PaymasterAddress.Hex())
	}
	return nil
}

func (s *Server) handleSendUserOpDirect(w http.ResponseWriter, payload sendUserOpPayload) {
	userOp, err := parseUserOperation(payload.UserOperation)
	if err != nil {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("parse userOperation failed: %w", err))
		return
	}

	relayerKey, relayerAddr, err := parsePrivateKey(s.cfg.RelayerPrivateKey)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("invalid RELAYER_PRIVATE_KEY: %w", err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, s.cfg.RPCURL)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("connect rpc failed: %w", err))
		return
	}
	defer client.Close()

	entryPoint := common.HexToAddress(payload.EntryPoint)
	calldata, err := packHandleOpsCalldata([]entryPointUserOperation{userOp}, relayerAddr)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("pack handleOps calldata failed: %w", err))
		return
	}

	nonce, err := client.PendingNonceAt(ctx, relayerAddr)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("get relayer nonce failed: %w", err))
		return
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("suggest gas price failed: %w", err))
		return
	}

	callMsg := ethereum.CallMsg{
		From:     relayerAddr,
		To:       &entryPoint,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     calldata,
	}
	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("estimate gas failed: %w", err))
		return
	}
	gasLimit = (gasLimit * 120) / 100

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &entryPoint,
		Value:    big.NewInt(0),
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     calldata,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(s.cfg.ChainID), relayerKey)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, fmt.Errorf("sign tx failed: %w", err))
		return
	}

	if err := client.SendTransaction(ctx, signedTx); err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("send tx failed: %w", err))
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"result":  signedTx.Hash().Hex(),
	})
}

func (s *Server) handleGetUserOpReceipt(w http.ResponseWriter, r *http.Request) {
	userOpHash := strings.TrimSpace(r.URL.Query().Get("hash"))
	if userOpHash == "" {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("hash is required"))
		return
	}

	if s.cfg.BundlerRPCURL == "" {
		s.handleGetDirectReceipt(w, userOpHash)
		return
	}

	rpcReq := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_getUserOperationReceipt",
		"params":  []any{userOpHash},
	}
	body, _ := json.Marshal(rpcReq)

	resp, err := http.Post(s.cfg.BundlerRPCURL, "application/json", bytes.NewReader(body))
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("call bundler rpc failed: %w", err))
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("read bundler rpc response failed: %w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBody)
}

func (s *Server) handleGetDirectReceipt(w http.ResponseWriter, txHash string) {
	if !isHexHash(txHash) {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid tx hash"))
		return
	}

	raw, err := postRPC(s.cfg.RPCURL, "eth_getTransactionReceipt", []any{txHash})
	if err != nil {
		writeErr(w, http.StatusBadGateway, err)
		return
	}

	var rpcResp map[string]any
	if err := json.Unmarshal(raw, &rpcResp); err != nil {
		writeErr(w, http.StatusBadGateway, fmt.Errorf("decode rpc response failed: %w", err))
		return
	}
	if rpcResp["error"] != nil {
		writeJSON(w, http.StatusOK, rpcResp)
		return
	}

	if rpcResp["result"] == nil {
		writeJSON(w, http.StatusOK, map[string]any{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  nil,
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"result": map[string]any{
			"receipt": rpcResp["result"],
		},
	})
}

func (s *Server) methodOnly(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			writeErr(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
			return
		}
		next(w, r)
	}
}

func parseUserOperation(src map[string]any) (entryPointUserOperation, error) {
	getString := func(key string) (string, error) {
		v, ok := src[key]
		if !ok {
			return "", fmt.Errorf("missing %s", key)
		}
		s, ok := v.(string)
		if !ok {
			return "", fmt.Errorf("%s must be string", key)
		}
		s = strings.TrimSpace(s)
		if s == "" {
			return "", fmt.Errorf("%s is empty", key)
		}
		return s, nil
	}

	getAddress := func(key string) (common.Address, error) {
		raw, err := getString(key)
		if err != nil {
			return common.Address{}, err
		}
		if !common.IsHexAddress(raw) {
			return common.Address{}, fmt.Errorf("invalid %s", key)
		}
		return common.HexToAddress(raw), nil
	}

	getBig := func(key string) (*big.Int, error) {
		raw, err := getString(key)
		if err != nil {
			return nil, err
		}
		return parseBigInt(raw, key)
	}

	getBytes := func(key string) ([]byte, error) {
		raw, err := getString(key)
		if err != nil {
			return nil, err
		}
		return parseHex(raw, key)
	}

	sender, err := getAddress("sender")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	nonce, err := getBig("nonce")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	initCode, err := getBytes("initCode")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	callData, err := getBytes("callData")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	callGasLimit, err := getBig("callGasLimit")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	verificationGasLimit, err := getBig("verificationGasLimit")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	preVerificationGas, err := getBig("preVerificationGas")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	maxFeePerGas, err := getBig("maxFeePerGas")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	maxPriorityFeePerGas, err := getBig("maxPriorityFeePerGas")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	paymasterAndData, err := getBytes("paymasterAndData")
	if err != nil {
		return entryPointUserOperation{}, err
	}
	signature, err := getBytes("signature")
	if err != nil {
		return entryPointUserOperation{}, err
	}

	return entryPointUserOperation{
		Sender:               sender,
		Nonce:                nonce,
		InitCode:             initCode,
		CallData:             callData,
		CallGasLimit:         callGasLimit,
		VerificationGasLimit: verificationGasLimit,
		PreVerificationGas:   preVerificationGas,
		MaxFeePerGas:         maxFeePerGas,
		MaxPriorityFeePerGas: maxPriorityFeePerGas,
		PaymasterAndData:     paymasterAndData,
		Signature:            signature,
	}, nil
}

func parsePrivateKey(raw string) (*ecdsa.PrivateKey, common.Address, error) {
	pk, err := gcrypto.HexToECDSA(strings.TrimPrefix(strings.TrimSpace(raw), "0x"))
	if err != nil {
		return nil, common.Address{}, err
	}
	addr := gcrypto.PubkeyToAddress(pk.PublicKey)
	return pk, addr, nil
}

func packHandleOpsCalldata(ops []entryPointUserOperation, beneficiary common.Address) ([]byte, error) {
	const entryPointABI = `[
	  {
	    "name":"handleOps",
	    "type":"function",
	    "stateMutability":"nonpayable",
	    "inputs":[
	      {
	        "name":"ops",
	        "type":"tuple[]",
	        "components":[
	          {"name":"sender","type":"address"},
	          {"name":"nonce","type":"uint256"},
	          {"name":"initCode","type":"bytes"},
	          {"name":"callData","type":"bytes"},
	          {"name":"callGasLimit","type":"uint256"},
	          {"name":"verificationGasLimit","type":"uint256"},
	          {"name":"preVerificationGas","type":"uint256"},
	          {"name":"maxFeePerGas","type":"uint256"},
	          {"name":"maxPriorityFeePerGas","type":"uint256"},
	          {"name":"paymasterAndData","type":"bytes"},
	          {"name":"signature","type":"bytes"}
	        ]
	      },
	      {"name":"beneficiary","type":"address"}
	    ],
	    "outputs":[]
	  }
	]`

	parsed, err := abi.JSON(strings.NewReader(entryPointABI))
	if err != nil {
		return nil, err
	}
	return parsed.Pack("handleOps", ops, beneficiary)
}

func postRPC(url string, method string, params []any) ([]byte, error) {
	rpcReq := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}
	body, _ := json.Marshal(rpcReq)

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("call rpc failed: %w", err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rpc response failed: %w", err)
	}
	return raw, nil
}

func parseBigInt(raw, field string) (*big.Int, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, fmt.Errorf("%s is required", field)
	}
	base := 10
	if strings.HasPrefix(raw, "0x") || strings.HasPrefix(raw, "0X") {
		raw = raw[2:]
		base = 16
	}
	n := new(big.Int)
	if _, ok := n.SetString(raw, base); !ok {
		return nil, fmt.Errorf("invalid %s", field)
	}
	return n, nil
}

func parseHex(raw, field string) ([]byte, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, fmt.Errorf("%s is required", field)
	}
	if strings.HasPrefix(raw, "0x") || strings.HasPrefix(raw, "0X") {
		raw = raw[2:]
	}
	if len(raw)%2 != 0 {
		return nil, fmt.Errorf("invalid %s hex length", field)
	}
	if raw == "" {
		return []byte{}, nil
	}
	b := common.FromHex("0x" + raw)
	if len(b) == 0 {
		return nil, fmt.Errorf("invalid %s hex", field)
	}
	return b, nil
}

func isHexHash(raw string) bool {
	raw = strings.TrimSpace(raw)
	if strings.HasPrefix(raw, "0x") || strings.HasPrefix(raw, "0X") {
		raw = raw[2:]
	}
	if len(raw) != 64 {
		return false
	}
	for _, c := range raw {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]any{
		"error": err.Error(),
	})
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
