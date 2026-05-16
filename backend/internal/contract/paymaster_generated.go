// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IUSDTPaymasterCall is an auto generated low-level Go binding around an user-defined struct.
type IUSDTPaymasterCall struct {
	To   common.Address
	Data []byte
}

// IUSDTPaymasterUserOperation is an auto generated low-level Go binding around an user-defined struct.
type IUSDTPaymasterUserOperation struct {
	User  common.Address
	Calls []IUSDTPaymasterCall
}

// USDTPaymasterMetaData contains all meta data concerning the USDTPaymaster contract.
var USDTPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addRelayer\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structIUSDTPaymaster.UserOperation\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structIUSDTPaymaster.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"usdtTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracleAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isRelayer\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeRelayer\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeRate\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOracle\",\"inputs\":[{\"name\":\"_oracleAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"usdtToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BatchExecuted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"compensation\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRateUpdated\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRecipientUpdated\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleUpdated\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RelayerAdded\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RelayerRemoved\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRelayer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// USDTPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use USDTPaymasterMetaData.ABI instead.
var USDTPaymasterABI = USDTPaymasterMetaData.ABI

// USDTPaymaster is an auto generated Go binding around an Ethereum contract.
type USDTPaymaster struct {
	USDTPaymasterCaller     // Read-only binding to the contract
	USDTPaymasterTransactor // Write-only binding to the contract
	USDTPaymasterFilterer   // Log filterer for contract events
}

// USDTPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDTPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDTPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDTPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDTPaymasterSession struct {
	Contract     *USDTPaymaster    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDTPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDTPaymasterCallerSession struct {
	Contract *USDTPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// USDTPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDTPaymasterTransactorSession struct {
	Contract     *USDTPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// USDTPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDTPaymasterRaw struct {
	Contract *USDTPaymaster // Generic contract binding to access the raw methods on
}

// USDTPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDTPaymasterCallerRaw struct {
	Contract *USDTPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// USDTPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDTPaymasterTransactorRaw struct {
	Contract *USDTPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDTPaymaster creates a new instance of USDTPaymaster, bound to a specific deployed contract.
func NewUSDTPaymaster(address common.Address, backend bind.ContractBackend) (*USDTPaymaster, error) {
	contract, err := bindUSDTPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDTPaymaster{USDTPaymasterCaller: USDTPaymasterCaller{contract: contract}, USDTPaymasterTransactor: USDTPaymasterTransactor{contract: contract}, USDTPaymasterFilterer: USDTPaymasterFilterer{contract: contract}}, nil
}

// NewUSDTPaymasterCaller creates a new read-only instance of USDTPaymaster, bound to a specific deployed contract.
func NewUSDTPaymasterCaller(address common.Address, caller bind.ContractCaller) (*USDTPaymasterCaller, error) {
	contract, err := bindUSDTPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterCaller{contract: contract}, nil
}

// NewUSDTPaymasterTransactor creates a new write-only instance of USDTPaymaster, bound to a specific deployed contract.
func NewUSDTPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*USDTPaymasterTransactor, error) {
	contract, err := bindUSDTPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterTransactor{contract: contract}, nil
}

// NewUSDTPaymasterFilterer creates a new log filterer instance of USDTPaymaster, bound to a specific deployed contract.
func NewUSDTPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*USDTPaymasterFilterer, error) {
	contract, err := bindUSDTPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterFilterer{contract: contract}, nil
}

// bindUSDTPaymaster binds a generic wrapper to an already deployed contract.
func bindUSDTPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDTPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDTPaymaster *USDTPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDTPaymaster.Contract.USDTPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDTPaymaster *USDTPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.USDTPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDTPaymaster *USDTPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.USDTPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDTPaymaster *USDTPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDTPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDTPaymaster *USDTPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDTPaymaster *USDTPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_USDTPaymaster *USDTPaymasterCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDTPaymaster.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_USDTPaymaster *USDTPaymasterSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _USDTPaymaster.Contract.UPGRADEINTERFACEVERSION(&_USDTPaymaster.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_USDTPaymaster *USDTPaymasterCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _USDTPaymaster.Contract.UPGRADEINTERFACEVERSION(&_USDTPaymaster.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_USDTPaymaster *USDTPaymasterCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDTPaymaster.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_USDTPaymaster *USDTPaymasterSession) FeeRate() (*big.Int, error) {
	return _USDTPaymaster.Contract.FeeRate(&_USDTPaymaster.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_USDTPaymaster *USDTPaymasterCallerSession) FeeRate() (*big.Int, error) {
	return _USDTPaymaster.Contract.FeeRate(&_USDTPaymaster.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_USDTPaymaster *USDTPaymasterCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDTPaymaster.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_USDTPaymaster *USDTPaymasterSession) FeeRecipient() (common.Address, error) {
	return _USDTPaymaster.Contract.FeeRecipient(&_USDTPaymaster.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_USDTPaymaster *USDTPaymasterCallerSession) FeeRecipient() (common.Address, error) {
	return _USDTPaymaster.Contract.FeeRecipient(&_USDTPaymaster.CallOpts)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_USDTPaymaster *USDTPaymasterCaller) IsRelayer(opts *bind.CallOpts, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _USDTPaymaster.contract.Call(opts, &out, "isRelayer", relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_USDTPaymaster *USDTPaymasterSession) IsRelayer(relayer common.Address) (bool, error) {
	return _USDTPaymaster.Contract.IsRelayer(&_USDTPaymaster.CallOpts, relayer)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_USDTPaymaster *USDTPaymasterCallerSession) IsRelayer(relayer common.Address) (bool, error) {
	return _USDTPaymaster.Contract.IsRelayer(&_USDTPaymaster.CallOpts, relayer)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_USDTPaymaster *USDTPaymasterCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDTPaymaster.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_USDTPaymaster *USDTPaymasterSession) Oracle() (common.Address, error) {
	return _USDTPaymaster.Contract.Oracle(&_USDTPaymaster.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_USDTPaymaster *USDTPaymasterCallerSession) Oracle() (common.Address, error) {
	return _USDTPaymaster.Contract.Oracle(&_USDTPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDTPaymaster *USDTPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDTPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDTPaymaster *USDTPaymasterSession) Owner() (common.Address, error) {
	return _USDTPaymaster.Contract.Owner(&_USDTPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDTPaymaster *USDTPaymasterCallerSession) Owner() (common.Address, error) {
	return _USDTPaymaster.Contract.Owner(&_USDTPaymaster.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_USDTPaymaster *USDTPaymasterCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _USDTPaymaster.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_USDTPaymaster *USDTPaymasterSession) ProxiableUUID() ([32]byte, error) {
	return _USDTPaymaster.Contract.ProxiableUUID(&_USDTPaymaster.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_USDTPaymaster *USDTPaymasterCallerSession) ProxiableUUID() ([32]byte, error) {
	return _USDTPaymaster.Contract.ProxiableUUID(&_USDTPaymaster.CallOpts)
}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_USDTPaymaster *USDTPaymasterCaller) UsdtToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDTPaymaster.contract.Call(opts, &out, "usdtToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_USDTPaymaster *USDTPaymasterSession) UsdtToken() (common.Address, error) {
	return _USDTPaymaster.Contract.UsdtToken(&_USDTPaymaster.CallOpts)
}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_USDTPaymaster *USDTPaymasterCallerSession) UsdtToken() (common.Address, error) {
	return _USDTPaymaster.Contract.UsdtToken(&_USDTPaymaster.CallOpts)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address relayer) returns()
func (_USDTPaymaster *USDTPaymasterTransactor) AddRelayer(opts *bind.TransactOpts, relayer common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "addRelayer", relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address relayer) returns()
func (_USDTPaymaster *USDTPaymasterSession) AddRelayer(relayer common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.AddRelayer(&_USDTPaymaster.TransactOpts, relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address relayer) returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) AddRelayer(relayer common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.AddRelayer(&_USDTPaymaster.TransactOpts, relayer)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xeee2121c.
//
// Solidity: function executeBatch((address,(address,bytes)[]) userOp, bytes signature) returns()
func (_USDTPaymaster *USDTPaymasterTransactor) ExecuteBatch(opts *bind.TransactOpts, userOp IUSDTPaymasterUserOperation, signature []byte) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "executeBatch", userOp, signature)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xeee2121c.
//
// Solidity: function executeBatch((address,(address,bytes)[]) userOp, bytes signature) returns()
func (_USDTPaymaster *USDTPaymasterSession) ExecuteBatch(userOp IUSDTPaymasterUserOperation, signature []byte) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.ExecuteBatch(&_USDTPaymaster.TransactOpts, userOp, signature)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xeee2121c.
//
// Solidity: function executeBatch((address,(address,bytes)[]) userOp, bytes signature) returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) ExecuteBatch(userOp IUSDTPaymasterUserOperation, signature []byte) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.ExecuteBatch(&_USDTPaymaster.TransactOpts, userOp, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address usdtTokenAddr, address oracleAddr, address _feeRecipient, address _owner) returns()
func (_USDTPaymaster *USDTPaymasterTransactor) Initialize(opts *bind.TransactOpts, usdtTokenAddr common.Address, oracleAddr common.Address, _feeRecipient common.Address, _owner common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "initialize", usdtTokenAddr, oracleAddr, _feeRecipient, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address usdtTokenAddr, address oracleAddr, address _feeRecipient, address _owner) returns()
func (_USDTPaymaster *USDTPaymasterSession) Initialize(usdtTokenAddr common.Address, oracleAddr common.Address, _feeRecipient common.Address, _owner common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.Initialize(&_USDTPaymaster.TransactOpts, usdtTokenAddr, oracleAddr, _feeRecipient, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address usdtTokenAddr, address oracleAddr, address _feeRecipient, address _owner) returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) Initialize(usdtTokenAddr common.Address, oracleAddr common.Address, _feeRecipient common.Address, _owner common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.Initialize(&_USDTPaymaster.TransactOpts, usdtTokenAddr, oracleAddr, _feeRecipient, _owner)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address relayer) returns()
func (_USDTPaymaster *USDTPaymasterTransactor) RemoveRelayer(opts *bind.TransactOpts, relayer common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "removeRelayer", relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address relayer) returns()
func (_USDTPaymaster *USDTPaymasterSession) RemoveRelayer(relayer common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.RemoveRelayer(&_USDTPaymaster.TransactOpts, relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address relayer) returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) RemoveRelayer(relayer common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.RemoveRelayer(&_USDTPaymaster.TransactOpts, relayer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDTPaymaster *USDTPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDTPaymaster *USDTPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _USDTPaymaster.Contract.RenounceOwnership(&_USDTPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _USDTPaymaster.Contract.RenounceOwnership(&_USDTPaymaster.TransactOpts)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 rate) returns()
func (_USDTPaymaster *USDTPaymasterTransactor) SetFeeRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "setFeeRate", rate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 rate) returns()
func (_USDTPaymaster *USDTPaymasterSession) SetFeeRate(rate *big.Int) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.SetFeeRate(&_USDTPaymaster.TransactOpts, rate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 rate) returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) SetFeeRate(rate *big.Int) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.SetFeeRate(&_USDTPaymaster.TransactOpts, rate)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_USDTPaymaster *USDTPaymasterTransactor) SetFeeRecipient(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "setFeeRecipient", recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_USDTPaymaster *USDTPaymasterSession) SetFeeRecipient(recipient common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.SetFeeRecipient(&_USDTPaymaster.TransactOpts, recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) SetFeeRecipient(recipient common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.SetFeeRecipient(&_USDTPaymaster.TransactOpts, recipient)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracleAddr) returns()
func (_USDTPaymaster *USDTPaymasterTransactor) SetOracle(opts *bind.TransactOpts, _oracleAddr common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "setOracle", _oracleAddr)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracleAddr) returns()
func (_USDTPaymaster *USDTPaymasterSession) SetOracle(_oracleAddr common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.SetOracle(&_USDTPaymaster.TransactOpts, _oracleAddr)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracleAddr) returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) SetOracle(_oracleAddr common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.SetOracle(&_USDTPaymaster.TransactOpts, _oracleAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDTPaymaster *USDTPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDTPaymaster *USDTPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.TransferOwnership(&_USDTPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.TransferOwnership(&_USDTPaymaster.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_USDTPaymaster *USDTPaymasterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _USDTPaymaster.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_USDTPaymaster *USDTPaymasterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.UpgradeToAndCall(&_USDTPaymaster.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_USDTPaymaster *USDTPaymasterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _USDTPaymaster.Contract.UpgradeToAndCall(&_USDTPaymaster.TransactOpts, newImplementation, data)
}

// USDTPaymasterBatchExecutedIterator is returned from FilterBatchExecuted and is used to iterate over the raw logs and unpacked data for BatchExecuted events raised by the USDTPaymaster contract.
type USDTPaymasterBatchExecutedIterator struct {
	Event *USDTPaymasterBatchExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterBatchExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterBatchExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterBatchExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterBatchExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterBatchExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterBatchExecuted represents a BatchExecuted event raised by the USDTPaymaster contract.
type USDTPaymasterBatchExecuted struct {
	User         common.Address
	GasUsed      *big.Int
	Compensation *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBatchExecuted is a free log retrieval operation binding the contract event 0x41b62e728bad24c9793175686912a5d3533ba7b6e0ab65afb69d36a7230354d2.
//
// Solidity: event BatchExecuted(address indexed user, uint256 gasUsed, uint256 compensation)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterBatchExecuted(opts *bind.FilterOpts, user []common.Address) (*USDTPaymasterBatchExecutedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "BatchExecuted", userRule)
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterBatchExecutedIterator{contract: _USDTPaymaster.contract, event: "BatchExecuted", logs: logs, sub: sub}, nil
}

// WatchBatchExecuted is a free log subscription operation binding the contract event 0x41b62e728bad24c9793175686912a5d3533ba7b6e0ab65afb69d36a7230354d2.
//
// Solidity: event BatchExecuted(address indexed user, uint256 gasUsed, uint256 compensation)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchBatchExecuted(opts *bind.WatchOpts, sink chan<- *USDTPaymasterBatchExecuted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "BatchExecuted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterBatchExecuted)
				if err := _USDTPaymaster.contract.UnpackLog(event, "BatchExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchExecuted is a log parse operation binding the contract event 0x41b62e728bad24c9793175686912a5d3533ba7b6e0ab65afb69d36a7230354d2.
//
// Solidity: event BatchExecuted(address indexed user, uint256 gasUsed, uint256 compensation)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseBatchExecuted(log types.Log) (*USDTPaymasterBatchExecuted, error) {
	event := new(USDTPaymasterBatchExecuted)
	if err := _USDTPaymaster.contract.UnpackLog(event, "BatchExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDTPaymasterFeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the USDTPaymaster contract.
type USDTPaymasterFeeRateUpdatedIterator struct {
	Event *USDTPaymasterFeeRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterFeeRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterFeeRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterFeeRateUpdated represents a FeeRateUpdated event raised by the USDTPaymaster contract.
type USDTPaymasterFeeRateUpdated struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x208f1b468d3d61f0f085e975bd9d04367c930d599642faad06695229f3eadcd8.
//
// Solidity: event FeeRateUpdated(uint256 rate)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*USDTPaymasterFeeRateUpdatedIterator, error) {

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterFeeRateUpdatedIterator{contract: _USDTPaymaster.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x208f1b468d3d61f0f085e975bd9d04367c930d599642faad06695229f3eadcd8.
//
// Solidity: event FeeRateUpdated(uint256 rate)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *USDTPaymasterFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterFeeRateUpdated)
				if err := _USDTPaymaster.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeRateUpdated is a log parse operation binding the contract event 0x208f1b468d3d61f0f085e975bd9d04367c930d599642faad06695229f3eadcd8.
//
// Solidity: event FeeRateUpdated(uint256 rate)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseFeeRateUpdated(log types.Log) (*USDTPaymasterFeeRateUpdated, error) {
	event := new(USDTPaymasterFeeRateUpdated)
	if err := _USDTPaymaster.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDTPaymasterFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the USDTPaymaster contract.
type USDTPaymasterFeeRecipientUpdatedIterator struct {
	Event *USDTPaymasterFeeRecipientUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterFeeRecipientUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterFeeRecipientUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the USDTPaymaster contract.
type USDTPaymasterFeeRecipientUpdated struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address recipient)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts) (*USDTPaymasterFeeRecipientUpdatedIterator, error) {

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterFeeRecipientUpdatedIterator{contract: _USDTPaymaster.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address recipient)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *USDTPaymasterFeeRecipientUpdated) (event.Subscription, error) {

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterFeeRecipientUpdated)
				if err := _USDTPaymaster.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address recipient)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseFeeRecipientUpdated(log types.Log) (*USDTPaymasterFeeRecipientUpdated, error) {
	event := new(USDTPaymasterFeeRecipientUpdated)
	if err := _USDTPaymaster.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDTPaymasterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the USDTPaymaster contract.
type USDTPaymasterInitializedIterator struct {
	Event *USDTPaymasterInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterInitialized represents a Initialized event raised by the USDTPaymaster contract.
type USDTPaymasterInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterInitialized(opts *bind.FilterOpts) (*USDTPaymasterInitializedIterator, error) {

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterInitializedIterator{contract: _USDTPaymaster.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *USDTPaymasterInitialized) (event.Subscription, error) {

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterInitialized)
				if err := _USDTPaymaster.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseInitialized(log types.Log) (*USDTPaymasterInitialized, error) {
	event := new(USDTPaymasterInitialized)
	if err := _USDTPaymaster.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDTPaymasterOracleUpdatedIterator is returned from FilterOracleUpdated and is used to iterate over the raw logs and unpacked data for OracleUpdated events raised by the USDTPaymaster contract.
type USDTPaymasterOracleUpdatedIterator struct {
	Event *USDTPaymasterOracleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterOracleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterOracleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterOracleUpdated represents a OracleUpdated event raised by the USDTPaymaster contract.
type USDTPaymasterOracleUpdated struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdated is a free log retrieval operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address oracle)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterOracleUpdated(opts *bind.FilterOpts) (*USDTPaymasterOracleUpdatedIterator, error) {

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "OracleUpdated")
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterOracleUpdatedIterator{contract: _USDTPaymaster.contract, event: "OracleUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleUpdated is a free log subscription operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address oracle)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchOracleUpdated(opts *bind.WatchOpts, sink chan<- *USDTPaymasterOracleUpdated) (event.Subscription, error) {

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "OracleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterOracleUpdated)
				if err := _USDTPaymaster.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleUpdated is a log parse operation binding the contract event 0x3df77beb5db05fcdd70a30fc8adf3f83f9501b68579455adbd100b8180940394.
//
// Solidity: event OracleUpdated(address oracle)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseOracleUpdated(log types.Log) (*USDTPaymasterOracleUpdated, error) {
	event := new(USDTPaymasterOracleUpdated)
	if err := _USDTPaymaster.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDTPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the USDTPaymaster contract.
type USDTPaymasterOwnershipTransferredIterator struct {
	Event *USDTPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the USDTPaymaster contract.
type USDTPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*USDTPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterOwnershipTransferredIterator{contract: _USDTPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDTPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterOwnershipTransferred)
				if err := _USDTPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*USDTPaymasterOwnershipTransferred, error) {
	event := new(USDTPaymasterOwnershipTransferred)
	if err := _USDTPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDTPaymasterRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the USDTPaymaster contract.
type USDTPaymasterRelayerAddedIterator struct {
	Event *USDTPaymasterRelayerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterRelayerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterRelayerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterRelayerAdded represents a RelayerAdded event raised by the USDTPaymaster contract.
type USDTPaymasterRelayerAdded struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterRelayerAdded(opts *bind.FilterOpts, relayer []common.Address) (*USDTPaymasterRelayerAddedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "RelayerAdded", relayerRule)
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterRelayerAddedIterator{contract: _USDTPaymaster.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *USDTPaymasterRelayerAdded, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "RelayerAdded", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterRelayerAdded)
				if err := _USDTPaymaster.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseRelayerAdded(log types.Log) (*USDTPaymasterRelayerAdded, error) {
	event := new(USDTPaymasterRelayerAdded)
	if err := _USDTPaymaster.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDTPaymasterRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the USDTPaymaster contract.
type USDTPaymasterRelayerRemovedIterator struct {
	Event *USDTPaymasterRelayerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterRelayerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterRelayerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterRelayerRemoved represents a RelayerRemoved event raised by the USDTPaymaster contract.
type USDTPaymasterRelayerRemoved struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterRelayerRemoved(opts *bind.FilterOpts, relayer []common.Address) (*USDTPaymasterRelayerRemovedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "RelayerRemoved", relayerRule)
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterRelayerRemovedIterator{contract: _USDTPaymaster.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *USDTPaymasterRelayerRemoved, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "RelayerRemoved", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterRelayerRemoved)
				if err := _USDTPaymaster.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseRelayerRemoved(log types.Log) (*USDTPaymasterRelayerRemoved, error) {
	event := new(USDTPaymasterRelayerRemoved)
	if err := _USDTPaymaster.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDTPaymasterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the USDTPaymaster contract.
type USDTPaymasterUpgradedIterator struct {
	Event *USDTPaymasterUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDTPaymasterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDTPaymasterUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDTPaymasterUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDTPaymasterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDTPaymasterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDTPaymasterUpgraded represents a Upgraded event raised by the USDTPaymaster contract.
type USDTPaymasterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_USDTPaymaster *USDTPaymasterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*USDTPaymasterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _USDTPaymaster.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &USDTPaymasterUpgradedIterator{contract: _USDTPaymaster.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_USDTPaymaster *USDTPaymasterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *USDTPaymasterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _USDTPaymaster.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDTPaymasterUpgraded)
				if err := _USDTPaymaster.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_USDTPaymaster *USDTPaymasterFilterer) ParseUpgraded(log types.Log) (*USDTPaymasterUpgraded, error) {
	event := new(USDTPaymasterUpgraded)
	if err := _USDTPaymaster.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
