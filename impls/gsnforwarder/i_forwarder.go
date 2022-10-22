// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gsnforwarder

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

// IForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type IForwarderForwardRequest struct {
	From           common.Address
	To             common.Address
	Value          *big.Int
	Gas            *big.Int
	Nonce          *big.Int
	Data           []byte
	ValidUntilTime *big.Int
}

// IForwarderMetaData contains all meta data concerning the IForwarder contract.
var IForwarderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"domainValue\",\"type\":\"bytes\"}],\"name\":\"DomainRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeStr\",\"type\":\"string\"}],\"name\":\"RequestTypeRegistered\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"forwardRequest\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"registerDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeSuffix\",\"type\":\"string\"}],\"name\":\"registerRequestType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"forwardRequest\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IForwarder *IForwarderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IForwarder.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IForwarder *IForwarderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IForwarder.Contract.SupportsInterface(&_IForwarder.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IForwarder *IForwarderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IForwarder.Contract.SupportsInterface(&_IForwarder.CallOpts, interfaceId)
}

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) forwardRequest, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes signature) view returns()
func (_IForwarder *IForwarderCaller) Verify(opts *bind.CallOpts, forwardRequest IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, signature []byte) error {
	var out []interface{}
	err := _IForwarder.contract.Call(opts, &out, "verify", forwardRequest, domainSeparator, requestTypeHash, suffixData, signature)

	if err != nil {
		return err
	}

	return err

}

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) forwardRequest, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes signature) view returns()
func (_IForwarder *IForwarderSession) Verify(forwardRequest IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, signature []byte) error {
	return _IForwarder.Contract.Verify(&_IForwarder.CallOpts, forwardRequest, domainSeparator, requestTypeHash, suffixData, signature)
}

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) forwardRequest, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes signature) view returns()
func (_IForwarder *IForwarderCallerSession) Verify(forwardRequest IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, signature []byte) error {
	return _IForwarder.Contract.Verify(&_IForwarder.CallOpts, forwardRequest, domainSeparator, requestTypeHash, suffixData, signature)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) forwardRequest, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes signature) payable returns(bool success, bytes ret)
func (_IForwarder *IForwarderTransactor) Execute(opts *bind.TransactOpts, forwardRequest IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, signature []byte) (*types.Transaction, error) {
	return _IForwarder.contract.Transact(opts, "execute", forwardRequest, domainSeparator, requestTypeHash, suffixData, signature)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) forwardRequest, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes signature) payable returns(bool success, bytes ret)
func (_IForwarder *IForwarderSession) Execute(forwardRequest IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, signature []byte) (*types.Transaction, error) {
	return _IForwarder.Contract.Execute(&_IForwarder.TransactOpts, forwardRequest, domainSeparator, requestTypeHash, suffixData, signature)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) forwardRequest, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes signature) payable returns(bool success, bytes ret)
func (_IForwarder *IForwarderTransactorSession) Execute(forwardRequest IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, signature []byte) (*types.Transaction, error) {
	return _IForwarder.Contract.Execute(&_IForwarder.TransactOpts, forwardRequest, domainSeparator, requestTypeHash, suffixData, signature)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_IForwarder *IForwarderTransactor) RegisterDomainSeparator(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	return _IForwarder.contract.Transact(opts, "registerDomainSeparator", name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_IForwarder *IForwarderSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _IForwarder.Contract.RegisterDomainSeparator(&_IForwarder.TransactOpts, name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_IForwarder *IForwarderTransactorSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _IForwarder.Contract.RegisterDomainSeparator(&_IForwarder.TransactOpts, name, version)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_IForwarder *IForwarderTransactor) RegisterRequestType(opts *bind.TransactOpts, typeName string, typeSuffix string) (*types.Transaction, error) {
	return _IForwarder.contract.Transact(opts, "registerRequestType", typeName, typeSuffix)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_IForwarder *IForwarderSession) RegisterRequestType(typeName string, typeSuffix string) (*types.Transaction, error) {
	return _IForwarder.Contract.RegisterRequestType(&_IForwarder.TransactOpts, typeName, typeSuffix)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_IForwarder *IForwarderTransactorSession) RegisterRequestType(typeName string, typeSuffix string) (*types.Transaction, error) {
	return _IForwarder.Contract.RegisterRequestType(&_IForwarder.TransactOpts, typeName, typeSuffix)
}

// IForwarderDomainRegisteredIterator is returned from FilterDomainRegistered and is used to iterate over the raw logs and unpacked data for DomainRegistered events raised by the IForwarder contract.
type IForwarderDomainRegisteredIterator struct {
	Event *IForwarderDomainRegistered // Event containing the contract specifics and raw log

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
func (it *IForwarderDomainRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IForwarderDomainRegistered)
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
		it.Event = new(IForwarderDomainRegistered)
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
func (it *IForwarderDomainRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IForwarderDomainRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IForwarderDomainRegistered represents a DomainRegistered event raised by the IForwarder contract.
type IForwarderDomainRegistered struct {
	DomainSeparator [32]byte
	DomainValue     []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDomainRegistered is a free log retrieval operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_IForwarder *IForwarderFilterer) FilterDomainRegistered(opts *bind.FilterOpts, domainSeparator [][32]byte) (*IForwarderDomainRegisteredIterator, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _IForwarder.contract.FilterLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return &IForwarderDomainRegisteredIterator{contract: _IForwarder.contract, event: "DomainRegistered", logs: logs, sub: sub}, nil
}

// WatchDomainRegistered is a free log subscription operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_IForwarder *IForwarderFilterer) WatchDomainRegistered(opts *bind.WatchOpts, sink chan<- *IForwarderDomainRegistered, domainSeparator [][32]byte) (event.Subscription, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _IForwarder.contract.WatchLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IForwarderDomainRegistered)
				if err := _IForwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
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

// ParseDomainRegistered is a log parse operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_IForwarder *IForwarderFilterer) ParseDomainRegistered(log types.Log) (*IForwarderDomainRegistered, error) {
	event := new(IForwarderDomainRegistered)
	if err := _IForwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IForwarderRequestTypeRegisteredIterator is returned from FilterRequestTypeRegistered and is used to iterate over the raw logs and unpacked data for RequestTypeRegistered events raised by the IForwarder contract.
type IForwarderRequestTypeRegisteredIterator struct {
	Event *IForwarderRequestTypeRegistered // Event containing the contract specifics and raw log

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
func (it *IForwarderRequestTypeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IForwarderRequestTypeRegistered)
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
		it.Event = new(IForwarderRequestTypeRegistered)
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
func (it *IForwarderRequestTypeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IForwarderRequestTypeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IForwarderRequestTypeRegistered represents a RequestTypeRegistered event raised by the IForwarder contract.
type IForwarderRequestTypeRegistered struct {
	TypeHash [32]byte
	TypeStr  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestTypeRegistered is a free log retrieval operation binding the contract event 0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202.
//
// Solidity: event RequestTypeRegistered(bytes32 indexed typeHash, string typeStr)
func (_IForwarder *IForwarderFilterer) FilterRequestTypeRegistered(opts *bind.FilterOpts, typeHash [][32]byte) (*IForwarderRequestTypeRegisteredIterator, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	logs, sub, err := _IForwarder.contract.FilterLogs(opts, "RequestTypeRegistered", typeHashRule)
	if err != nil {
		return nil, err
	}
	return &IForwarderRequestTypeRegisteredIterator{contract: _IForwarder.contract, event: "RequestTypeRegistered", logs: logs, sub: sub}, nil
}

// WatchRequestTypeRegistered is a free log subscription operation binding the contract event 0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202.
//
// Solidity: event RequestTypeRegistered(bytes32 indexed typeHash, string typeStr)
func (_IForwarder *IForwarderFilterer) WatchRequestTypeRegistered(opts *bind.WatchOpts, sink chan<- *IForwarderRequestTypeRegistered, typeHash [][32]byte) (event.Subscription, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	logs, sub, err := _IForwarder.contract.WatchLogs(opts, "RequestTypeRegistered", typeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IForwarderRequestTypeRegistered)
				if err := _IForwarder.contract.UnpackLog(event, "RequestTypeRegistered", log); err != nil {
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

// ParseRequestTypeRegistered is a log parse operation binding the contract event 0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202.
//
// Solidity: event RequestTypeRegistered(bytes32 indexed typeHash, string typeStr)
func (_IForwarder *IForwarderFilterer) ParseRequestTypeRegistered(log types.Log) (*IForwarderRequestTypeRegistered, error) {
	event := new(IForwarderRequestTypeRegistered)
	if err := _IForwarder.contract.UnpackLog(event, "RequestTypeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
