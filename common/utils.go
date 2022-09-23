package common

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetFunctionSignature returns the 4-byte function signature of a Solidity function.
func GetFunctionSignature(fn string) []byte {
	hash := crypto.Keccak256Hash([]byte(fn))
	return hash[:4]
}

// GetForwardRequestDigestToSign returns a 32-byte digest for signing
func GetForwardRequestDigestToSign(req ForwardRequest, chainID *big.Int,
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

	hashPreimage, err := req.Pack()
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

func padBytesLeft(in []byte, n int) []byte {
	if len(in) > n {
		return in // error probably
	}

	out := make([]byte, n-len(in))
	return append(out, in...)
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
