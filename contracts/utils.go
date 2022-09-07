package contracts

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// TODO: make this configurable
	forwardRequestTypehash = crypto.Keccak256Hash([]byte("ForwardRequest(address from,address to,uint256 value,uint256 gas,uint256 nonce,bytes data)"))
)

func padBytesLeft(in []byte, n int) []byte {
	if len(in) > n {
		return in // error probably
	}

	out := make([]byte, n-len(in))
	return append(out, in...)
}

// GetFunctionSignature returns the 4-byte function signature of a Solidity function.
func GetFunctionSignature(fn string) []byte {
	hash := crypto.Keccak256Hash([]byte(fn))
	return hash[:4]
}

// address = forwarder contract address
func getEIP712DomainSeparator(name, version []byte, chainID *big.Int, address ethcommon.Address) ([32]byte, error) {
	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create bytes32 type: %w", err)
	}

	addressTy, err := abi.NewType("address", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create address type: %w", err)
	}

	eip712DomainHash := crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
	var chainIDArr [32]byte
	copy(chainIDArr[:], padBytesLeft(chainID.Bytes(), 32))

	args := &abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: addressTy,
		},
	}
	domainSeparatorPreimage, err := args.Pack(
		eip712DomainHash,
		crypto.Keccak256Hash(name),
		crypto.Keccak256Hash(version),
		chainIDArr,
		address,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create domainSeparatorPreimage: %w", err)
	}

	return crypto.Keccak256Hash(domainSeparatorPreimage), nil
}

// GetForwardRequestDigestToSign returns a 32-byte digest for signing
func GetForwardRequestDigestToSign(req *MinimalForwarderForwardRequest, chainID *big.Int,
	forwarderAddress ethcommon.Address) ([32]byte, error) {
	domainSeparator, err := getEIP712DomainSeparator(
		[]byte("MinimalForwarder"), // TODO: make configurable
		[]byte("0.0.1"),
		chainID,
		forwarderAddress,
	)
	if err != nil {
		return [32]byte{}, err
	}

	uint256Ty, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create uint256 type: %w", err)
	}

	addressTy, err := abi.NewType("address", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create address type: %w", err)
	}

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create bytes32 type: %w", err)
	}

	hashedData := crypto.Keccak256Hash(req.Data)

	args := &abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: addressTy,
		},
		{
			Type: addressTy,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: bytes32Ty,
		},
	}
	hashPreimage, err := args.Pack(
		forwardRequestTypehash,
		req.From,
		req.To,
		req.Value,
		req.Gas,
		req.Nonce,
		hashedData,
	)
	if err != nil {
		return [32]byte{}, err
	}

	hash := crypto.Keccak256Hash(hashPreimage)

	// from ECDSA.toTypedDataHash
	prefix, err := hex.DecodeString("1901")
	if err != nil {
		return [32]byte{}, err
	}

	digestPreimage := append(append(prefix, domainSeparator[:]...), hash[:]...)
	return crypto.Keccak256Hash(digestPreimage), nil
}
