// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IForwarderMetaData contains all meta data concerning the IForwarder contract.
var IForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use IForwarderMetaData.ABI instead.
var IForwarderABI = IForwarderMetaData.ABI

// IForwarder is an auto generated Go binding around an Ethereum contract.
type IForwarder struct {
	IForwarderCaller     // Read-only binding to the contract
	IForwarderTransactor // Write-only binding to the contract
	IForwarderFilterer   // Log filterer for contract events
}

// IForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IForwarderSession struct {
	Contract     *IForwarder       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IForwarderCallerSession struct {
	Contract *IForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IForwarderTransactorSession struct {
	Contract     *IForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IForwarderRaw struct {
	Contract *IForwarder // Generic contract binding to access the raw methods on
}

// IForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IForwarderCallerRaw struct {
	Contract *IForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// IForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IForwarderTransactorRaw struct {
	Contract *IForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIForwarder creates a new instance of IForwarder, bound to a specific deployed contract.
func NewIForwarder(address common.Address, backend bind.ContractBackend) (*IForwarder, error) {
	contract, err := bindIForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IForwarder{IForwarderCaller: IForwarderCaller{contract: contract}, IForwarderTransactor: IForwarderTransactor{contract: contract}, IForwarderFilterer: IForwarderFilterer{contract: contract}}, nil
}

// NewIForwarderCaller creates a new read-only instance of IForwarder, bound to a specific deployed contract.
func NewIForwarderCaller(address common.Address, caller bind.ContractCaller) (*IForwarderCaller, error) {
	contract, err := bindIForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IForwarderCaller{contract: contract}, nil
}

// NewIForwarderTransactor creates a new write-only instance of IForwarder, bound to a specific deployed contract.
func NewIForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*IForwarderTransactor, error) {
	contract, err := bindIForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IForwarderTransactor{contract: contract}, nil
}

// NewIForwarderFilterer creates a new log filterer instance of IForwarder, bound to a specific deployed contract.
func NewIForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*IForwarderFilterer, error) {
	contract, err := bindIForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IForwarderFilterer{contract: contract}, nil
}

// bindIForwarder binds a generic wrapper to an already deployed contract.
func bindIForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IForwarder *IForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IForwarder.Contract.IForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IForwarder *IForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IForwarder.Contract.IForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IForwarder *IForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IForwarder.Contract.IForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IForwarder *IForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IForwarder *IForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IForwarder *IForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IForwarder.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_IForwarder *IForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IForwarder.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_IForwarder *IForwarderSession) GetNonce(from common.Address) (*big.Int, error) {
	return _IForwarder.Contract.GetNonce(&_IForwarder.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_IForwarder *IForwarderCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _IForwarder.Contract.GetNonce(&_IForwarder.CallOpts, from)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_IForwarder *IForwarderCaller) Verify(opts *bind.CallOpts, req IForwarderForwardRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _IForwarder.contract.Call(opts, &out, "verify", req, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_IForwarder *IForwarderSession) Verify(req IForwarderForwardRequest, signature []byte) (bool, error) {
	return _IForwarder.Contract.Verify(&_IForwarder.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_IForwarder *IForwarderCallerSession) Verify(req IForwarderForwardRequest, signature []byte) (bool, error) {
	return _IForwarder.Contract.Verify(&_IForwarder.CallOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_IForwarder *IForwarderTransactor) Execute(opts *bind.TransactOpts, req IForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _IForwarder.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_IForwarder *IForwarderSession) Execute(req IForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _IForwarder.Contract.Execute(&_IForwarder.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_IForwarder *IForwarderTransactorSession) Execute(req IForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _IForwarder.Contract.Execute(&_IForwarder.TransactOpts, req, signature)
}
