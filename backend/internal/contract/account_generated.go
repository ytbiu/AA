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

// Simple7702AccountMetaData contains all meta data concerning the Simple7702Account contract.
var Simple7702AccountMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structIUSDTPaymaster.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_paymaster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paymaster\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUSDTPaymaster\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymaster\",\"inputs\":[{\"name\":\"_paymaster\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BatchExecuted\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"callCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// Simple7702AccountABI is the input ABI used to generate the binding from.
// Deprecated: Use Simple7702AccountMetaData.ABI instead.
var Simple7702AccountABI = Simple7702AccountMetaData.ABI

// Simple7702Account is an auto generated Go binding around an Ethereum contract.
type Simple7702Account struct {
	Simple7702AccountCaller     // Read-only binding to the contract
	Simple7702AccountTransactor // Write-only binding to the contract
	Simple7702AccountFilterer   // Log filterer for contract events
}

// Simple7702AccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type Simple7702AccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Simple7702AccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Simple7702AccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Simple7702AccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Simple7702AccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Simple7702AccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Simple7702AccountSession struct {
	Contract     *Simple7702Account // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Simple7702AccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Simple7702AccountCallerSession struct {
	Contract *Simple7702AccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Simple7702AccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Simple7702AccountTransactorSession struct {
	Contract     *Simple7702AccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Simple7702AccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type Simple7702AccountRaw struct {
	Contract *Simple7702Account // Generic contract binding to access the raw methods on
}

// Simple7702AccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Simple7702AccountCallerRaw struct {
	Contract *Simple7702AccountCaller // Generic read-only contract binding to access the raw methods on
}

// Simple7702AccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Simple7702AccountTransactorRaw struct {
	Contract *Simple7702AccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimple7702Account creates a new instance of Simple7702Account, bound to a specific deployed contract.
func NewSimple7702Account(address common.Address, backend bind.ContractBackend) (*Simple7702Account, error) {
	contract, err := bindSimple7702Account(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simple7702Account{Simple7702AccountCaller: Simple7702AccountCaller{contract: contract}, Simple7702AccountTransactor: Simple7702AccountTransactor{contract: contract}, Simple7702AccountFilterer: Simple7702AccountFilterer{contract: contract}}, nil
}

// NewSimple7702AccountCaller creates a new read-only instance of Simple7702Account, bound to a specific deployed contract.
func NewSimple7702AccountCaller(address common.Address, caller bind.ContractCaller) (*Simple7702AccountCaller, error) {
	contract, err := bindSimple7702Account(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Simple7702AccountCaller{contract: contract}, nil
}

// NewSimple7702AccountTransactor creates a new write-only instance of Simple7702Account, bound to a specific deployed contract.
func NewSimple7702AccountTransactor(address common.Address, transactor bind.ContractTransactor) (*Simple7702AccountTransactor, error) {
	contract, err := bindSimple7702Account(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Simple7702AccountTransactor{contract: contract}, nil
}

// NewSimple7702AccountFilterer creates a new log filterer instance of Simple7702Account, bound to a specific deployed contract.
func NewSimple7702AccountFilterer(address common.Address, filterer bind.ContractFilterer) (*Simple7702AccountFilterer, error) {
	contract, err := bindSimple7702Account(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Simple7702AccountFilterer{contract: contract}, nil
}

// bindSimple7702Account binds a generic wrapper to an already deployed contract.
func bindSimple7702Account(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Simple7702AccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple7702Account *Simple7702AccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple7702Account.Contract.Simple7702AccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple7702Account *Simple7702AccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple7702Account.Contract.Simple7702AccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple7702Account *Simple7702AccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple7702Account.Contract.Simple7702AccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple7702Account *Simple7702AccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple7702Account.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple7702Account *Simple7702AccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple7702Account.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple7702Account *Simple7702AccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple7702Account.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Simple7702Account *Simple7702AccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Simple7702Account.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Simple7702Account *Simple7702AccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Simple7702Account.Contract.UPGRADEINTERFACEVERSION(&_Simple7702Account.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Simple7702Account *Simple7702AccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Simple7702Account.Contract.UPGRADEINTERFACEVERSION(&_Simple7702Account.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_Simple7702Account *Simple7702AccountCaller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Simple7702Account.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_Simple7702Account *Simple7702AccountSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _Simple7702Account.Contract.IsValidSignature(&_Simple7702Account.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_Simple7702Account *Simple7702AccountCallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _Simple7702Account.Contract.IsValidSignature(&_Simple7702Account.CallOpts, hash, signature)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simple7702Account *Simple7702AccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simple7702Account.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simple7702Account *Simple7702AccountSession) Owner() (common.Address, error) {
	return _Simple7702Account.Contract.Owner(&_Simple7702Account.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simple7702Account *Simple7702AccountCallerSession) Owner() (common.Address, error) {
	return _Simple7702Account.Contract.Owner(&_Simple7702Account.CallOpts)
}

// Paymaster is a free data retrieval call binding the contract method 0x16e4cbf9.
//
// Solidity: function paymaster() view returns(address)
func (_Simple7702Account *Simple7702AccountCaller) Paymaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simple7702Account.contract.Call(opts, &out, "paymaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Paymaster is a free data retrieval call binding the contract method 0x16e4cbf9.
//
// Solidity: function paymaster() view returns(address)
func (_Simple7702Account *Simple7702AccountSession) Paymaster() (common.Address, error) {
	return _Simple7702Account.Contract.Paymaster(&_Simple7702Account.CallOpts)
}

// Paymaster is a free data retrieval call binding the contract method 0x16e4cbf9.
//
// Solidity: function paymaster() view returns(address)
func (_Simple7702Account *Simple7702AccountCallerSession) Paymaster() (common.Address, error) {
	return _Simple7702Account.Contract.Paymaster(&_Simple7702Account.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Simple7702Account *Simple7702AccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Simple7702Account.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Simple7702Account *Simple7702AccountSession) ProxiableUUID() ([32]byte, error) {
	return _Simple7702Account.Contract.ProxiableUUID(&_Simple7702Account.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Simple7702Account *Simple7702AccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Simple7702Account.Contract.ProxiableUUID(&_Simple7702Account.CallOpts)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x21566279.
//
// Solidity: function executeBatch((address,bytes)[] calls) returns()
func (_Simple7702Account *Simple7702AccountTransactor) ExecuteBatch(opts *bind.TransactOpts, calls []IUSDTPaymasterCall) (*types.Transaction, error) {
	return _Simple7702Account.contract.Transact(opts, "executeBatch", calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x21566279.
//
// Solidity: function executeBatch((address,bytes)[] calls) returns()
func (_Simple7702Account *Simple7702AccountSession) ExecuteBatch(calls []IUSDTPaymasterCall) (*types.Transaction, error) {
	return _Simple7702Account.Contract.ExecuteBatch(&_Simple7702Account.TransactOpts, calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x21566279.
//
// Solidity: function executeBatch((address,bytes)[] calls) returns()
func (_Simple7702Account *Simple7702AccountTransactorSession) ExecuteBatch(calls []IUSDTPaymasterCall) (*types.Transaction, error) {
	return _Simple7702Account.Contract.ExecuteBatch(&_Simple7702Account.TransactOpts, calls)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _paymaster, address _owner) returns()
func (_Simple7702Account *Simple7702AccountTransactor) Initialize(opts *bind.TransactOpts, _paymaster common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Simple7702Account.contract.Transact(opts, "initialize", _paymaster, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _paymaster, address _owner) returns()
func (_Simple7702Account *Simple7702AccountSession) Initialize(_paymaster common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Simple7702Account.Contract.Initialize(&_Simple7702Account.TransactOpts, _paymaster, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _paymaster, address _owner) returns()
func (_Simple7702Account *Simple7702AccountTransactorSession) Initialize(_paymaster common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Simple7702Account.Contract.Initialize(&_Simple7702Account.TransactOpts, _paymaster, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Simple7702Account *Simple7702AccountTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple7702Account.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Simple7702Account *Simple7702AccountSession) RenounceOwnership() (*types.Transaction, error) {
	return _Simple7702Account.Contract.RenounceOwnership(&_Simple7702Account.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Simple7702Account *Simple7702AccountTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Simple7702Account.Contract.RenounceOwnership(&_Simple7702Account.TransactOpts)
}

// SetPaymaster is a paid mutator transaction binding the contract method 0x2a97fa77.
//
// Solidity: function setPaymaster(address _paymaster) returns()
func (_Simple7702Account *Simple7702AccountTransactor) SetPaymaster(opts *bind.TransactOpts, _paymaster common.Address) (*types.Transaction, error) {
	return _Simple7702Account.contract.Transact(opts, "setPaymaster", _paymaster)
}

// SetPaymaster is a paid mutator transaction binding the contract method 0x2a97fa77.
//
// Solidity: function setPaymaster(address _paymaster) returns()
func (_Simple7702Account *Simple7702AccountSession) SetPaymaster(_paymaster common.Address) (*types.Transaction, error) {
	return _Simple7702Account.Contract.SetPaymaster(&_Simple7702Account.TransactOpts, _paymaster)
}

// SetPaymaster is a paid mutator transaction binding the contract method 0x2a97fa77.
//
// Solidity: function setPaymaster(address _paymaster) returns()
func (_Simple7702Account *Simple7702AccountTransactorSession) SetPaymaster(_paymaster common.Address) (*types.Transaction, error) {
	return _Simple7702Account.Contract.SetPaymaster(&_Simple7702Account.TransactOpts, _paymaster)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Simple7702Account *Simple7702AccountTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Simple7702Account.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Simple7702Account *Simple7702AccountSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Simple7702Account.Contract.TransferOwnership(&_Simple7702Account.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Simple7702Account *Simple7702AccountTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Simple7702Account.Contract.TransferOwnership(&_Simple7702Account.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Simple7702Account *Simple7702AccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Simple7702Account.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Simple7702Account *Simple7702AccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Simple7702Account.Contract.UpgradeToAndCall(&_Simple7702Account.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Simple7702Account *Simple7702AccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Simple7702Account.Contract.UpgradeToAndCall(&_Simple7702Account.TransactOpts, newImplementation, data)
}

// Simple7702AccountBatchExecutedIterator is returned from FilterBatchExecuted and is used to iterate over the raw logs and unpacked data for BatchExecuted events raised by the Simple7702Account contract.
type Simple7702AccountBatchExecutedIterator struct {
	Event *Simple7702AccountBatchExecuted // Event containing the contract specifics and raw log

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
func (it *Simple7702AccountBatchExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Simple7702AccountBatchExecuted)
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
		it.Event = new(Simple7702AccountBatchExecuted)
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
func (it *Simple7702AccountBatchExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Simple7702AccountBatchExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Simple7702AccountBatchExecuted represents a BatchExecuted event raised by the Simple7702Account contract.
type Simple7702AccountBatchExecuted struct {
	Caller    common.Address
	CallCount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchExecuted is a free log retrieval operation binding the contract event 0x7ffb1df8fb1cb14bf915045200c5b519c9cdd3c996c3fa9e781810674194029c.
//
// Solidity: event BatchExecuted(address caller, uint256 callCount)
func (_Simple7702Account *Simple7702AccountFilterer) FilterBatchExecuted(opts *bind.FilterOpts) (*Simple7702AccountBatchExecutedIterator, error) {

	logs, sub, err := _Simple7702Account.contract.FilterLogs(opts, "BatchExecuted")
	if err != nil {
		return nil, err
	}
	return &Simple7702AccountBatchExecutedIterator{contract: _Simple7702Account.contract, event: "BatchExecuted", logs: logs, sub: sub}, nil
}

// WatchBatchExecuted is a free log subscription operation binding the contract event 0x7ffb1df8fb1cb14bf915045200c5b519c9cdd3c996c3fa9e781810674194029c.
//
// Solidity: event BatchExecuted(address caller, uint256 callCount)
func (_Simple7702Account *Simple7702AccountFilterer) WatchBatchExecuted(opts *bind.WatchOpts, sink chan<- *Simple7702AccountBatchExecuted) (event.Subscription, error) {

	logs, sub, err := _Simple7702Account.contract.WatchLogs(opts, "BatchExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Simple7702AccountBatchExecuted)
				if err := _Simple7702Account.contract.UnpackLog(event, "BatchExecuted", log); err != nil {
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

// ParseBatchExecuted is a log parse operation binding the contract event 0x7ffb1df8fb1cb14bf915045200c5b519c9cdd3c996c3fa9e781810674194029c.
//
// Solidity: event BatchExecuted(address caller, uint256 callCount)
func (_Simple7702Account *Simple7702AccountFilterer) ParseBatchExecuted(log types.Log) (*Simple7702AccountBatchExecuted, error) {
	event := new(Simple7702AccountBatchExecuted)
	if err := _Simple7702Account.contract.UnpackLog(event, "BatchExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Simple7702AccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Simple7702Account contract.
type Simple7702AccountInitializedIterator struct {
	Event *Simple7702AccountInitialized // Event containing the contract specifics and raw log

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
func (it *Simple7702AccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Simple7702AccountInitialized)
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
		it.Event = new(Simple7702AccountInitialized)
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
func (it *Simple7702AccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Simple7702AccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Simple7702AccountInitialized represents a Initialized event raised by the Simple7702Account contract.
type Simple7702AccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Simple7702Account *Simple7702AccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*Simple7702AccountInitializedIterator, error) {

	logs, sub, err := _Simple7702Account.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Simple7702AccountInitializedIterator{contract: _Simple7702Account.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Simple7702Account *Simple7702AccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Simple7702AccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Simple7702Account.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Simple7702AccountInitialized)
				if err := _Simple7702Account.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Simple7702Account *Simple7702AccountFilterer) ParseInitialized(log types.Log) (*Simple7702AccountInitialized, error) {
	event := new(Simple7702AccountInitialized)
	if err := _Simple7702Account.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Simple7702AccountOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Simple7702Account contract.
type Simple7702AccountOwnershipTransferredIterator struct {
	Event *Simple7702AccountOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Simple7702AccountOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Simple7702AccountOwnershipTransferred)
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
		it.Event = new(Simple7702AccountOwnershipTransferred)
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
func (it *Simple7702AccountOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Simple7702AccountOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Simple7702AccountOwnershipTransferred represents a OwnershipTransferred event raised by the Simple7702Account contract.
type Simple7702AccountOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Simple7702Account *Simple7702AccountFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Simple7702AccountOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Simple7702Account.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Simple7702AccountOwnershipTransferredIterator{contract: _Simple7702Account.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Simple7702Account *Simple7702AccountFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Simple7702AccountOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Simple7702Account.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Simple7702AccountOwnershipTransferred)
				if err := _Simple7702Account.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Simple7702Account *Simple7702AccountFilterer) ParseOwnershipTransferred(log types.Log) (*Simple7702AccountOwnershipTransferred, error) {
	event := new(Simple7702AccountOwnershipTransferred)
	if err := _Simple7702Account.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Simple7702AccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Simple7702Account contract.
type Simple7702AccountUpgradedIterator struct {
	Event *Simple7702AccountUpgraded // Event containing the contract specifics and raw log

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
func (it *Simple7702AccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Simple7702AccountUpgraded)
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
		it.Event = new(Simple7702AccountUpgraded)
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
func (it *Simple7702AccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Simple7702AccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Simple7702AccountUpgraded represents a Upgraded event raised by the Simple7702Account contract.
type Simple7702AccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Simple7702Account *Simple7702AccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Simple7702AccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Simple7702Account.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Simple7702AccountUpgradedIterator{contract: _Simple7702Account.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Simple7702Account *Simple7702AccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Simple7702AccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Simple7702Account.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Simple7702AccountUpgraded)
				if err := _Simple7702Account.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Simple7702Account *Simple7702AccountFilterer) ParseUpgraded(log types.Log) (*Simple7702AccountUpgraded, error) {
	event := new(Simple7702AccountUpgraded)
	if err := _Simple7702Account.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
