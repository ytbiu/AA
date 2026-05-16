package api

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gin-gonic/gin"
	"github.com/holiman/uint256"

	"aa-wallet-backend/internal/models"
	"aa-wallet-backend/internal/relayer"
)

func (h *Handlers) Authorize7702(c *gin.Context) {
	var req models.Authorize7702Request
	if err := c.ShouldBindJSON(&req); err != nil {
		respondBadRequest(c, err)
		return
	}

	log.Printf("[Authorize7702] Request: user=%s, chainId=%d, nonce=%d", req.UserAddress, req.ChainID, req.Nonce)

	userAddr := parseAddress(req.UserAddress)
	implAddr := h.account.GetAddress()

	signer, err := h.verify7702Signature(req.ChainID, implAddr, req.Nonce, req.V, req.R, req.S, userAddr)
	if err != nil {
		respondInternalError(c, "signature_verification_failed", err.Error())
		return
	}

	if signer != userAddr {
		respondInternalError(c, "invalid_signer", "signature signer does not match user address")
		return
	}

	relayer, err := h.relayerPool.SelectIdle()
	if err != nil {
		respondInternalError(c, "no_relayer", "no available relayer")
		return
	}

	r := new(big.Int).SetBytes(common.Hex2Bytes(strings.TrimPrefix(req.R, "0x")))
	s := new(big.Int).SetBytes(common.Hex2Bytes(strings.TrimPrefix(req.S, "0x")))

	tx, err := h.send7702SetCodeTransaction(relayer, implAddr, req.ChainID, req.Nonce, req.V, r, s)
	if err != nil {
		respondInternalError(c, "setcode_tx_failed", err.Error())
		return
	}

	receipt, err := h.waitForTx(tx, 30e9)
	if err != nil {
		c.JSON(http.StatusOK, models.Authorize7702Response{
			TxHash:        tx.Hash().Hex(),
			Status:        "pending",
			BoundContract: implAddr.Hex(),
		})
		return
	}

	h.relayerPool.MarkComplete(relayer.Address)

	status := "success"
	if receipt.Status == 0 {
		status = "failed"
	}

	c.JSON(http.StatusOK, models.Authorize7702Response{
		TxHash:        tx.Hash().Hex(),
		Status:        status,
		BoundContract: implAddr.Hex(),
	})
}

func (h *Handlers) Clear7702(c *gin.Context) {
	var req models.Clear7702Request
	if err := c.ShouldBindJSON(&req); err != nil {
		respondBadRequest(c, err)
		return
	}

	userAddr := parseAddress(req.UserAddress)
	emptyAddr := common.Address{}

	signer, err := h.verify7702Signature(req.ChainID, emptyAddr, req.Nonce, req.V, req.R, req.S, userAddr)
	if err != nil {
		respondInternalError(c, "signature_verification_failed", err.Error())
		return
	}

	if signer != userAddr {
		respondInternalError(c, "invalid_signer", "signature signer does not match user address")
		return
	}

	relayer, err := h.relayerPool.SelectIdle()
	if err != nil {
		respondInternalError(c, "no_relayer", "no available relayer")
		return
	}

	r := new(big.Int).SetBytes(common.Hex2Bytes(strings.TrimPrefix(req.R, "0x")))
	s := new(big.Int).SetBytes(common.Hex2Bytes(strings.TrimPrefix(req.S, "0x")))

	tx, err := h.send7702SetCodeTransaction(relayer, emptyAddr, req.ChainID, req.Nonce, req.V, r, s)
	if err != nil {
		respondInternalError(c, "setcode_tx_failed", err.Error())
		return
	}

	receipt, err := h.waitForTx(tx, 30e9)
	if err != nil {
		c.JSON(http.StatusOK, models.Clear7702Response{
			TxHash: tx.Hash().Hex(),
			Status: "pending",
		})
		return
	}

	h.relayerPool.MarkComplete(relayer.Address)

	status := "success"
	if receipt.Status == 0 {
		status = "failed"
	}

	c.JSON(http.StatusOK, models.Clear7702Response{
		TxHash: tx.Hash().Hex(),
		Status: status,
	})
}

func (h *Handlers) verify7702Signature(chainID uint64, impl common.Address, nonce uint64, v uint8, rHex, sHex string, expectedUser common.Address) (common.Address, error) {
	rlpData, err := rlp.EncodeToBytes([]interface{}{chainID, impl, nonce})
	if err != nil {
		return common.Address{}, fmt.Errorf("rlp encode failed: %v", err)
	}

	authData := append([]byte{0x05}, rlpData...)
	authHash := crypto.Keccak256Hash(authData)

	rBytes := common.Hex2Bytes(strings.TrimPrefix(rHex, "0x"))
	sBytes := common.Hex2Bytes(strings.TrimPrefix(sHex, "0x"))
	if len(rBytes) != 32 || len(sBytes) != 32 {
		return common.Address{}, fmt.Errorf("r and s must be 32 bytes")
	}

	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	return recoverSignerFrom7702Auth(authHash, v, r, s)
}

func recoverSignerFrom7702Auth(hash common.Hash, v uint8, r, s *big.Int) (common.Address, error) {
	yParity := v
	if yParity > 1 {
		yParity = yParity - 27
	}

	sig := make([]byte, 65)
	copy(sig[:32], r.Bytes())
	copy(sig[32:64], s.Bytes())
	sig[64] = yParity

	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to recover public key: %v", err)
	}
	return crypto.PubkeyToAddress(*pubKey), nil
}

func (h *Handlers) send7702SetCodeTransaction(
	relayer *relayer.Relayer,
	impl common.Address,
	chainID uint64,
	nonce uint64,
	v uint8,
	r, s *big.Int,
) (*types.Transaction, error) {
	relayerNonce, err := h.ethClient.GetEthClient().NonceAt(context.Background(), relayer.Address, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get relayer nonce: %v", err)
	}

	gasPrice, err := h.ethClient.SuggestGasPrice()
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %v", err)
	}

	yParity := v
	if yParity > 1 {
		yParity = yParity - 27
	}

	chainIDUint256 := uint256.NewInt(chainID)
	auth := types.SetCodeAuthorization{
		ChainID: *chainIDUint256,
		Address: impl,
		Nonce:   nonce,
		V:       yParity,
		R:       *uint256.MustFromBig(r),
		S:       *uint256.MustFromBig(s),
	}

	setCodeTx := &types.SetCodeTx{
		ChainID:    chainIDUint256,
		Nonce:      relayerNonce,
		GasTipCap:  uint256.MustFromBig(gasPrice),
		GasFeeCap:  uint256.MustFromBig(gasPrice),
		Gas:        100000,
		To:         relayer.Address, // 发送到 relayer 自己，避免用户变成 EIP-7702 后空调用失败
		Value:      new(uint256.Int),
		Data:       nil,
		AccessList: nil,
		AuthList:   []types.SetCodeAuthorization{auth},
	}

	signer := types.NewPragueSigner(new(big.Int).SetUint64(chainID))
	signedTx, err := types.SignTx(types.NewTx(setCodeTx), signer, relayer.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	err = h.ethClient.GetEthClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx, nil
}
