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

type PaymasterContract struct {
	address  common.Address
	client   *ethclient.Client
	contract *USDTPaymaster
}

func NewPaymaster(address string, client *ethclient.Client) *PaymasterContract {
	addr := common.HexToAddress(address)
	contract, err := NewUSDTPaymaster(addr, client)
	if err != nil {
		return &PaymasterContract{
			address: addr,
			client:  client,
		}
	}
	return &PaymasterContract{
		address:  addr,
		client:   client,
		contract: contract,
	}
}

func (p *PaymasterContract) GetAddress() common.Address {
	return p.address
}

func (p *PaymasterContract) IsRelayer(relayerAddr common.Address) (bool, error) {
	if p.contract == nil {
		return false, fmt.Errorf("contract not initialized")
	}
	return p.contract.IsRelayer(&bind.CallOpts{}, relayerAddr)
}

func (p *PaymasterContract) AddRelayerWithTransactor(transactor *bind.TransactOpts, relayerAddr common.Address) (*types.Transaction, error) {
	if p.contract == nil {
		return nil, fmt.Errorf("contract not initialized")
	}
	return p.contract.AddRelayer(transactor, relayerAddr)
}

func (p *PaymasterContract) RemoveRelayerWithTransactor(transactor *bind.TransactOpts, relayerAddr common.Address) (*types.Transaction, error) {
	if p.contract == nil {
		return nil, fmt.Errorf("contract not initialized")
	}
	return p.contract.RemoveRelayer(transactor, relayerAddr)
}

func (p *PaymasterContract) SetFeeRateWithTransactor(transactor *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	if p.contract == nil {
		return nil, fmt.Errorf("contract not initialized")
	}
	return p.contract.SetFeeRate(transactor, rate)
}

func (p *PaymasterContract) SetOracleWithTransactor(transactor *bind.TransactOpts, oracleAddr common.Address) (*types.Transaction, error) {
	if p.contract == nil {
		return nil, fmt.Errorf("contract not initialized")
	}
	return p.contract.SetOracle(transactor, oracleAddr)
}

func (p *PaymasterContract) ExecuteBatchWithTransactor(transactor *bind.TransactOpts, userOp IUSDTPaymasterUserOperation, signature []byte) (*types.Transaction, error) {
	if p.contract == nil {
		return nil, fmt.Errorf("contract not initialized")
	}
	return p.contract.ExecuteBatch(transactor, userOp, signature)
}

func (p *PaymasterContract) GetFeeRate() (*big.Int, error) {
	if p.contract == nil {
		return big.NewInt(0), nil
	}
	return p.contract.FeeRate(&bind.CallOpts{})
}

func (p *PaymasterContract) GetFeeRecipient() (common.Address, error) {
	if p.contract == nil {
		return common.Address{}, nil
	}
	return p.contract.FeeRecipient(&bind.CallOpts{})
}

func (p *PaymasterContract) GetOracle() (common.Address, error) {
	if p.contract == nil {
		return common.Address{}, nil
	}
	return p.contract.Oracle(&bind.CallOpts{})
}

func (p *PaymasterContract) GetUSDTToken() (common.Address, error) {
	if p.contract == nil {
		return common.Address{}, nil
	}
	return p.contract.UsdtToken(&bind.CallOpts{})
}

func (p *PaymasterContract) WaitForTx(tx *types.Transaction, timeout time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, p.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx: %v", err)
	}
	return receipt, nil
}

func (p *PaymasterContract) ParseBatchExecuted(log types.Log) (*USDTPaymasterBatchExecuted, error) {
	if p.contract == nil {
		return nil, fmt.Errorf("contract not initialized")
	}
	return p.contract.ParseBatchExecuted(log)
}
