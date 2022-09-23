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
func GetForwardRequestDigestToSign(
	req ForwardRequest,
	name,
	version string,
	chainID *big.Int,
	forwarderAddress ethcommon.Address,
) ([32]byte, error) {
	domainSeparator, err := GetEIP712DomainSeparator(
		name,
		version,
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

// GetEIP712DomainSeparator ...
// Note: address = forwarder contract address
func GetEIP712DomainSeparator(name, version string, chainID *big.Int, address ethcommon.Address) ([32]byte, error) {
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
		crypto.Keccak256Hash([]byte(name)),
		crypto.Keccak256Hash([]byte(version)),
		chainIDArr,
		address,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create domainSeparatorPreimage: %w", err)
	}

	return crypto.Keccak256Hash(domainSeparatorPreimage), nil
}

func padBytesLeft(in []byte, n int) []byte {
	if len(in) > n {
		return in // error probably
	}

	out := make([]byte, n-len(in))
	return append(out, in...)
}
