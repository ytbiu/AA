package contract

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type USDTContract struct {
	address  common.Address
	client   *ethclient.Client
	contract *MockUSDT
}

func NewUSDT(address string, client *ethclient.Client) *USDTContract {
	addr := common.HexToAddress(address)
	contract, err := NewMockUSDT(addr, client)
	if err != nil {
		// Return a struct without the contract binding if initialization fails
		return &USDTContract{
			address: addr,
			client:  client,
		}
	}
	return &USDTContract{
		address:  addr,
		client:   client,
		contract: contract,
	}
}

func (u *USDTContract) GetAddress() common.Address {
	return u.address
}

func (u *USDTContract) BalanceOf(account common.Address, opts *bind.CallOpts) (*big.Int, error) {
	if u.contract == nil {
		return big.NewInt(0), nil
	}
	return u.contract.BalanceOf(opts, account)
}

func (u *USDTContract) FaucetAmount(opts *bind.CallOpts) (*big.Int, error) {
	if u.contract == nil {
		// Default 100 USDT with 18 decimals
		amount := new(big.Int)
		amount.SetString("100000000000000000000", 10)
		return amount, nil
	}
	return u.contract.FAUCETAMOUNT(opts)
}

func (u *USDTContract) Mint(to string, amount string) (string, error) {
	return "", fmt.Errorf("Mint requires transactor - use MintWithTransactor instead")
}

// MintWithTransactor mints USDT to an address using the provided transactor
func (u *USDTContract) MintWithTransactor(to common.Address, amount *big.Int, transactor *bind.TransactOpts) (*types.Transaction, error) {
	if u.contract == nil {
		return nil, fmt.Errorf("contract not initialized")
	}
	return u.contract.Mint(transactor, to, amount)
}

// WaitForTx waits for a transaction to be mined and returns the receipt
func (u *USDTContract) WaitForTx(tx *types.Transaction, timeout time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, u.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx: %v", err)
	}
	return receipt, nil
}
