package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type TxManager struct {
	client  Client
	chainID *big.Int
}

func NewTxManager(client *Client) *TxManager {
	return &TxManager{
		client:  *client,
		chainID: client.GetChainID(),
	}
}

func (tm *TxManager) SendTransaction(
	privKey *ecdsa.PrivateKey,
	to common.Address,
	data []byte,
	gasLimit uint64,
) (common.Hash, error) {
	from := crypto.PubkeyToAddress(privKey.PublicKey)

	nonce, err := tm.client.GetPendingNonce(from)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := tm.client.SuggestGasPrice()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get gas price: %v", err)
	}

	tx := types.NewTransaction(
		nonce,
		to,
		big.NewInt(0),
		gasLimit,
		gasPrice,
		data,
	)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(tm.chainID), privKey)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to sign tx: %v", err)
	}

	err = tm.client.GetEthClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send tx: %v", err)
	}

	return signedTx.Hash(), nil
}

func (tm *TxManager) WaitForReceipt(txHash common.Hash, timeout time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout waiting for tx receipt")
		case <-ticker.C:
			receipt, err := tm.client.GetEthClient().TransactionReceipt(context.Background(), txHash)
			if err == ethereum.NotFound {
				continue
			}
			if err != nil {
				return nil, fmt.Errorf("error getting receipt: %v", err)
			}
			return receipt, nil
		}
	}
}

func (tm *TxManager) GetTransactionStatus(txHash common.Hash) (string, error) {
	_, pending, err := tm.client.GetEthClient().TransactionByHash(context.Background(), txHash)
	if err != nil {
		return "", err
	}
	if pending {
		return "pending", nil
	}

	receipt, err := tm.client.GetEthClient().TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return "", err
	}
	if receipt.Status == 1 {
		return "success", nil
	}
	return "failed", nil
}

func (tm *TxManager) EstimateGas(from, to common.Address, data []byte) (uint64, error) {
	msg := ethereum.CallMsg{
		From: from,
		To:   &to,
		Data: data,
	}
	return tm.client.GetEthClient().EstimateGas(context.Background(), msg)
}
