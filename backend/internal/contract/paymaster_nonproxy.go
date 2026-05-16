package contract

import (
	"github.com/ethereum/go-ethereum/common"
)

// USDTPaymasterNonProxyUserOperation 等同于 IUSDTPaymasterUserOperation
type USDTPaymasterNonProxyUserOperation struct {
	User  common.Address
	Calls []IUSDTPaymasterCall
}

// USDTPaymasterNonProxyCall 等同于 IUSDTPaymasterCall
type USDTPaymasterNonProxyCall struct {
	To   common.Address
	Data []byte
}

// USDTPaymasterNonProxy 合约地址
type USDTPaymasterNonProxy struct {
	address common.Address
}

func NewUSDTPaymasterNonProxy(address common.Address) *USDTPaymasterNonProxy {
	return &USDTPaymasterNonProxy{address: address}
}

func (p *USDTPaymasterNonProxy) GetAddress() common.Address {
	return p.address
}
