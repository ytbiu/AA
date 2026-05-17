package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const priceOracleABI = `[{"inputs":[],"name":"getBNBPriceInUSDT","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"router","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`

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

func (o *OracleContract) GetBNBPriceInUSDT(caller *bind.CallOpts) (*big.Int, error) {
	parsedABI, err := abi.JSON(strings.NewReader(priceOracleABI))
	if err != nil {
		return nil, err
	}

	bound := bind.NewBoundContract(o.address, parsedABI, o.client, nil, nil)
	var out []interface{}
	if err := bound.Call(caller, &out, "getBNBPriceInUSDT"); err != nil {
		return nil, err
	}

	price := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return price, nil
}
