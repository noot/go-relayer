// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mforwarder

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
)

// IMinimalForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type IMinimalForwarderForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

// IMinimalForwarderMetaData contains all meta data concerning the IMinimalForwarder contract.
var IMinimalForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IMinimalForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use IMinimalForwarderMetaData.ABI instead.
var IMinimalForwarderABI = IMinimalForwarderMetaData.ABI

// IMinimalForwarder is an auto generated Go binding around an Ethereum contract.
type IMinimalForwarder struct {
	IMinimalForwarderCaller     // Read-only binding to the contract
	IMinimalForwarderTransactor // Write-only binding to the contract
	IMinimalForwarderFilterer   // Log filterer for contract events
}

// IMinimalForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMinimalForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinimalForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMinimalForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinimalForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMinimalForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinimalForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMinimalForwarderSession struct {
	Contract     *IMinimalForwarder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IMinimalForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMinimalForwarderCallerSession struct {
	Contract *IMinimalForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IMinimalForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMinimalForwarderTransactorSession struct {
	Contract     *IMinimalForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IMinimalForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMinimalForwarderRaw struct {
	Contract *IMinimalForwarder // Generic contract binding to access the raw methods on
}

// IMinimalForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMinimalForwarderCallerRaw struct {
	Contract *IMinimalForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// IMinimalForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMinimalForwarderTransactorRaw struct {
	Contract *IMinimalForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMinimalForwarder creates a new instance of IMinimalForwarder, bound to a specific deployed contract.
func NewIMinimalForwarder(address common.Address, backend bind.ContractBackend) (*IMinimalForwarder, error) {
	contract, err := bindIMinimalForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMinimalForwarder{IMinimalForwarderCaller: IMinimalForwarderCaller{contract: contract}, IMinimalForwarderTransactor: IMinimalForwarderTransactor{contract: contract}, IMinimalForwarderFilterer: IMinimalForwarderFilterer{contract: contract}}, nil
}

// NewIMinimalForwarderCaller creates a new read-only instance of IMinimalForwarder, bound to a specific deployed contract.
func NewIMinimalForwarderCaller(address common.Address, caller bind.ContractCaller) (*IMinimalForwarderCaller, error) {
	contract, err := bindIMinimalForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMinimalForwarderCaller{contract: contract}, nil
}

// NewIMinimalForwarderTransactor creates a new write-only instance of IMinimalForwarder, bound to a specific deployed contract.
func NewIMinimalForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*IMinimalForwarderTransactor, error) {
	contract, err := bindIMinimalForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMinimalForwarderTransactor{contract: contract}, nil
}

// NewIMinimalForwarderFilterer creates a new log filterer instance of IMinimalForwarder, bound to a specific deployed contract.
func NewIMinimalForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*IMinimalForwarderFilterer, error) {
	contract, err := bindIMinimalForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMinimalForwarderFilterer{contract: contract}, nil
}

// bindIMinimalForwarder binds a generic wrapper to an already deployed contract.
func bindIMinimalForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMinimalForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMinimalForwarder *IMinimalForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMinimalForwarder.Contract.IMinimalForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMinimalForwarder *IMinimalForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMinimalForwarder.Contract.IMinimalForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMinimalForwarder *IMinimalForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMinimalForwarder.Contract.IMinimalForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMinimalForwarder *IMinimalForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMinimalForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMinimalForwarder *IMinimalForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMinimalForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMinimalForwarder *IMinimalForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMinimalForwarder.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_IMinimalForwarder *IMinimalForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IMinimalForwarder.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_IMinimalForwarder *IMinimalForwarderSession) GetNonce(from common.Address) (*big.Int, error) {
	return _IMinimalForwarder.Contract.GetNonce(&_IMinimalForwarder.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_IMinimalForwarder *IMinimalForwarderCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _IMinimalForwarder.Contract.GetNonce(&_IMinimalForwarder.CallOpts, from)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_IMinimalForwarder *IMinimalForwarderCaller) Verify(opts *bind.CallOpts, req IMinimalForwarderForwardRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _IMinimalForwarder.contract.Call(opts, &out, "verify", req, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_IMinimalForwarder *IMinimalForwarderSession) Verify(req IMinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _IMinimalForwarder.Contract.Verify(&_IMinimalForwarder.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_IMinimalForwarder *IMinimalForwarderCallerSession) Verify(req IMinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _IMinimalForwarder.Contract.Verify(&_IMinimalForwarder.CallOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_IMinimalForwarder *IMinimalForwarderTransactor) Execute(opts *bind.TransactOpts, req IMinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _IMinimalForwarder.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_IMinimalForwarder *IMinimalForwarderSession) Execute(req IMinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _IMinimalForwarder.Contract.Execute(&_IMinimalForwarder.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_IMinimalForwarder *IMinimalForwarderTransactorSession) Execute(req IMinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _IMinimalForwarder.Contract.Execute(&_IMinimalForwarder.TransactOpts, req, signature)
}
