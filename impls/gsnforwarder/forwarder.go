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
	_ = abi.ConvertType
)

// IForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type _duplicate_IForwarderForwardRequest struct {
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
	Bin: "0x60806040523480156200001157600080fd5b5060006040518060a0016040528060618152602001620025b160619139604051602001620000409190620001a4565b604051602081830303815290604052905062000062816200006960201b60201c565b50620002bb565b600081805190602001209050600160008083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20283604051620000d29190620001d7565b60405180910390a25050565b6000620000eb82620001fb565b620000f7818562000206565b93506200010981856020860162000222565b620001148162000258565b840191505092915050565b60006200012c82620001fb565b62000138818562000217565b93506200014a81856020860162000222565b80840191505092915050565b600062000165600f8362000217565b9150620001728262000269565b600f82019050919050565b60006200018c60018362000217565b9150620001998262000292565b600182019050919050565b6000620001b18262000156565b9150620001bf82846200011f565b9150620001cc826200017d565b915081905092915050565b60006020820190508181036000830152620001f38184620000de565b905092915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60005b838110156200024257808201518184015260208101905062000225565b8381111562000252576000848401525b50505050565b6000601f19601f8301169050919050565b7f466f727761726452657175657374280000000000000000000000000000000000600082015250565b7f2900000000000000000000000000000000000000000000000000000000000000600082015250565b6122e680620002cb6000396000f3fe6080604052600436106100a05760003560e01c8063ad9f99c711610064578063ad9f99c7146101b7578063c3f28abd146101e0578063c722f1771461020b578063d9210be514610248578063e024dc7f14610271578063e2b62f2d146102a2576100a7565b806301ffc9a7146100ac578063066a310c146100e957806321fe98df146101145780632d0335ab146101515780639c7b45921461018e576100a7565b366100a757005b600080fd5b3480156100b857600080fd5b506100d360048036038101906100ce919061120c565b6102df565b6040516100e09190611914565b60405180910390f35b3480156100f557600080fd5b506100fe610359565b60405161010b9190611a19565b60405180910390f35b34801561012057600080fd5b5061013b600480360381019061013691906111df565b610375565b6040516101489190611914565b60405180910390f35b34801561015d57600080fd5b50610178600480360381019061017391906111b2565b610395565b6040516101859190611b7b565b60405180910390f35b34801561019a57600080fd5b506101b560048036038101906101b09190611239565b6103de565b005b3480156101c357600080fd5b506101de60048036038101906101d991906112ba565b6104d4565b005b3480156101ec57600080fd5b506101f56104f5565b6040516102029190611a19565b60405180910390f35b34801561021757600080fd5b50610232600480360381019061022d91906111df565b610511565b60405161023f9190611914565b60405180910390f35b34801561025457600080fd5b5061026f600480360381019061026a9190611239565b610531565b005b61028b600480360381019061028691906112ba565b61069c565b60405161029992919061192f565b60405180910390f35b3480156102ae57600080fd5b506102c960048036038101906102c49190611392565b6108e1565b6040516102d691906119f7565b60405180910390f35b60007f25e23e64000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103525750610351826109a4565b5b9050919050565b6040518060a00160405280606181526020016121fe6061913981565b60006020528060005260406000206000915054906101000a900460ff1681565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000469050600060405180608001604052806052815260200161225f60529139805190602001208686604051610415929190611838565b6040518091039020858560405161042d929190611838565b6040518091039020843060405160200161044b95949392919061195f565b6040516020818303038152906040529050600081805190602001209050600180600083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8836040516104c391906119f7565b60405180910390a250505050505050565b6104dd87610a0e565b6104ec87878787878787610aa8565b50505050505050565b60405180608001604052806052815260200161225f6052913981565b60016020528060005260406000206000915054906101000a900460ff1681565b60005b8484905081101561064757600085858381811061055457610553611f03565b5b9050013560f81c60f81b90507f2800000000000000000000000000000000000000000000000000000000000000817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141580156105f457507f2900000000000000000000000000000000000000000000000000000000000000817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b610633576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062a90611b3b565b60405180910390fd5b50808061063f90611df5565b915050610534565b50600084846040518060a00160405280606181526020016121fe60619139858560405160200161067b959493929190611892565b604051602081830303815290604052905061069581610cbd565b5050505050565b600060606106af89898989898989610aa8565b6106b889610d30565b60008960c0013514806106ce5750428960c00135115b61070d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070490611afb565b60405180910390fd5b6000808a604001351461072057619c4090505b60008a8060a001906107329190611b96565b8c600001602081019061074591906111b2565b60405160200161075793929190611851565b6040516020818303038152906040529050818b606001356107789190611c47565b6040603f5a6107879190611cce565b6107919190611c9d565b10156107d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c990611adb565b60405180910390fd5b8a60200160208101906107e591906111b2565b73ffffffffffffffffffffffffffffffffffffffff168b606001358c6040013583604051610813919061187b565b600060405180830381858888f193505050503d8060008114610851576040519150601f19603f3d011682016040523d82523d6000602084013e610856565b606091505b50809450819550505060008b60400135141580156108745750600047115b156108d3578a600001602081019061088c91906111b2565b73ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156108d1573d6000803e3d6000fd5b505b505097509795505050505050565b6060838560000160208101906108f791906111b2565b73ffffffffffffffffffffffffffffffffffffffff1686602001602081019061092091906111b2565b73ffffffffffffffffffffffffffffffffffffffff168760400135886060013589608001358a8060a001906109559190611b96565b604051610963929190611838565b60405180910390208b60c001358a8a60405160200161098b9a99989796959493929190611797565b6040516020818303038152906040529050949350505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b806080013560026000836000016020810190610a2a91906111b2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610aa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9c90611abb565b60405180910390fd5b50565b6001600087815260200190815260200160002060009054906101000a900460ff16610b08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aff90611b1b565b60405180910390fd5b60008086815260200190815260200160002060009054906101000a900460ff16610b67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5e90611b5b565b60405180910390fd5b600086610b76898888886108e1565b80519060200120604051602001610b8e9291906118dd565b604051602081830303815290604052805190602001209050600073ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161480610c745750876000016020810190610bef91906111b2565b73ffffffffffffffffffffffffffffffffffffffff16610c5c84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505083610dde90919063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16145b610cb3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610caa90611a9b565b60405180910390fd5b5050505050505050565b600081805190602001209050600160008083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20283604051610d249190611a19565b60405180910390a25050565b806080013560026000836000016020810190610d4c91906111b2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190610d9790611df5565b9190505514610ddb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd290611abb565b60405180910390fd5b50565b6000806000610ded8585610e05565b91509150610dfa81610e57565b819250505092915050565b600080604183511415610e475760008060006020860151925060408601519150606086015160001a9050610e3b87828585610fc5565b94509450505050610e50565b60006002915091505b9250929050565b60006004811115610e6b57610e6a611ed4565b5b816004811115610e7e57610e7d611ed4565b5b1415610e8957610fc2565b60016004811115610e9d57610e9c611ed4565b5b816004811115610eb057610eaf611ed4565b5b1415610ef1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee890611a3b565b60405180910390fd5b60026004811115610f0557610f04611ed4565b5b816004811115610f1857610f17611ed4565b5b1415610f59576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5090611a5b565b60405180910390fd5b60036004811115610f6d57610f6c611ed4565b5b816004811115610f8057610f7f611ed4565b5b1415610fc1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb890611a7b565b60405180910390fd5b5b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c111561100057600060039150915061109f565b60006001878787876040516000815260200160405260405161102594939291906119b2565b6020604051602081039080840390855afa158015611047573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156110965760006001925092505061109f565b80600092509250505b94509492505050565b6000813590506110b7816121b8565b92915050565b6000813590506110cc816121cf565b92915050565b6000813590506110e1816121e6565b92915050565b60008083601f8401126110fd576110fc611f37565b5b8235905067ffffffffffffffff81111561111a57611119611f32565b5b60208301915083600182028301111561113657611135611f4b565b5b9250929050565b60008083601f84011261115357611152611f37565b5b8235905067ffffffffffffffff8111156111705761116f611f32565b5b60208301915083600182028301111561118c5761118b611f4b565b5b9250929050565b600060e082840312156111a9576111a8611f41565b5b81905092915050565b6000602082840312156111c8576111c7611f5a565b5b60006111d6848285016110a8565b91505092915050565b6000602082840312156111f5576111f4611f5a565b5b6000611203848285016110bd565b91505092915050565b60006020828403121561122257611221611f5a565b5b6000611230848285016110d2565b91505092915050565b6000806000806040858703121561125357611252611f5a565b5b600085013567ffffffffffffffff81111561127157611270611f55565b5b61127d8782880161113d565b9450945050602085013567ffffffffffffffff8111156112a05761129f611f55565b5b6112ac8782880161113d565b925092505092959194509250565b600080600080600080600060a0888a0312156112d9576112d8611f5a565b5b600088013567ffffffffffffffff8111156112f7576112f6611f55565b5b6113038a828b01611193565b97505060206113148a828b016110bd565b96505060406113258a828b016110bd565b955050606088013567ffffffffffffffff81111561134657611345611f55565b5b6113528a828b016110e7565b9450945050608088013567ffffffffffffffff81111561137557611374611f55565b5b6113818a828b016110e7565b925092505092959891949750929550565b600080600080606085870312156113ac576113ab611f5a565b5b600085013567ffffffffffffffff8111156113ca576113c9611f55565b5b6113d687828801611193565b94505060206113e7878288016110bd565b935050604085013567ffffffffffffffff81111561140857611407611f55565b5b611414878288016110e7565b925092505092959194509250565b61142b81611d28565b82525050565b61144261143d82611d28565b611e3e565b82525050565b61145181611d3a565b82525050565b61146081611d46565b82525050565b61147761147282611d46565b611e50565b82525050565b60006114898385611c20565b9350611496838584611db3565b82840190509392505050565b60006114ad82611bf9565b6114b78185611c0f565b93506114c7818560208601611dc2565b6114d081611f5f565b840191505092915050565b60006114e682611bf9565b6114f08185611c20565b9350611500818560208601611dc2565b80840191505092915050565b60006115188385611c3c565b9350611525838584611db3565b82840190509392505050565b600061153c82611c04565b6115468185611c2b565b9350611556818560208601611dc2565b61155f81611f5f565b840191505092915050565b600061157582611c04565b61157f8185611c3c565b935061158f818560208601611dc2565b80840191505092915050565b60006115a8601883611c2b565b91506115b382611f7d565b602082019050919050565b60006115cb601f83611c2b565b91506115d682611fa6565b602082019050919050565b60006115ee600283611c3c565b91506115f982611fcf565b600282019050919050565b6000611611600183611c3c565b915061161c82611ff8565b600182019050919050565b6000611634600183611c3c565b915061163f82612021565b600182019050919050565b6000611657602283611c2b565b91506116628261204a565b604082019050919050565b600061167a601783611c2b565b915061168582612099565b602082019050919050565b600061169d601383611c2b565b91506116a8826120c2565b602082019050919050565b60006116c0601583611c2b565b91506116cb826120eb565b602082019050919050565b60006116e3601483611c2b565b91506116ee82612114565b602082019050919050565b6000611706601d83611c2b565b91506117118261213d565b602082019050919050565b6000611729601583611c2b565b915061173482612166565b602082019050919050565b600061174c601a83611c2b565b91506117578261218f565b602082019050919050565b61176b81611d9c565b82525050565b61178261177d82611d9c565b611e6c565b82525050565b61179181611da6565b82525050565b60006117a3828d611466565b6020820191506117b3828c611771565b6020820191506117c3828b611771565b6020820191506117d3828a611771565b6020820191506117e38289611771565b6020820191506117f38288611771565b6020820191506118038287611466565b6020820191506118138286611771565b60208201915061182482848661147d565b91508190509b9a5050505050505050505050565b600061184582848661147d565b91508190509392505050565b600061185e82858761147d565b915061186a8284611431565b601482019150819050949350505050565b600061188782846114db565b915081905092915050565b600061189f82878961150c565b91506118aa82611627565b91506118b6828661156a565b91506118c182611604565b91506118ce82848661150c565b91508190509695505050505050565b60006118e8826115e1565b91506118f48285611466565b6020820191506119048284611466565b6020820191508190509392505050565b60006020820190506119296000830184611448565b92915050565b60006040820190506119446000830185611448565b818103602083015261195681846114a2565b90509392505050565b600060a0820190506119746000830188611457565b6119816020830187611457565b61198e6040830186611457565b61199b6060830185611762565b6119a86080830184611422565b9695505050505050565b60006080820190506119c76000830187611457565b6119d46020830186611788565b6119e16040830185611457565b6119ee6060830184611457565b95945050505050565b60006020820190508181036000830152611a1181846114a2565b905092915050565b60006020820190508181036000830152611a338184611531565b905092915050565b60006020820190508181036000830152611a548161159b565b9050919050565b60006020820190508181036000830152611a74816115be565b9050919050565b60006020820190508181036000830152611a948161164a565b9050919050565b60006020820190508181036000830152611ab48161166d565b9050919050565b60006020820190508181036000830152611ad481611690565b9050919050565b60006020820190508181036000830152611af4816116b3565b9050919050565b60006020820190508181036000830152611b14816116d6565b9050919050565b60006020820190508181036000830152611b34816116f9565b9050919050565b60006020820190508181036000830152611b548161171c565b9050919050565b60006020820190508181036000830152611b748161173f565b9050919050565b6000602082019050611b906000830184611762565b92915050565b60008083356001602003843603038112611bb357611bb2611f46565b5b80840192508235915067ffffffffffffffff821115611bd557611bd4611f3c565b5b602083019250600182023603831315611bf157611bf0611f50565b5b509250929050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b6000611c5282611d9c565b9150611c5d83611d9c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c9257611c91611e76565b5b828201905092915050565b6000611ca882611d9c565b9150611cb383611d9c565b925082611cc357611cc2611ea5565b5b828204905092915050565b6000611cd982611d9c565b9150611ce483611d9c565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d1d57611d1c611e76565b5b828202905092915050565b6000611d3382611d7c565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611de0578082015181840152602081019050611dc5565b83811115611def576000848401525b50505050565b6000611e0082611d9c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e3357611e32611e76565b5b600182019050919050565b6000611e4982611e5a565b9050919050565b6000819050919050565b6000611e6582611f70565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b7f2c00000000000000000000000000000000000000000000000000000000000000600082015250565b7f2800000000000000000000000000000000000000000000000000000000000000600082015250565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b7f4657443a207369676e6174757265206d69736d61746368000000000000000000600082015250565b7f4657443a206e6f6e6365206d69736d6174636800000000000000000000000000600082015250565b7f4657443a20696e73756666696369656e74206761730000000000000000000000600082015250565b7f4657443a20726571756573742065787069726564000000000000000000000000600082015250565b7f4657443a20756e7265676973746572656420646f6d61696e207365702e000000600082015250565b7f4657443a20696e76616c696420747970656e616d650000000000000000000000600082015250565b7f4657443a20756e72656769737465726564207479706568617368000000000000600082015250565b6121c181611d28565b81146121cc57600080fd5b50565b6121d881611d46565b81146121e357600080fd5b50565b6121ef81611d50565b81146121fa57600080fd5b5056fe616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c54696d65454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429a26469706673582212204afd44b6e82f8317f8eeecaa3980d08cec8fcbcf9bb788e05eedaaa01da9382c64736f6c63430008070033616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c54696d65",
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
	parsed, err := ForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
