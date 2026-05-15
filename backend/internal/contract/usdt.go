package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type USDTContract struct {
	address common.Address
	client  *ethclient.Client
}

func NewUSDT(address string, client *ethclient.Client) *USDTContract {
	return &USDTContract{
		address: common.HexToAddress(address),
		client:  client,
	}
}

func (u *USDTContract) GetAddress() common.Address {
	return u.address
}

func (u *USDTContract) BalanceOf(account common.Address, caller *bind.CallOpts) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (u *USDTContract) FaucetAmount(caller *bind.CallOpts) (*big.Int, error) {
	amount := new(big.Int)
	amount.SetString("100000000000000000000", 10)
	return amount, nil
}
