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

// MockMetaData contains all meta data concerning the Mock contract.
var MockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedForwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FallbackCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161073e38038061073e833981810160405281019061003291906100d1565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505050506100fe565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061009e82610073565b9050919050565b6100ae81610093565b81146100b957600080fd5b50565b6000815190506100cb816100a5565b92915050565b6000602082840312156100e7576100e661006e565b5b60006100f5848285016100bc565b91505092915050565b608051610625610119600039600061025001526106256000f3fe60806040526004361061002d5760003560e01c8063441a3e701461006f578063572b6c051461009857610034565b3661003457005b34801561004057600080fd5b507ff5c8b1061e6c4978efeae4105f6f09578e04098b17e6aeb3fde64b1be4989b1760405160405180910390a1005b34801561007b57600080fd5b5061009660048036038101906100919190610319565b6100d5565b005b3480156100a457600080fd5b506100bf60048036038101906100ba91906103b7565b61024c565b6040516100cc91906103ff565b60405180910390f35b814711610117576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010e9061049d565b60405180910390fd5b818110610159576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015090610509565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561019f573d6000803e3d6000fd5b506101a86102a4565b73ffffffffffffffffffffffffffffffffffffffff166108fc82846101cd9190610558565b9081150290604051600060405180830381858888f193505050501580156101f8573d6000803e3d6000fd5b507ff341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb5676102226102a4565b32838561022f9190610558565b8460405161024094939291906105aa565b60405180910390a15050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b60006102af3361024c565b156102c357601436033560601c90506102d2565b6102cb6102d6565b90506102d3565b5b90565b600033905090565b600080fd5b6000819050919050565b6102f6816102e3565b811461030157600080fd5b50565b600081359050610313816102ed565b92915050565b600080604083850312156103305761032f6102de565b5b600061033e85828601610304565b925050602061034f85828601610304565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061038482610359565b9050919050565b61039481610379565b811461039f57600080fd5b50565b6000813590506103b18161038b565b92915050565b6000602082840312156103cd576103cc6102de565b5b60006103db848285016103a2565b91505092915050565b60008115159050919050565b6103f9816103e4565b82525050565b600060208201905061041460008301846103f0565b92915050565b600082825260208201905092915050565b7f76616c756520746f207472616e7366657220697320686967686572207468616e60008201527f20636f6e74726163742062616c616e6365000000000000000000000000000000602082015250565b600061048760318361041a565b91506104928261042b565b604082019050919050565b600060208201905081810360008301526104b68161047a565b9050919050565b7f66656520697320686967686572207468616e2076616c75650000000000000000600082015250565b60006104f360188361041a565b91506104fe826104bd565b602082019050919050565b60006020820190508181036000830152610522816104e6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610563826102e3565b915061056e836102e3565b925082820390508181111561058657610585610529565b5b92915050565b61059581610379565b82525050565b6105a4816102e3565b82525050565b60006080820190506105bf600083018761058c565b6105cc602083018661058c565b6105d9604083018561059b565b6105e6606083018461059b565b9594505050505056fea2646970667358221220cf0749a7b7b2ad1e11259682ebd882542b22e65bca52c2f363df73cf1f595b6364736f6c63430008100033",
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
// Solidity: fallback() returns()
func (_Mock *MockTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Mock.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Mock *MockSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Mock.Contract.Fallback(&_Mock.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
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
