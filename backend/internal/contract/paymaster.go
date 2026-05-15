package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PaymasterContract struct {
	address common.Address
	client  *ethclient.Client
}

func NewPaymaster(address string, client *ethclient.Client) *PaymasterContract {
	return &PaymasterContract{
		address: common.HexToAddress(address),
		client:  client,
	}
}

func (p *PaymasterContract) GetAddress() common.Address {
	return p.address
}

func (p *PaymasterContract) IsRelayer(relayer common.Address, caller *bind.CallOpts) (bool, error) {
	return true, nil
}

func (p *PaymasterContract) GetFeeRate(caller *bind.CallOpts) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (p *PaymasterContract) GetFeeRecipient(caller *bind.CallOpts) (common.Address, error) {
	return common.HexToAddress("0x0"), nil
}
