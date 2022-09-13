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

// MinimalForwarderMetaData contains all meta data concerning the MinimalForwarder contract.
var MinimalForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b506040518060400160405280601081526020017f4d696e696d616c466f72776172646572000000000000000000000000000000008152506040518060400160405280600581526020017f302e302e3100000000000000000000000000000000000000000000000000000081525060008280519060200120905060008280519060200120905060007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f90508260e081815250508161010081815250504660a08181525050620000e88184846200013760201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505080610120818152505050505050506200024b565b6000838383463060405160200162000154959493929190620001ee565b6040516020818303038152906040528051906020012090509392505050565b6000819050919050565b620001888162000173565b82525050565b6000819050919050565b620001a3816200018e565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001d682620001a9565b9050919050565b620001e881620001c9565b82525050565b600060a0820190506200020560008301886200017d565b6200021460208301876200017d565b6200022360408301866200017d565b62000232606083018562000198565b620002416080830184620001dd565b9695505050505050565b60805160a05160c05160e051610100516101205161127a6200029b6000396000610591015260006105d3015260006105b2015260006104e70152600061053d01526000610566015261127a6000f3fe6080604052600436106100345760003560e01c80632d0335ab1461003957806347153f8214610076578063bf5d3bdb146100a7575b600080fd5b34801561004557600080fd5b50610060600480360381019061005b919061096b565b6100e4565b60405161006d91906109b1565b60405180910390f35b610090600480360381019061008b9190610a55565b61012c565b60405161009e929190610b7c565b60405180910390f35b3480156100b357600080fd5b506100ce60048036038101906100c99190610a55565b6102d7565b6040516100db9190610bac565b60405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000606061013b8585856102ed565b61017a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017190610c4a565b60405180910390fd5b6001856080013561018b9190610c99565b6000808760000160208101906101a1919061096b565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000808660200160208101906101f5919061096b565b73ffffffffffffffffffffffffffffffffffffffff1687606001358860400135898060a001906102259190610cdc565b8b6000016020810190610238919061096b565b60405160200161024a93929190610dc6565b6040516020818303038152906040526040516102669190610e21565b600060405180830381858888f193505050503d80600081146102a4576040519150601f19603f3d011682016040523d82523d6000602084013e6102a9565b606091505b5091509150603f87606001356102bf9190610e67565b5a116102c757fe5b8181935093505050935093915050565b60006102e48484846102ed565b90509392505050565b6000806103f684848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506103e87fdd8f4b70b0f4393e889bd39128a30628a78b61816a9eb8199759e7a349657e4888600001602081019061036f919061096b565b896020016020810190610382919061096b565b8a604001358b606001358c608001358d8060a001906103a19190610cdc565b6040516103af929190610e98565b60405180910390206040516020016103cd9796959493929190610ed9565b604051602081830303815290604052805190602001206104a2565b6104bc90919063ffffffff16565b90508460800135600080876000016020810190610413919061096b565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541480156104985750846000016020810190610469919061096b565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b9150509392505050565b60006104b56104af6104e3565b836105fd565b9050919050565b60008060006104cb8585610630565b915091506104d881610681565b819250505092915050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614801561055f57507f000000000000000000000000000000000000000000000000000000000000000046145b1561058c577f000000000000000000000000000000000000000000000000000000000000000090506105fa565b6105f77f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006107e7565b90505b90565b60008282604051602001610612929190610fc0565b60405160208183030381529060405280519060200120905092915050565b60008060418351036106715760008060006020860151925060408601519150606086015160001a905061066587828585610821565b9450945050505061067a565b60006002915091505b9250929050565b6000600481111561069557610694610ff7565b5b8160048111156106a8576106a7610ff7565b5b03156107e457600160048111156106c2576106c1610ff7565b5b8160048111156106d5576106d4610ff7565b5b03610715576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070c90611072565b60405180910390fd5b6002600481111561072957610728610ff7565b5b81600481111561073c5761073b610ff7565b5b0361077c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610773906110de565b60405180910390fd5b600360048111156107905761078f610ff7565b5b8160048111156107a3576107a2610ff7565b5b036107e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107da90611170565b60405180910390fd5b5b50565b60008383834630604051602001610802959493929190611190565b6040516020818303038152906040528051906020012090509392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c111561085c5760006003915091506108fa565b60006001878787876040516000815260200160405260405161088194939291906111ff565b6020604051602081039080840390855afa1580156108a3573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108f1576000600192509250506108fa565b80600092509250505b94509492505050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109388261090d565b9050919050565b6109488161092d565b811461095357600080fd5b50565b6000813590506109658161093f565b92915050565b60006020828403121561098157610980610903565b5b600061098f84828501610956565b91505092915050565b6000819050919050565b6109ab81610998565b82525050565b60006020820190506109c660008301846109a2565b92915050565b600080fd5b600060c082840312156109e7576109e66109cc565b5b81905092915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610a1557610a146109f0565b5b8235905067ffffffffffffffff811115610a3257610a316109f5565b5b602083019150836001820283011115610a4e57610a4d6109fa565b5b9250929050565b600080600060408486031215610a6e57610a6d610903565b5b600084013567ffffffffffffffff811115610a8c57610a8b610908565b5b610a98868287016109d1565b935050602084013567ffffffffffffffff811115610ab957610ab8610908565b5b610ac5868287016109ff565b92509250509250925092565b60008115159050919050565b610ae681610ad1565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b26578082015181840152602081019050610b0b565b60008484015250505050565b6000601f19601f8301169050919050565b6000610b4e82610aec565b610b588185610af7565b9350610b68818560208601610b08565b610b7181610b32565b840191505092915050565b6000604082019050610b916000830185610add565b8181036020830152610ba38184610b43565b90509392505050565b6000602082019050610bc16000830184610add565b92915050565b600082825260208201905092915050565b7f4d696e696d616c466f727761726465723a207369676e617475726520646f657360008201527f206e6f74206d6174636820726571756573740000000000000000000000000000602082015250565b6000610c34603283610bc7565b9150610c3f82610bd8565b604082019050919050565b60006020820190508181036000830152610c6381610c27565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ca482610998565b9150610caf83610998565b9250828201905080821115610cc757610cc6610c6a565b5b92915050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112610cf957610cf8610ccd565b5b80840192508235915067ffffffffffffffff821115610d1b57610d1a610cd2565b5b602083019250600182023603831315610d3757610d36610cd7565b5b509250929050565b600081905092915050565b82818337600083830152505050565b6000610d658385610d3f565b9350610d72838584610d4a565b82840190509392505050565b60008160601b9050919050565b6000610d9682610d7e565b9050919050565b6000610da882610d8b565b9050919050565b610dc0610dbb8261092d565b610d9d565b82525050565b6000610dd3828587610d59565b9150610ddf8284610daf565b601482019150819050949350505050565b6000610dfb82610aec565b610e058185610d3f565b9350610e15818560208601610b08565b80840191505092915050565b6000610e2d8284610df0565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610e7282610998565b9150610e7d83610998565b925082610e8d57610e8c610e38565b5b828204905092915050565b6000610ea5828486610d59565b91508190509392505050565b6000819050919050565b610ec481610eb1565b82525050565b610ed38161092d565b82525050565b600060e082019050610eee600083018a610ebb565b610efb6020830189610eca565b610f086040830188610eca565b610f1560608301876109a2565b610f2260808301866109a2565b610f2f60a08301856109a2565b610f3c60c0830184610ebb565b98975050505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b6000610f89600283610f48565b9150610f9482610f53565b600282019050919050565b6000819050919050565b610fba610fb582610eb1565b610f9f565b82525050565b6000610fcb82610f7c565b9150610fd78285610fa9565b602082019150610fe78284610fa9565b6020820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b600061105c601883610bc7565b915061106782611026565b602082019050919050565b6000602082019050818103600083015261108b8161104f565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b60006110c8601f83610bc7565b91506110d382611092565b602082019050919050565b600060208201905081810360008301526110f7816110bb565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b600061115a602283610bc7565b9150611165826110fe565b604082019050919050565b600060208201905081810360008301526111898161114d565b9050919050565b600060a0820190506111a56000830188610ebb565b6111b26020830187610ebb565b6111bf6040830186610ebb565b6111cc60608301856109a2565b6111d96080830184610eca565b9695505050505050565b600060ff82169050919050565b6111f9816111e3565b82525050565b60006080820190506112146000830187610ebb565b61122160208301866111f0565b61122e6040830185610ebb565b61123b6060830184610ebb565b9594505050505056fea2646970667358221220b12d9605e84879ab0c76642a1f49995e3146562569ef8134bd9b158297fc5d6b64736f6c63430008100033",
}

// MinimalForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use MinimalForwarderMetaData.ABI instead.
var MinimalForwarderABI = MinimalForwarderMetaData.ABI

// MinimalForwarderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MinimalForwarderMetaData.Bin instead.
var MinimalForwarderBin = MinimalForwarderMetaData.Bin

// DeployMinimalForwarder deploys a new Ethereum contract, binding an instance of MinimalForwarder to it.
func DeployMinimalForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MinimalForwarder, error) {
	parsed, err := MinimalForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MinimalForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MinimalForwarder{MinimalForwarderCaller: MinimalForwarderCaller{contract: contract}, MinimalForwarderTransactor: MinimalForwarderTransactor{contract: contract}, MinimalForwarderFilterer: MinimalForwarderFilterer{contract: contract}}, nil
}

// MinimalForwarder is an auto generated Go binding around an Ethereum contract.
type MinimalForwarder struct {
	MinimalForwarderCaller     // Read-only binding to the contract
	MinimalForwarderTransactor // Write-only binding to the contract
	MinimalForwarderFilterer   // Log filterer for contract events
}

// MinimalForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinimalForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinimalForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinimalForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinimalForwarderSession struct {
	Contract     *MinimalForwarder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinimalForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinimalForwarderCallerSession struct {
	Contract *MinimalForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MinimalForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinimalForwarderTransactorSession struct {
	Contract     *MinimalForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MinimalForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinimalForwarderRaw struct {
	Contract *MinimalForwarder // Generic contract binding to access the raw methods on
}

// MinimalForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinimalForwarderCallerRaw struct {
	Contract *MinimalForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// MinimalForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinimalForwarderTransactorRaw struct {
	Contract *MinimalForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinimalForwarder creates a new instance of MinimalForwarder, bound to a specific deployed contract.
func NewMinimalForwarder(address common.Address, backend bind.ContractBackend) (*MinimalForwarder, error) {
	contract, err := bindMinimalForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarder{MinimalForwarderCaller: MinimalForwarderCaller{contract: contract}, MinimalForwarderTransactor: MinimalForwarderTransactor{contract: contract}, MinimalForwarderFilterer: MinimalForwarderFilterer{contract: contract}}, nil
}

// NewMinimalForwarderCaller creates a new read-only instance of MinimalForwarder, bound to a specific deployed contract.
func NewMinimalForwarderCaller(address common.Address, caller bind.ContractCaller) (*MinimalForwarderCaller, error) {
	contract, err := bindMinimalForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderCaller{contract: contract}, nil
}

// NewMinimalForwarderTransactor creates a new write-only instance of MinimalForwarder, bound to a specific deployed contract.
func NewMinimalForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*MinimalForwarderTransactor, error) {
	contract, err := bindMinimalForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderTransactor{contract: contract}, nil
}

// NewMinimalForwarderFilterer creates a new log filterer instance of MinimalForwarder, bound to a specific deployed contract.
func NewMinimalForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*MinimalForwarderFilterer, error) {
	contract, err := bindMinimalForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderFilterer{contract: contract}, nil
}

// bindMinimalForwarder binds a generic wrapper to an already deployed contract.
func bindMinimalForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinimalForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinimalForwarder *MinimalForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinimalForwarder.Contract.MinimalForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinimalForwarder *MinimalForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinimalForwarder.Contract.MinimalForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinimalForwarder *MinimalForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinimalForwarder.Contract.MinimalForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinimalForwarder *MinimalForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinimalForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinimalForwarder *MinimalForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinimalForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinimalForwarder *MinimalForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinimalForwarder.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_MinimalForwarder *MinimalForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MinimalForwarder.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_MinimalForwarder *MinimalForwarderSession) GetNonce(from common.Address) (*big.Int, error) {
	return _MinimalForwarder.Contract.GetNonce(&_MinimalForwarder.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_MinimalForwarder *MinimalForwarderCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _MinimalForwarder.Contract.GetNonce(&_MinimalForwarder.CallOpts, from)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_MinimalForwarder *MinimalForwarderCaller) Verify(opts *bind.CallOpts, req IMinimalForwarderForwardRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _MinimalForwarder.contract.Call(opts, &out, "verify", req, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_MinimalForwarder *MinimalForwarderSession) Verify(req IMinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _MinimalForwarder.Contract.Verify(&_MinimalForwarder.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_MinimalForwarder *MinimalForwarderCallerSession) Verify(req IMinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _MinimalForwarder.Contract.Verify(&_MinimalForwarder.CallOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MinimalForwarder *MinimalForwarderTransactor) Execute(opts *bind.TransactOpts, req IMinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MinimalForwarder.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MinimalForwarder *MinimalForwarderSession) Execute(req IMinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MinimalForwarder.Contract.Execute(&_MinimalForwarder.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MinimalForwarder *MinimalForwarderTransactorSession) Execute(req IMinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MinimalForwarder.Contract.Execute(&_MinimalForwarder.TransactOpts, req, signature)
}
