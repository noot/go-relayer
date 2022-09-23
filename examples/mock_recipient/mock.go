// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock

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

// MockMetaData contains all meta data concerning the Mock contract.
var MockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedForwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FallbackCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161080c38038061080c833981810160405281019061003291906100d1565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505050506100fe565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061009e82610073565b9050919050565b6100ae81610093565b81146100b957600080fd5b50565b6000815190506100cb816100a5565b92915050565b6000602082840312156100e7576100e661006e565b5b60006100f5848285016100bc565b91505092915050565b6080516106f3610119600039600061028c01526106f36000f3fe60806040526004361061002d5760003560e01c8063441a3e7014610062578063572b6c051461008b57610034565b3661003457005b7ff5c8b1061e6c4978efeae4105f6f09578e04098b17e6aeb3fde64b1be4989b1760405160405180910390a1005b34801561006e57600080fd5b5061008960048036038101906100849190610355565b6100c8565b005b34801561009757600080fd5b506100b260048036038101906100ad91906103f3565b610288565b6040516100bf919061043b565b60405180910390f35b6100d133610288565b610110576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610107906104d9565b60405180910390fd5b81471015610153576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014a9061056b565b60405180910390fd5b818110610195576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018c906105d7565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156101db573d6000803e3d6000fd5b506101e46102e0565b73ffffffffffffffffffffffffffffffffffffffff166108fc82846102099190610626565b9081150290604051600060405180830381858888f19350505050158015610234573d6000803e3d6000fd5b507ff341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb56761025e6102e0565b32838561026b9190610626565b8460405161027c9493929190610678565b60405180910390a15050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b60006102eb33610288565b156102ff57601436033560601c905061030e565b610307610312565b905061030f565b5b90565b600033905090565b600080fd5b6000819050919050565b6103328161031f565b811461033d57600080fd5b50565b60008135905061034f81610329565b92915050565b6000806040838503121561036c5761036b61031a565b5b600061037a85828601610340565b925050602061038b85828601610340565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103c082610395565b9050919050565b6103d0816103b5565b81146103db57600080fd5b50565b6000813590506103ed816103c7565b92915050565b6000602082840312156104095761040861031a565b5b6000610417848285016103de565b91505092915050565b60008115159050919050565b61043581610420565b82525050565b6000602082019050610450600083018461042c565b92915050565b600082825260208201905092915050565b7f66756e6374696f6e206d7573742062652063616c6c656420627920666f72776160008201527f7264657200000000000000000000000000000000000000000000000000000000602082015250565b60006104c3602483610456565b91506104ce82610467565b604082019050919050565b600060208201905081810360008301526104f2816104b6565b9050919050565b7f76616c756520746f207472616e7366657220697320686967686572207468616e60008201527f20636f6e74726163742062616c616e6365000000000000000000000000000000602082015250565b6000610555603183610456565b9150610560826104f9565b604082019050919050565b6000602082019050818103600083015261058481610548565b9050919050565b7f66656520697320686967686572207468616e2076616c75650000000000000000600082015250565b60006105c1601883610456565b91506105cc8261058b565b602082019050919050565b600060208201905081810360008301526105f0816105b4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106318261031f565b915061063c8361031f565b9250828203905081811115610654576106536105f7565b5b92915050565b610663816103b5565b82525050565b6106728161031f565b82525050565b600060808201905061068d600083018761065a565b61069a602083018661065a565b6106a76040830185610669565b6106b46060830184610669565b9594505050505056fea2646970667358221220f4e4805969e5f3f629cc1a71da6b870f3e0dc2cd1c98facddb85509decbcee5e64736f6c63430008100033",
}

// MockABI is the input ABI used to generate the binding from.
// Deprecated: Use MockMetaData.ABI instead.
var MockABI = MockMetaData.ABI

// MockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockMetaData.Bin instead.
var MockBin = MockMetaData.Bin

// DeployMock deploys a new Ethereum contract, binding an instance of Mock to it.
func DeployMock(auth *bind.TransactOpts, backend bind.ContractBackend, trustedForwarder common.Address) (common.Address, *types.Transaction, *Mock, error) {
	parsed, err := MockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockBin), backend, trustedForwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mock{MockCaller: MockCaller{contract: contract}, MockTransactor: MockTransactor{contract: contract}, MockFilterer: MockFilterer{contract: contract}}, nil
}

// Mock is an auto generated Go binding around an Ethereum contract.
type Mock struct {
	MockCaller     // Read-only binding to the contract
	MockTransactor // Write-only binding to the contract
	MockFilterer   // Log filterer for contract events
}

// MockCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockSession struct {
	Contract     *Mock             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockCallerSession struct {
	Contract *MockCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTransactorSession struct {
	Contract     *MockTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockRaw struct {
	Contract *Mock // Generic contract binding to access the raw methods on
}

// MockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockCallerRaw struct {
	Contract *MockCaller // Generic read-only contract binding to access the raw methods on
}

// MockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTransactorRaw struct {
	Contract *MockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMock creates a new instance of Mock, bound to a specific deployed contract.
func NewMock(address common.Address, backend bind.ContractBackend) (*Mock, error) {
	contract, err := bindMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mock{MockCaller: MockCaller{contract: contract}, MockTransactor: MockTransactor{contract: contract}, MockFilterer: MockFilterer{contract: contract}}, nil
}

// NewMockCaller creates a new read-only instance of Mock, bound to a specific deployed contract.
func NewMockCaller(address common.Address, caller bind.ContractCaller) (*MockCaller, error) {
	contract, err := bindMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockCaller{contract: contract}, nil
}

// NewMockTransactor creates a new write-only instance of Mock, bound to a specific deployed contract.
func NewMockTransactor(address common.Address, transactor bind.ContractTransactor) (*MockTransactor, error) {
	contract, err := bindMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTransactor{contract: contract}, nil
}

// NewMockFilterer creates a new log filterer instance of Mock, bound to a specific deployed contract.
func NewMockFilterer(address common.Address, filterer bind.ContractFilterer) (*MockFilterer, error) {
	contract, err := bindMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockFilterer{contract: contract}, nil
}

// bindMock binds a generic wrapper to an already deployed contract.
func bindMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mock *MockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mock.Contract.MockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mock *MockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mock.Contract.MockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mock *MockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mock.Contract.MockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mock *MockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mock *MockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mock *MockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mock.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Mock *MockCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Mock *MockSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Mock.Contract.IsTrustedForwarder(&_Mock.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Mock *MockCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Mock.Contract.IsTrustedForwarder(&_Mock.CallOpts, forwarder)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 value, uint256 fee) returns()
func (_Mock *MockTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Mock.contract.Transact(opts, "withdraw", value, fee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 value, uint256 fee) returns()
func (_Mock *MockSession) Withdraw(value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Mock.Contract.Withdraw(&_Mock.TransactOpts, value, fee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 value, uint256 fee) returns()
func (_Mock *MockTransactorSession) Withdraw(value *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Mock.Contract.Withdraw(&_Mock.TransactOpts, value, fee)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Mock *MockTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Mock.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Mock *MockSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Mock.Contract.Fallback(&_Mock.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Mock *MockTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Mock.Contract.Fallback(&_Mock.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Mock *MockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Mock *MockSession) Receive() (*types.Transaction, error) {
	return _Mock.Contract.Receive(&_Mock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Mock *MockTransactorSession) Receive() (*types.Transaction, error) {
	return _Mock.Contract.Receive(&_Mock.TransactOpts)
}

// MockFallbackCalledIterator is returned from FilterFallbackCalled and is used to iterate over the raw logs and unpacked data for FallbackCalled events raised by the Mock contract.
type MockFallbackCalledIterator struct {
	Event *MockFallbackCalled // Event containing the contract specifics and raw log

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
func (it *MockFallbackCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockFallbackCalled)
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
		it.Event = new(MockFallbackCalled)
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
func (it *MockFallbackCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockFallbackCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockFallbackCalled represents a FallbackCalled event raised by the Mock contract.
type MockFallbackCalled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFallbackCalled is a free log retrieval operation binding the contract event 0xf5c8b1061e6c4978efeae4105f6f09578e04098b17e6aeb3fde64b1be4989b17.
//
// Solidity: event FallbackCalled()
func (_Mock *MockFilterer) FilterFallbackCalled(opts *bind.FilterOpts) (*MockFallbackCalledIterator, error) {

	logs, sub, err := _Mock.contract.FilterLogs(opts, "FallbackCalled")
	if err != nil {
		return nil, err
	}
	return &MockFallbackCalledIterator{contract: _Mock.contract, event: "FallbackCalled", logs: logs, sub: sub}, nil
}

// WatchFallbackCalled is a free log subscription operation binding the contract event 0xf5c8b1061e6c4978efeae4105f6f09578e04098b17e6aeb3fde64b1be4989b17.
//
// Solidity: event FallbackCalled()
func (_Mock *MockFilterer) WatchFallbackCalled(opts *bind.WatchOpts, sink chan<- *MockFallbackCalled) (event.Subscription, error) {

	logs, sub, err := _Mock.contract.WatchLogs(opts, "FallbackCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockFallbackCalled)
				if err := _Mock.contract.UnpackLog(event, "FallbackCalled", log); err != nil {
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

// ParseFallbackCalled is a log parse operation binding the contract event 0xf5c8b1061e6c4978efeae4105f6f09578e04098b17e6aeb3fde64b1be4989b17.
//
// Solidity: event FallbackCalled()
func (_Mock *MockFilterer) ParseFallbackCalled(log types.Log) (*MockFallbackCalled, error) {
	event := new(MockFallbackCalled)
	if err := _Mock.contract.UnpackLog(event, "FallbackCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Mock contract.
type MockWithdrawIterator struct {
	Event *MockWithdraw // Event containing the contract specifics and raw log

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
func (it *MockWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockWithdraw)
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
		it.Event = new(MockWithdraw)
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
func (it *MockWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockWithdraw represents a Withdraw event raised by the Mock contract.
type MockWithdraw struct {
	Recipient common.Address
	Relayer   common.Address
	Value     *big.Int
	Fee       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address recipient, address relayer, uint256 value, uint256 fee)
func (_Mock *MockFilterer) FilterWithdraw(opts *bind.FilterOpts) (*MockWithdrawIterator, error) {

	logs, sub, err := _Mock.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &MockWithdrawIterator{contract: _Mock.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address recipient, address relayer, uint256 value, uint256 fee)
func (_Mock *MockFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MockWithdraw) (event.Subscription, error) {

	logs, sub, err := _Mock.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockWithdraw)
				if err := _Mock.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address recipient, address relayer, uint256 value, uint256 fee)
func (_Mock *MockFilterer) ParseWithdraw(log types.Log) (*MockWithdraw, error) {
	event := new(MockWithdraw)
	if err := _Mock.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
