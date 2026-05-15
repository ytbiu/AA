package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type OracleContract struct {
	address common.Address
	client  *ethclient.Client
}

func NewOracle(address string, client *ethclient.Client) *OracleContract {
	return &OracleContract{
		address: common.HexToAddress(address),
		client:  client,
	}
}

func (o *OracleContract) GetAddress() common.Address {
	return o.address
}

func (o *OracleContract) GetBNBPrice(caller *bind.CallOpts) (*big.Int, error) {
	return big.NewInt(300000000000), nil // $300 * 10^9
}

func (o *OracleContract) GetUSDTPrice(caller *bind.CallOpts) (*big.Int, error) {
	return big.NewInt(100000000000), nil // $1 * 10^9
}
