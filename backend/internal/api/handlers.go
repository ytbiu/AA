package api

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"

	"aa-wallet-backend/internal/contract"
	"aa-wallet-backend/internal/models"
	"aa-wallet-backend/internal/relayer"
	"aa-wallet-backend/pkg/eth"
)

var erc20TransferABI = `[{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"}]`
var erc20ApproveABI = `[{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"}]`

type Handlers struct {
	relayerPool *relayer.Pool
	paymaster   *contract.PaymasterContract
	usdt        *contract.USDTContract
	account     *contract.AccountContract
	ethClient   *eth.Client
}

func NewHandlers(pool *relayer.Pool, paymaster *contract.PaymasterContract, usdt *contract.USDTContract, account *contract.AccountContract, ethClient *eth.Client) *Handlers {
	return &Handlers{
		relayerPool: pool,
		paymaster:   paymaster,
		usdt:        usdt,
		account:     account,
		ethClient:   ethClient,
	}
}

func (h *Handlers) createTransactor(relayer *relayer.Relayer) (*bind.TransactOpts, error) {
	transactor, err := bind.NewKeyedTransactorWithChainID(relayer.PrivateKey, h.ethClient.GetChainID())
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	gasPrice, err := h.ethClient.SuggestGasPrice()
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %v", err)
	}
	transactor.GasPrice = gasPrice
	transactor.GasLimit = uint64(500000)
	transactor.Context = context.Background()
	return transactor, nil
}

func (h *Handlers) waitForTx(tx *types.Transaction, timeout time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, h.ethClient.GetEthClient(), tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx: %v", err)
	}
	return receipt, nil
}

func getTransferData(to common.Address, amount *big.Int) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc20TransferABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}
	return parsedABI.Pack("transfer", to, amount)
}

func getApproveData(spender common.Address, amount *big.Int) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc20ApproveABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}
	return parsedABI.Pack("approve", spender, amount)
}

func recoverSigner(hash common.Hash, signature []byte) (common.Address, error) {
	if len(signature) != 65 {
		return common.Address{}, fmt.Errorf("invalid signature length: %d", len(signature))
	}
	sig := make([]byte, 65)
	copy(sig, signature)
	if sig[64] >= 27 {
		sig[64] -= 27
	}
	ethSignedHash := crypto.Keccak256Hash(append([]byte("\x19Ethereum Signed Message:\n32"), hash.Bytes()...))
	pubKey, err := crypto.SigToPub(ethSignedHash.Bytes(), sig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to recover public key: %v", err)
	}
	return crypto.PubkeyToAddress(*pubKey), nil
}

func packUserOperation(userOp contract.IUSDTPaymasterUserOperation) []byte {
	parsedABI, err := abi.JSON(strings.NewReader(contract.USDTPaymasterABI))
	if err != nil {
		log.Printf("packUserOperation: failed to parse ABI: %v", err)
		return nil
	}
	method := parsedABI.Methods["executeBatch"]
	if len(method.Inputs) == 0 {
		return nil
	}
	calls := make([]struct {
		To   common.Address
		Data []byte
	}, len(userOp.Calls))
	for i, call := range userOp.Calls {
		calls[i].To = call.To
		calls[i].Data = call.Data
	}
	args := abi.Arguments{method.Inputs[0]}
	packed, err := args.Pack(struct {
		User  common.Address
		Calls []struct {
			To   common.Address
			Data []byte
		}
	}{User: userOp.User, Calls: calls})
	if err != nil {
		log.Printf("packUserOperation: pack error: %v", err)
		return nil
	}
	return packed
}

func signMessage(privateKey *ecdsa.PrivateKey, hash common.Hash) ([]byte, error) {
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %v", err)
	}
	if signature[64] < 27 {
		signature[64] += 27
	}
	return signature, nil
}

func (h *Handlers) GetUserStatus(c *gin.Context) {
	address := c.Param("address")
	userAddr := common.HexToAddress(address)

	balance, err := h.usdt.BalanceOf(userAddr, &bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusOK, models.UserStatusResponse{Address: address, Is7702Bound: false, BoundContract: "", USDTBalance: "0"})
		return
	}

	code, err := h.ethClient.GetEthClient().CodeAt(context.Background(), userAddr, nil)
	if err != nil {
		c.JSON(http.StatusOK, models.UserStatusResponse{Address: address, Is7702Bound: false, BoundContract: "", USDTBalance: balance.String()})
		return
	}

	is7702Bound := len(code) > 0
	boundContract := ""
	if is7702Bound {
		boundContract = h.account.GetAddress().Hex()
	}

	c.JSON(http.StatusOK, models.UserStatusResponse{Address: address, Is7702Bound: is7702Bound, BoundContract: boundContract, USDTBalance: balance.String()})
}

func (h *Handlers) GetFaucetInfo(c *gin.Context) {
	faucetAmount, _ := new(big.Int).SetString("100000000000000000000", 10)
	c.JSON(http.StatusOK, models.FaucetInfoResponse{FaucetAmount: faucetAmount.String(), UsdtAddress: h.usdt.GetAddress().Hex()})
}

func (h *Handlers) ClaimFaucet(c *gin.Context) {
	address := c.Param("address")

	relayer, err := h.relayerPool.SelectIdle()
	if err != nil {
		respondInternalError(c, "no_relayer", "no available relayer")
		return
	}

	transactor, err := h.createTransactor(relayer)
	if err != nil {
		respondInternalError(c, "transactor_failed", err.Error())
		return
	}

	amount, _ := new(big.Int).SetString("100000000000000000000", 10)
	tx, err := h.usdt.MintWithTransactor(common.HexToAddress(address), amount, transactor)
	if err != nil {
		respondInternalError(c, "mint_failed", err.Error())
		return
	}

	receipt, err := h.waitForTx(tx, 30*time.Second)
	if err != nil {
		h.relayerPool.MarkComplete(relayer.Address)
		c.JSON(http.StatusOK, models.FaucetClaimResponse{TxHash: tx.Hash().Hex(), Amount: "100", Status: "pending"})
		return
	}

	h.relayerPool.MarkComplete(relayer.Address)
	status := "success"
	if receipt.Status == 0 {
		status = "failed"
	}

	c.JSON(http.StatusOK, models.FaucetClaimResponse{TxHash: tx.Hash().Hex(), Amount: "100", Status: status})
}

func (h *Handlers) GetRelayers(c *gin.Context) {
	infos := h.relayerPool.GetAll()
	relayerInfos := make([]models.RelayerInfo, len(infos))
	for i, info := range infos {
		isActive, _ := h.paymaster.IsRelayer(common.HexToAddress(info.Address))
		relayerInfos[i] = models.RelayerInfo{Address: info.Address, PendingTx: info.PendingTx, IsActive: isActive}
	}
	c.JSON(http.StatusOK, models.RelayerStatusResponse{Relayers: relayerInfos})
}

func (h *Handlers) TransferUSDT(c *gin.Context) {
	var req models.TransferUSDTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondBadRequest(c, err)
		return
	}

	log.Printf("[TransferUSDT] Request: user=%s, target=%s, amount=%s", req.UserAddress, req.TargetAddress, req.Amount)

	userAddr := common.HexToAddress(req.UserAddress)
	targetAddr := common.HexToAddress(req.TargetAddress)

	amount := new(big.Int)
	if _, ok := amount.SetString(req.Amount, 10); !ok {
		respondBadRequest(c, fmt.Errorf("amount must be a valid integer"))
		return
	}

	sigHex := req.Signature
	if strings.HasPrefix(sigHex, "0x") {
		sigHex = sigHex[2:]
	}
	signature := common.Hex2Bytes(sigHex)
	if len(signature) != 65 {
		respondBadRequest(c, fmt.Errorf("signature must be 65 bytes"))
		return
	}

	buffer, _ := new(big.Int).SetString("1000000000000000000", 10)
	approveAmount := new(big.Int).Add(amount, buffer)
	approveData, err := getApproveData(h.paymaster.GetAddress(), approveAmount)
	if err != nil {
		respondInternalError(c, "encode_failed", err.Error())
		return
	}

	transferData, err := getTransferData(targetAddr, amount)
	if err != nil {
		respondInternalError(c, "encode_failed", err.Error())
		return
	}

	userOp := contract.IUSDTPaymasterUserOperation{
		User: userAddr,
		Calls: []contract.IUSDTPaymasterCall{
			{To: h.usdt.GetAddress(), Data: approveData},
			{To: h.usdt.GetAddress(), Data: transferData},
		},
	}

	packed := packUserOperation(userOp)
	userOpHash := crypto.Keccak256Hash(packed)

	signer, err := recoverSigner(userOpHash, signature)
	if err != nil {
		log.Printf("[TransferUSDT] Signature recovery error: %v", err)
	} else {
		log.Printf("[TransferUSDT] Recovered signer: %s, expected user: %s", signer.Hex(), userAddr.Hex())
	}

	relayer, err := h.relayerPool.SelectIdle()
	if err != nil {
		respondInternalError(c, "no_relayer", "no available relayer")
		return
	}

	h.relayerPool.MarkPending(relayer.Address)

	transactor, err := h.createTransactor(relayer)
	if err != nil {
		h.relayerPool.MarkComplete(relayer.Address)
		respondInternalError(c, "transactor_failed", err.Error())
		return
	}

	tx, err := h.paymaster.ExecuteBatchWithTransactor(transactor, userOp, signature)
	if err != nil {
		h.relayerPool.MarkComplete(relayer.Address)
		respondInternalError(c, "execute_batch_failed", err.Error())
		return
	}

	receipt, err := h.waitForTx(tx, 30*time.Second)
	if err != nil {
		c.JSON(http.StatusOK, models.TransferUSDTResponse{TxHash: tx.Hash().Hex(), Status: "pending", Compensation: "0", GasUsed: 0})
		return
	}

	h.relayerPool.MarkComplete(relayer.Address)

	status := "success"
	if receipt.Status == 0 {
		status = "failed"
	}

	var gasUsed uint64 = receipt.GasUsed
	compensation := "0"

	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Address == h.paymaster.GetAddress() {
			if log.Topics[0] == common.HexToHash("0x41b62e728bad24c9793175686912a5d3533ba7b6e0ab65afb69d36a7230354d2") {
				event, err := h.paymaster.ParseBatchExecuted(*log)
				if err == nil {
					gasUsed = event.GasUsed.Uint64()
					compensation = event.Compensation.String()
				}
			}
		}
	}

	c.JSON(http.StatusOK, models.TransferUSDTResponse{TxHash: tx.Hash().Hex(), Status: status, Compensation: compensation, GasUsed: gasUsed})
}
