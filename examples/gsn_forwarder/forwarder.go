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

// ForwarderMetaData contains all meta data concerning the Forwarder contract.
var ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"domainValue\",\"type\":\"bytes\"}],\"name\":\"DomainRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeStr\",\"type\":\"string\"}],\"name\":\"RequestTypeRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERIC_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"}],\"name\":\"_getEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"registerDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeSuffix\",\"type\":\"string\"}],\"name\":\"registerRequestType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"typeHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060006040518060a00160405280606181526020016200257160619139604051602001620000409190620001f7565b604051602081830303815290604052905062000062816200006960201b60201c565b50620002b1565b600081805190602001209050600160008083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20283604051620000d291906200028d565b60405180910390a25050565b600081905092915050565b7f466f727761726452657175657374280000000000000000000000000000000000600082015250565b600062000121600f83620000de565b91506200012e82620000e9565b600f82019050919050565b600081519050919050565b60005b838110156200016457808201518184015260208101905062000147565b60008484015250505050565b60006200017d8262000139565b620001898185620000de565b93506200019b81856020860162000144565b80840191505092915050565b7f2900000000000000000000000000000000000000000000000000000000000000600082015250565b6000620001df600183620000de565b9150620001ec82620001a7565b600182019050919050565b6000620002048262000112565b915062000212828462000170565b91506200021f82620001d0565b915081905092915050565b600082825260208201905092915050565b6000601f19601f8301169050919050565b6000620002598262000139565b6200026581856200022a565b93506200027781856020860162000144565b62000282816200023b565b840191505092915050565b60006020820190508181036000830152620002a981846200024c565b905092915050565b6122b080620002c16000396000f3fe6080604052600436106100a05760003560e01c8063ad9f99c711610064578063ad9f99c7146101b7578063c3f28abd146101e0578063c722f1771461020b578063d9210be514610248578063e024dc7f14610271578063e2b62f2d146102a2576100a7565b806301ffc9a7146100ac578063066a310c146100e957806321fe98df146101145780632d0335ab146101515780639c7b45921461018e576100a7565b366100a757005b600080fd5b3480156100b857600080fd5b506100d360048036038101906100ce9190611100565b6102df565b6040516100e09190611148565b60405180910390f35b3480156100f557600080fd5b506100fe610359565b60405161010b91906111f3565b60405180910390f35b34801561012057600080fd5b5061013b6004803603810190610136919061124b565b610375565b6040516101489190611148565b60405180910390f35b34801561015d57600080fd5b50610178600480360381019061017391906112d6565b610395565b604051610185919061131c565b60405180910390f35b34801561019a57600080fd5b506101b560048036038101906101b0919061139c565b6103de565b005b3480156101c357600080fd5b506101de60048036038101906101d99190611497565b6104d4565b005b3480156101ec57600080fd5b506101f56104f5565b60405161020291906111f3565b60405180910390f35b34801561021757600080fd5b50610232600480360381019061022d919061124b565b610511565b60405161023f9190611148565b60405180910390f35b34801561025457600080fd5b5061026f600480360381019061026a919061139c565b610531565b005b61028b60048036038101906102869190611497565b61069c565b6040516102999291906115c4565b60405180910390f35b3480156102ae57600080fd5b506102c960048036038101906102c491906115f4565b6108e1565b6040516102d69190611684565b60405180910390f35b60007f25e23e64000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103525750610351826109a4565b5b9050919050565b6040518060a00160405280606181526020016121c86061913981565b60006020528060005260406000206000915054906101000a900460ff1681565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60004690506000604051806080016040528060528152602001612229605291398051906020012086866040516104159291906116e5565b6040518091039020858560405161042d9291906116e5565b6040518091039020843060405160200161044b95949392919061171c565b6040516020818303038152906040529050600081805190602001209050600180600083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8836040516104c39190611684565b60405180910390a250505050505050565b6104dd87610a0e565b6104ec87878787878787610aa8565b50505050505050565b6040518060800160405280605281526020016122296052913981565b60016020528060005260406000206000915054906101000a900460ff1681565b60005b848490508110156106475760008585838181106105545761055361176f565b5b9050013560f81c60f81b90507f2800000000000000000000000000000000000000000000000000000000000000817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141580156105f457507f2900000000000000000000000000000000000000000000000000000000000000817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b610633576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062a906117ea565b60405180910390fd5b50808061063f90611839565b915050610534565b50600084846040518060a00160405280606181526020016121c860619139858560405160200161067b95949392919061197a565b604051602081830303815290604052905061069581610cbd565b5050505050565b600060606106af89898989898989610aa8565b6106b889610d30565b60008960c0013514806106ce5750428960c00135115b61070d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070490611a11565b60405180910390fd5b6000808a604001351461072057619c4090505b60008a8060a001906107329190611a40565b8c600001602081019061074591906112d6565b60405160200161075793929190611aeb565b6040516020818303038152906040529050818b606001356107789190611b15565b6040603f5a6107879190611b49565b6107919190611bd2565b10156107d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c990611c4f565b60405180910390fd5b8a60200160208101906107e591906112d6565b73ffffffffffffffffffffffffffffffffffffffff168b606001358c60400135836040516108139190611ca0565b600060405180830381858888f193505050503d8060008114610851576040519150601f19603f3d011682016040523d82523d6000602084013e610856565b606091505b50809450819550505060008b60400135141580156108745750600047115b156108d3578a600001602081019061088c91906112d6565b73ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156108d1573d6000803e3d6000fd5b505b505097509795505050505050565b6060838560000160208101906108f791906112d6565b73ffffffffffffffffffffffffffffffffffffffff1686602001602081019061092091906112d6565b73ffffffffffffffffffffffffffffffffffffffff168760400135886060013589608001358a8060a001906109559190611a40565b6040516109639291906116e5565b60405180910390208b60c001358a8a60405160200161098b9a99989796959493929190611cf9565b6040516020818303038152906040529050949350505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b806080013560026000836000016020810190610a2a91906112d6565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610aa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9c90611de6565b60405180910390fd5b50565b6001600087815260200190815260200160002060009054906101000a900460ff16610b08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aff90611e52565b60405180910390fd5b60008086815260200190815260200160002060009054906101000a900460ff16610b67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5e90611ebe565b60405180910390fd5b600086610b76898888886108e1565b80519060200120604051602001610b8e929190611f2a565b604051602081830303815290604052805190602001209050600073ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161480610c745750876000016020810190610bef91906112d6565b73ffffffffffffffffffffffffffffffffffffffff16610c5c84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505083610dde90919063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16145b610cb3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610caa90611fad565b60405180910390fd5b5050505050505050565b600081805190602001209050600160008083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20283604051610d2491906111f3565b60405180910390a25050565b806080013560026000836000016020810190610d4c91906112d6565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190610d9790611839565b9190505514610ddb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd290611de6565b60405180910390fd5b50565b6000806000610ded8585610e05565b91509150610dfa81610e56565b819250505092915050565b6000806041835103610e465760008060006020860151925060408601519150606086015160001a9050610e3a87828585610fbc565b94509450505050610e4f565b60006002915091505b9250929050565b60006004811115610e6a57610e69611fcd565b5b816004811115610e7d57610e7c611fcd565b5b0315610fb95760016004811115610e9757610e96611fcd565b5b816004811115610eaa57610ea9611fcd565b5b03610eea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee190612048565b60405180910390fd5b60026004811115610efe57610efd611fcd565b5b816004811115610f1157610f10611fcd565b5b03610f51576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f48906120b4565b60405180910390fd5b60036004811115610f6557610f64611fcd565b5b816004811115610f7857610f77611fcd565b5b03610fb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610faf90612146565b60405180910390fd5b5b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c1115610ff7576000600391509150611095565b60006001878787876040516000815260200160405260405161101c9493929190612182565b6020604051602081039080840390855afa15801561103e573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361108c57600060019250925050611095565b80600092509250505b94509492505050565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6110dd816110a8565b81146110e857600080fd5b50565b6000813590506110fa816110d4565b92915050565b6000602082840312156111165761111561109e565b5b6000611124848285016110eb565b91505092915050565b60008115159050919050565b6111428161112d565b82525050565b600060208201905061115d6000830184611139565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561119d578082015181840152602081019050611182565b60008484015250505050565b6000601f19601f8301169050919050565b60006111c582611163565b6111cf818561116e565b93506111df81856020860161117f565b6111e8816111a9565b840191505092915050565b6000602082019050818103600083015261120d81846111ba565b905092915050565b6000819050919050565b61122881611215565b811461123357600080fd5b50565b6000813590506112458161121f565b92915050565b6000602082840312156112615761126061109e565b5b600061126f84828501611236565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006112a382611278565b9050919050565b6112b381611298565b81146112be57600080fd5b50565b6000813590506112d0816112aa565b92915050565b6000602082840312156112ec576112eb61109e565b5b60006112fa848285016112c1565b91505092915050565b6000819050919050565b61131681611303565b82525050565b6000602082019050611331600083018461130d565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261135c5761135b611337565b5b8235905067ffffffffffffffff8111156113795761137861133c565b5b60208301915083600182028301111561139557611394611341565b5b9250929050565b600080600080604085870312156113b6576113b561109e565b5b600085013567ffffffffffffffff8111156113d4576113d36110a3565b5b6113e087828801611346565b9450945050602085013567ffffffffffffffff811115611403576114026110a3565b5b61140f87828801611346565b925092505092959194509250565b600080fd5b600060e082840312156114385761143761141d565b5b81905092915050565b60008083601f84011261145757611456611337565b5b8235905067ffffffffffffffff8111156114745761147361133c565b5b6020830191508360018202830111156114905761148f611341565b5b9250929050565b600080600080600080600060a0888a0312156114b6576114b561109e565b5b600088013567ffffffffffffffff8111156114d4576114d36110a3565b5b6114e08a828b01611422565b97505060206114f18a828b01611236565b96505060406115028a828b01611236565b955050606088013567ffffffffffffffff811115611523576115226110a3565b5b61152f8a828b01611441565b9450945050608088013567ffffffffffffffff811115611552576115516110a3565b5b61155e8a828b01611441565b925092505092959891949750929550565b600081519050919050565b600082825260208201905092915050565b60006115968261156f565b6115a0818561157a565b93506115b081856020860161117f565b6115b9816111a9565b840191505092915050565b60006040820190506115d96000830185611139565b81810360208301526115eb818461158b565b90509392505050565b6000806000806060858703121561160e5761160d61109e565b5b600085013567ffffffffffffffff81111561162c5761162b6110a3565b5b61163887828801611422565b945050602061164987828801611236565b935050604085013567ffffffffffffffff81111561166a576116696110a3565b5b61167687828801611441565b925092505092959194509250565b6000602082019050818103600083015261169e818461158b565b905092915050565b600081905092915050565b82818337600083830152505050565b60006116cc83856116a6565b93506116d98385846116b1565b82840190509392505050565b60006116f28284866116c0565b91508190509392505050565b61170781611215565b82525050565b61171681611298565b82525050565b600060a08201905061173160008301886116fe565b61173e60208301876116fe565b61174b60408301866116fe565b611758606083018561130d565b611765608083018461170d565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4657443a20696e76616c696420747970656e616d650000000000000000000000600082015250565b60006117d460158361116e565b91506117df8261179e565b602082019050919050565b60006020820190508181036000830152611803816117c7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061184482611303565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118765761187561180a565b5b600182019050919050565b600081905092915050565b60006118988385611881565b93506118a58385846116b1565b82840190509392505050565b7f2800000000000000000000000000000000000000000000000000000000000000600082015250565b60006118e7600183611881565b91506118f2826118b1565b600182019050919050565b600061190882611163565b6119128185611881565b935061192281856020860161117f565b80840191505092915050565b7f2c00000000000000000000000000000000000000000000000000000000000000600082015250565b6000611964600183611881565b915061196f8261192e565b600182019050919050565b600061198782878961188c565b9150611992826118da565b915061199e82866118fd565b91506119a982611957565b91506119b682848661188c565b91508190509695505050505050565b7f4657443a20726571756573742065787069726564000000000000000000000000600082015250565b60006119fb60148361116e565b9150611a06826119c5565b602082019050919050565b60006020820190508181036000830152611a2a816119ee565b9050919050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112611a5d57611a5c611a31565b5b80840192508235915067ffffffffffffffff821115611a7f57611a7e611a36565b5b602083019250600182023603831315611a9b57611a9a611a3b565b5b509250929050565b60008160601b9050919050565b6000611abb82611aa3565b9050919050565b6000611acd82611ab0565b9050919050565b611ae5611ae082611298565b611ac2565b82525050565b6000611af88285876116c0565b9150611b048284611ad4565b601482019150819050949350505050565b6000611b2082611303565b9150611b2b83611303565b9250828201905080821115611b4357611b4261180a565b5b92915050565b6000611b5482611303565b9150611b5f83611303565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611b9857611b9761180a565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611bdd82611303565b9150611be883611303565b925082611bf857611bf7611ba3565b5b828204905092915050565b7f4657443a20696e73756666696369656e74206761730000000000000000000000600082015250565b6000611c3960158361116e565b9150611c4482611c03565b602082019050919050565b60006020820190508181036000830152611c6881611c2c565b9050919050565b6000611c7a8261156f565b611c8481856116a6565b9350611c9481856020860161117f565b80840191505092915050565b6000611cac8284611c6f565b915081905092915050565b6000819050919050565b611cd2611ccd82611215565b611cb7565b82525050565b6000819050919050565b611cf3611cee82611303565b611cd8565b82525050565b6000611d05828d611cc1565b602082019150611d15828c611ce2565b602082019150611d25828b611ce2565b602082019150611d35828a611ce2565b602082019150611d458289611ce2565b602082019150611d558288611ce2565b602082019150611d658287611cc1565b602082019150611d758286611ce2565b602082019150611d868284866116c0565b91508190509b9a5050505050505050505050565b7f4657443a206e6f6e6365206d69736d6174636800000000000000000000000000600082015250565b6000611dd060138361116e565b9150611ddb82611d9a565b602082019050919050565b60006020820190508181036000830152611dff81611dc3565b9050919050565b7f4657443a20756e7265676973746572656420646f6d61696e207365702e000000600082015250565b6000611e3c601d8361116e565b9150611e4782611e06565b602082019050919050565b60006020820190508181036000830152611e6b81611e2f565b9050919050565b7f4657443a20756e72656769737465726564207479706568617368000000000000600082015250565b6000611ea8601a8361116e565b9150611eb382611e72565b602082019050919050565b60006020820190508181036000830152611ed781611e9b565b9050919050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b6000611f14600283611881565b9150611f1f82611ede565b600282019050919050565b6000611f3582611f07565b9150611f418285611cc1565b602082019150611f518284611cc1565b6020820191508190509392505050565b7f4657443a207369676e6174757265206d69736d61746368000000000000000000600082015250565b6000611f9760178361116e565b9150611fa282611f61565b602082019050919050565b60006020820190508181036000830152611fc681611f8a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b600061203260188361116e565b915061203d82611ffc565b602082019050919050565b6000602082019050818103600083015261206181612025565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b600061209e601f8361116e565b91506120a982612068565b602082019050919050565b600060208201905081810360008301526120cd81612091565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b600061213060228361116e565b915061213b826120d4565b604082019050919050565b6000602082019050818103600083015261215f81612123565b9050919050565b600060ff82169050919050565b61217c81612166565b82525050565b600060808201905061219760008301876116fe565b6121a46020830186612173565b6121b160408301856116fe565b6121be60608301846116fe565b9594505050505056fe616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c54696d65454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429a26469706673582212200aec853f62c730eaeb1b4f449e97ff745d435636078e33843fa37d927092063b64736f6c63430008100033616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c54696d65",
}

// ForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ForwarderMetaData.ABI instead.
var ForwarderABI = ForwarderMetaData.ABI

// ForwarderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ForwarderMetaData.Bin instead.
var ForwarderBin = ForwarderMetaData.Bin

// DeployForwarder deploys a new Ethereum contract, binding an instance of Forwarder to it.
func DeployForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Forwarder, error) {
	parsed, err := ForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// Forwarder is an auto generated Go binding around an Ethereum contract.
type Forwarder struct {
	ForwarderCaller     // Read-only binding to the contract
	ForwarderTransactor // Write-only binding to the contract
	ForwarderFilterer   // Log filterer for contract events
}

// ForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwarderSession struct {
	Contract     *Forwarder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwarderCallerSession struct {
	Contract *ForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwarderTransactorSession struct {
	Contract     *ForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwarderRaw struct {
	Contract *Forwarder // Generic contract binding to access the raw methods on
}

// ForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwarderCallerRaw struct {
	Contract *ForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwarderTransactorRaw struct {
	Contract *ForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwarder creates a new instance of Forwarder, bound to a specific deployed contract.
func NewForwarder(address common.Address, backend bind.ContractBackend) (*Forwarder, error) {
	contract, err := bindForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// NewForwarderCaller creates a new read-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderCaller(address common.Address, caller bind.ContractCaller) (*ForwarderCaller, error) {
	contract, err := bindForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderCaller{contract: contract}, nil
}

// NewForwarderTransactor creates a new write-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderTransactor, error) {
	contract, err := bindForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderTransactor{contract: contract}, nil
}

// NewForwarderFilterer creates a new log filterer instance of Forwarder, bound to a specific deployed contract.
func NewForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwarderFilterer, error) {
	contract, err := bindForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwarderFilterer{contract: contract}, nil
}

// bindForwarder binds a generic wrapper to an already deployed contract.
func bindForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.ForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Forwarder *ForwarderCaller) EIP712DOMAINTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "EIP712_DOMAIN_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Forwarder *ForwarderSession) EIP712DOMAINTYPE() (string, error) {
	return _Forwarder.Contract.EIP712DOMAINTYPE(&_Forwarder.CallOpts)
}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Forwarder *ForwarderCallerSession) EIP712DOMAINTYPE() (string, error) {
	return _Forwarder.Contract.EIP712DOMAINTYPE(&_Forwarder.CallOpts)
}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_Forwarder *ForwarderCaller) GENERICPARAMS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "GENERIC_PARAMS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_Forwarder *ForwarderSession) GENERICPARAMS() (string, error) {
	return _Forwarder.Contract.GENERICPARAMS(&_Forwarder.CallOpts)
}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_Forwarder *ForwarderCallerSession) GENERICPARAMS() (string, error) {
	return _Forwarder.Contract.GENERICPARAMS(&_Forwarder.CallOpts)
}

// GetEncoded is a free data retrieval call binding the contract method 0xe2b62f2d.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderCaller) GetEncoded(opts *bind.CallOpts, req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "_getEncoded", req, requestTypeHash, suffixData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncoded is a free data retrieval call binding the contract method 0xe2b62f2d.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderSession) GetEncoded(req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	return _Forwarder.Contract.GetEncoded(&_Forwarder.CallOpts, req, requestTypeHash, suffixData)
}

// GetEncoded is a free data retrieval call binding the contract method 0xe2b62f2d.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderCallerSession) GetEncoded(req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	return _Forwarder.Contract.GetEncoded(&_Forwarder.CallOpts, req, requestTypeHash, suffixData)
}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderCaller) Domains(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "domains", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderSession) Domains(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.Domains(&_Forwarder.CallOpts, arg0)
}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderCallerSession) Domains(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.Domains(&_Forwarder.CallOpts, arg0)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Forwarder *ForwarderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Forwarder *ForwarderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Forwarder.Contract.SupportsInterface(&_Forwarder.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Forwarder *ForwarderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Forwarder.Contract.SupportsInterface(&_Forwarder.CallOpts, interfaceId)
}

// TypeHashes is a free data retrieval call binding the contract method 0x21fe98df.
//
// Solidity: function typeHashes(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderCaller) TypeHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "typeHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TypeHashes is a free data retrieval call binding the contract method 0x21fe98df.
//
// Solidity: function typeHashes(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderSession) TypeHashes(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.TypeHashes(&_Forwarder.CallOpts, arg0)
}

// TypeHashes is a free data retrieval call binding the contract method 0x21fe98df.
//
// Solidity: function typeHashes(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderCallerSession) TypeHashes(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.TypeHashes(&_Forwarder.CallOpts, arg0)
}

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderCaller) Verify(opts *bind.CallOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "verify", req, domainSeparator, requestTypeHash, suffixData, sig)

	if err != nil {
		return err
	}

	return err

}

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderSession) Verify(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderCallerSession) Verify(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactor) Execute(opts *bind.TransactOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "execute", req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactorSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderTransactor) RegisterDomainSeparator(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "registerDomainSeparator", name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterDomainSeparator(&_Forwarder.TransactOpts, name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderTransactorSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterDomainSeparator(&_Forwarder.TransactOpts, name, version)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderTransactor) RegisterRequestType(opts *bind.TransactOpts, typeName string, typeSuffix string) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "registerRequestType", typeName, typeSuffix)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderSession) RegisterRequestType(typeName string, typeSuffix string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterRequestType(&_Forwarder.TransactOpts, typeName, typeSuffix)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderTransactorSession) RegisterRequestType(typeName string, typeSuffix string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterRequestType(&_Forwarder.TransactOpts, typeName, typeSuffix)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Forwarder *ForwarderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Forwarder *ForwarderSession) Receive() (*types.Transaction, error) {
	return _Forwarder.Contract.Receive(&_Forwarder.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Forwarder *ForwarderTransactorSession) Receive() (*types.Transaction, error) {
	return _Forwarder.Contract.Receive(&_Forwarder.TransactOpts)
}

// ForwarderDomainRegisteredIterator is returned from FilterDomainRegistered and is used to iterate over the raw logs and unpacked data for DomainRegistered events raised by the Forwarder contract.
type ForwarderDomainRegisteredIterator struct {
	Event *ForwarderDomainRegistered // Event containing the contract specifics and raw log

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
func (it *ForwarderDomainRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderDomainRegistered)
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
		it.Event = new(ForwarderDomainRegistered)
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
func (it *ForwarderDomainRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderDomainRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderDomainRegistered represents a DomainRegistered event raised by the Forwarder contract.
type ForwarderDomainRegistered struct {
	DomainSeparator [32]byte
	DomainValue     []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDomainRegistered is a free log retrieval operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Forwarder *ForwarderFilterer) FilterDomainRegistered(opts *bind.FilterOpts, domainSeparator [][32]byte) (*ForwarderDomainRegisteredIterator, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderDomainRegisteredIterator{contract: _Forwarder.contract, event: "DomainRegistered", logs: logs, sub: sub}, nil
}

// WatchDomainRegistered is a free log subscription operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Forwarder *ForwarderFilterer) WatchDomainRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderDomainRegistered, domainSeparator [][32]byte) (event.Subscription, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderDomainRegistered)
				if err := _Forwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
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
func (_Forwarder *ForwarderFilterer) ParseDomainRegistered(log types.Log) (*ForwarderDomainRegistered, error) {
	event := new(ForwarderDomainRegistered)
	if err := _Forwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForwarderRequestTypeRegisteredIterator is returned from FilterRequestTypeRegistered and is used to iterate over the raw logs and unpacked data for RequestTypeRegistered events raised by the Forwarder contract.
type ForwarderRequestTypeRegisteredIterator struct {
	Event *ForwarderRequestTypeRegistered // Event containing the contract specifics and raw log

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
func (it *ForwarderRequestTypeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderRequestTypeRegistered)
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
		it.Event = new(ForwarderRequestTypeRegistered)
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
func (it *ForwarderRequestTypeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderRequestTypeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderRequestTypeRegistered represents a RequestTypeRegistered event raised by the Forwarder contract.
type ForwarderRequestTypeRegistered struct {
	TypeHash [32]byte
	TypeStr  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestTypeRegistered is a free log retrieval operation binding the contract event 0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202.
//
// Solidity: event RequestTypeRegistered(bytes32 indexed typeHash, string typeStr)
func (_Forwarder *ForwarderFilterer) FilterRequestTypeRegistered(opts *bind.FilterOpts, typeHash [][32]byte) (*ForwarderRequestTypeRegisteredIterator, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "RequestTypeRegistered", typeHashRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderRequestTypeRegisteredIterator{contract: _Forwarder.contract, event: "RequestTypeRegistered", logs: logs, sub: sub}, nil
}

// WatchRequestTypeRegistered is a free log subscription operation binding the contract event 0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202.
//
// Solidity: event RequestTypeRegistered(bytes32 indexed typeHash, string typeStr)
func (_Forwarder *ForwarderFilterer) WatchRequestTypeRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderRequestTypeRegistered, typeHash [][32]byte) (event.Subscription, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "RequestTypeRegistered", typeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderRequestTypeRegistered)
				if err := _Forwarder.contract.UnpackLog(event, "RequestTypeRegistered", log); err != nil {
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
func (_Forwarder *ForwarderFilterer) ParseRequestTypeRegistered(log types.Log) (*ForwarderRequestTypeRegistered, error) {
	event := new(ForwarderRequestTypeRegistered)
	if err := _Forwarder.contract.UnpackLog(event, "RequestTypeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
