package contract

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AccountContract struct {
	address  common.Address
	client   *ethclient.Client
	contract *Simple7702Account
}

func NewAccountContract(address string, client *ethclient.Client) *AccountContract {
	addr := common.HexToAddress(address)
	contract, err := NewSimple7702Account(addr, client)
	if err != nil {
		return &AccountContract{
			address: addr,
			client:  client,
		}
	}
	return &AccountContract{
		address:  addr,
		client:   client,
		contract: contract,
	}
}

func (a *AccountContract) GetAddress() common.Address {
	return a.address
}

func (a *AccountContract) GetOwner() (common.Address, error) {
	if a.contract == nil {
		return common.Address{}, fmt.Errorf("contract not initialized")
	}
	return a.contract.Owner(&bind.CallOpts{})
}

func (a *AccountContract) GetPaymaster() (common.Address, error) {
	if a.contract == nil {
		return common.Address{}, fmt.Errorf("contract not initialized")
	}
	return a.contract.Paymaster(&bind.CallOpts{})
}

func (a *AccountContract) ExecuteBatchWithTransactor(transactor *bind.TransactOpts, calls []IUSDTPaymasterCall) (*types.Transaction, error) {
	if a.contract == nil {
		return nil, fmt.Errorf("contract not initialized")
	}
	return a.contract.ExecuteBatch(transactor, calls)
}

func (a *AccountContract) IsValidSignature(hash [32]byte, signature []byte) (bool, error) {
	if a.contract == nil {
		return false, fmt.Errorf("contract not initialized")
	}
	result, err := a.contract.IsValidSignature(&bind.CallOpts{}, hash, signature)
	if err != nil {
		return false, err
	}
	return result == [4]byte{0x16, 0x26, 0xba, 0x7e}, nil // EIP-1271 magic value
}

func (a *AccountContract) WaitForTx(tx *types.Transaction, timeout time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, a.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx: %v", err)
	}
	return receipt, nil
}
