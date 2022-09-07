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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedForwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610582380380610582833981810160405281019061003291906100d1565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505050506100fe565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061009e82610073565b9050919050565b6100ae81610093565b81146100b957600080fd5b50565b6000815190506100cb816100a5565b92915050565b6000602082840312156100e7576100e661006e565b5b60006100f5848285016100bc565b91505092915050565b608051610469610119600039600061018901526104696000f3fe60806040526004361061002d5760003560e01c8063441a3e7014610039578063572b6c051461006257610034565b3661003457005b600080fd5b34801561004557600080fd5b50610060600480360381019061005b9190610252565b61009f565b005b34801561006e57600080fd5b50610089600480360381019061008491906102f0565b610185565b6040516100969190610338565b60405180910390f35b8181106100e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100d8906103b0565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610127573d6000803e3d6000fd5b506101306101dd565b73ffffffffffffffffffffffffffffffffffffffff166108fc828461015591906103ff565b9081150290604051600060405180830381858888f19350505050158015610180573d6000803e3d6000fd5b505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b60006101e833610185565b156101fc57601436033560601c905061020b565b61020461020f565b905061020c565b5b90565b600033905090565b600080fd5b6000819050919050565b61022f8161021c565b811461023a57600080fd5b50565b60008135905061024c81610226565b92915050565b6000806040838503121561026957610268610217565b5b60006102778582860161023d565b92505060206102888582860161023d565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102bd82610292565b9050919050565b6102cd816102b2565b81146102d857600080fd5b50565b6000813590506102ea816102c4565b92915050565b60006020828403121561030657610305610217565b5b6000610314848285016102db565b91505092915050565b60008115159050919050565b6103328161031d565b82525050565b600060208201905061034d6000830184610329565b92915050565b600082825260208201905092915050565b7f66656520697320686967686572207468616e2076616c75650000000000000000600082015250565b600061039a601883610353565b91506103a582610364565b602082019050919050565b600060208201905081810360008301526103c98161038d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061040a8261021c565b91506104158361021c565b925082820390508181111561042d5761042c6103d0565b5b9291505056fea2646970667358221220a8e962db8651de1687cbdd2f5eb467c4919375979e4336806f38a72004d550e164736f6c63430008100033",
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
