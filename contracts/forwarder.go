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

// MinimalForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type MinimalForwarderForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

// MinimalForwarderMetaData contains all meta data concerning the MinimalForwarder contract.
var MinimalForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b506040518060400160405280601081526020017f4d696e696d616c466f72776172646572000000000000000000000000000000008152506040518060400160405280600581526020017f302e302e3100000000000000000000000000000000000000000000000000000081525060008280519060200120905060008280519060200120905060007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f90508260e081815250508161010081815250504660a08181525050620000e88184846200013760201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505080610120818152505050505050506200024b565b6000838383463060405160200162000154959493929190620001ee565b6040516020818303038152906040528051906020012090509392505050565b6000819050919050565b620001888162000173565b82525050565b6000819050919050565b620001a3816200018e565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001d682620001a9565b9050919050565b620001e881620001c9565b82525050565b600060a0820190506200020560008301886200017d565b6200021460208301876200017d565b6200022360408301866200017d565b62000232606083018562000198565b620002416080830184620001dd565b9695505050505050565b60805160a05160c05160e05161010051610120516112646200029b600039600061057b015260006105bd0152600061059c015260006104d1015260006105270152600061055001526112646000f3fe6080604052600436106100345760003560e01c80632d0335ab1461003957806347153f8214610076578063bf5d3bdb146100a7575b600080fd5b34801561004557600080fd5b50610060600480360381019061005b9190610955565b6100e4565b60405161006d919061099b565b60405180910390f35b610090600480360381019061008b9190610a3f565b61012c565b60405161009e929190610b66565b60405180910390f35b3480156100b357600080fd5b506100ce60048036038101906100c99190610a3f565b6102d7565b6040516100db9190610b96565b60405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000606061013b8585856102d7565b61017a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017190610c34565b60405180910390fd5b6001856080013561018b9190610c83565b6000808760000160208101906101a19190610955565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000808660200160208101906101f59190610955565b73ffffffffffffffffffffffffffffffffffffffff1687606001358860400135898060a001906102259190610cc6565b8b60000160208101906102389190610955565b60405160200161024a93929190610db0565b6040516020818303038152906040526040516102669190610e0b565b600060405180830381858888f193505050503d80600081146102a4576040519150601f19603f3d011682016040523d82523d6000602084013e6102a9565b606091505b5091509150603f87606001356102bf9190610e51565b5a116102c757fe5b8181935093505050935093915050565b6000806103e084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506103d27fdd8f4b70b0f4393e889bd39128a30628a78b61816a9eb8199759e7a349657e488860000160208101906103599190610955565b89602001602081019061036c9190610955565b8a604001358b606001358c608001358d8060a0019061038b9190610cc6565b604051610399929190610e82565b60405180910390206040516020016103b79796959493929190610ec3565b6040516020818303038152906040528051906020012061048c565b6104a690919063ffffffff16565b905084608001356000808760000160208101906103fd9190610955565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414801561048257508460000160208101906104539190610955565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b9150509392505050565b600061049f6104996104cd565b836105e7565b9050919050565b60008060006104b5858561061a565b915091506104c28161066b565b819250505092915050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614801561054957507f000000000000000000000000000000000000000000000000000000000000000046145b15610576577f000000000000000000000000000000000000000000000000000000000000000090506105e4565b6105e17f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006107d1565b90505b90565b600082826040516020016105fc929190610faa565b60405160208183030381529060405280519060200120905092915050565b600080604183510361065b5760008060006020860151925060408601519150606086015160001a905061064f8782858561080b565b94509450505050610664565b60006002915091505b9250929050565b6000600481111561067f5761067e610fe1565b5b81600481111561069257610691610fe1565b5b03156107ce57600160048111156106ac576106ab610fe1565b5b8160048111156106bf576106be610fe1565b5b036106ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f69061105c565b60405180910390fd5b6002600481111561071357610712610fe1565b5b81600481111561072657610725610fe1565b5b03610766576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075d906110c8565b60405180910390fd5b6003600481111561077a57610779610fe1565b5b81600481111561078d5761078c610fe1565b5b036107cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c49061115a565b60405180910390fd5b5b50565b600083838346306040516020016107ec95949392919061117a565b6040516020818303038152906040528051906020012090509392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c11156108465760006003915091506108e4565b60006001878787876040516000815260200160405260405161086b94939291906111e9565b6020604051602081039080840390855afa15801561088d573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108db576000600192509250506108e4565b80600092509250505b94509492505050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610922826108f7565b9050919050565b61093281610917565b811461093d57600080fd5b50565b60008135905061094f81610929565b92915050565b60006020828403121561096b5761096a6108ed565b5b600061097984828501610940565b91505092915050565b6000819050919050565b61099581610982565b82525050565b60006020820190506109b0600083018461098c565b92915050565b600080fd5b600060c082840312156109d1576109d06109b6565b5b81905092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126109ff576109fe6109da565b5b8235905067ffffffffffffffff811115610a1c57610a1b6109df565b5b602083019150836001820283011115610a3857610a376109e4565b5b9250929050565b600080600060408486031215610a5857610a576108ed565b5b600084013567ffffffffffffffff811115610a7657610a756108f2565b5b610a82868287016109bb565b935050602084013567ffffffffffffffff811115610aa357610aa26108f2565b5b610aaf868287016109e9565b92509250509250925092565b60008115159050919050565b610ad081610abb565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b10578082015181840152602081019050610af5565b60008484015250505050565b6000601f19601f8301169050919050565b6000610b3882610ad6565b610b428185610ae1565b9350610b52818560208601610af2565b610b5b81610b1c565b840191505092915050565b6000604082019050610b7b6000830185610ac7565b8181036020830152610b8d8184610b2d565b90509392505050565b6000602082019050610bab6000830184610ac7565b92915050565b600082825260208201905092915050565b7f4d696e696d616c466f727761726465723a207369676e617475726520646f657360008201527f206e6f74206d6174636820726571756573740000000000000000000000000000602082015250565b6000610c1e603283610bb1565b9150610c2982610bc2565b604082019050919050565b60006020820190508181036000830152610c4d81610c11565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c8e82610982565b9150610c9983610982565b9250828201905080821115610cb157610cb0610c54565b5b92915050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112610ce357610ce2610cb7565b5b80840192508235915067ffffffffffffffff821115610d0557610d04610cbc565b5b602083019250600182023603831315610d2157610d20610cc1565b5b509250929050565b600081905092915050565b82818337600083830152505050565b6000610d4f8385610d29565b9350610d5c838584610d34565b82840190509392505050565b60008160601b9050919050565b6000610d8082610d68565b9050919050565b6000610d9282610d75565b9050919050565b610daa610da582610917565b610d87565b82525050565b6000610dbd828587610d43565b9150610dc98284610d99565b601482019150819050949350505050565b6000610de582610ad6565b610def8185610d29565b9350610dff818560208601610af2565b80840191505092915050565b6000610e178284610dda565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610e5c82610982565b9150610e6783610982565b925082610e7757610e76610e22565b5b828204905092915050565b6000610e8f828486610d43565b91508190509392505050565b6000819050919050565b610eae81610e9b565b82525050565b610ebd81610917565b82525050565b600060e082019050610ed8600083018a610ea5565b610ee56020830189610eb4565b610ef26040830188610eb4565b610eff606083018761098c565b610f0c608083018661098c565b610f1960a083018561098c565b610f2660c0830184610ea5565b98975050505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b6000610f73600283610f32565b9150610f7e82610f3d565b600282019050919050565b6000819050919050565b610fa4610f9f82610e9b565b610f89565b82525050565b6000610fb582610f66565b9150610fc18285610f93565b602082019150610fd18284610f93565b6020820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b6000611046601883610bb1565b915061105182611010565b602082019050919050565b6000602082019050818103600083015261107581611039565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b60006110b2601f83610bb1565b91506110bd8261107c565b602082019050919050565b600060208201905081810360008301526110e1816110a5565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b6000611144602283610bb1565b915061114f826110e8565b604082019050919050565b6000602082019050818103600083015261117381611137565b9050919050565b600060a08201905061118f6000830188610ea5565b61119c6020830187610ea5565b6111a96040830186610ea5565b6111b6606083018561098c565b6111c36080830184610eb4565b9695505050505050565b600060ff82169050919050565b6111e3816111cd565b82525050565b60006080820190506111fe6000830187610ea5565b61120b60208301866111da565b6112186040830185610ea5565b6112256060830184610ea5565b9594505050505056fea2646970667358221220e877ce1bf70f6ac87ccdc70d159fd1af508d286de0526f252e88e1d0cf36072a64736f6c63430008100033",
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
func (_MinimalForwarder *MinimalForwarderCaller) Verify(opts *bind.CallOpts, req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
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
func (_MinimalForwarder *MinimalForwarderSession) Verify(req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _MinimalForwarder.Contract.Verify(&_MinimalForwarder.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_MinimalForwarder *MinimalForwarderCallerSession) Verify(req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _MinimalForwarder.Contract.Verify(&_MinimalForwarder.CallOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MinimalForwarder *MinimalForwarderTransactor) Execute(opts *bind.TransactOpts, req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MinimalForwarder.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MinimalForwarder *MinimalForwarderSession) Execute(req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MinimalForwarder.Contract.Execute(&_MinimalForwarder.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MinimalForwarder *MinimalForwarderTransactorSession) Execute(req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MinimalForwarder.Contract.Execute(&_MinimalForwarder.TransactOpts, req, signature)
}
