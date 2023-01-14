package gsnforwarder

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/athanorlabs/go-relayer/common"
)

var _ common.ForwardRequest = &IForwarderForwardRequest{}

var (
	DefaultName            = "Forwarder"
	DefaultVersion         = "0.0.1"
	ForwardRequestTypehash = crypto.Keccak256Hash([]byte("ForwardRequest(address from,address to,uint256 value,uint256 gas,uint256 nonce,bytes data,uint256 validUntilTime)"))
)

func (r *IForwarderForwardRequest) FromSubmitTransactionRequest(
	req *common.SubmitTransactionRequest,
) {
	r.From = req.From
	r.To = req.To
	r.Value = req.Value
	r.Gas = req.Gas
	r.Nonce = req.Nonce
	r.Data = req.Data
	r.ValidUntilTime = req.ValidUntilTime
}

func (r *IForwarderForwardRequest) Pack(suffixData []byte) ([]byte, error) {
	uint256Ty, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create uint256 type: %w", err)
	}

	addressTy, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create address type: %w", err)
	}

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bytes32 type: %w", err)
	}

	// bytesTy, err := abi.NewType("bytes", "", nil)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create bytes type: %w", err)
	// }

	if r.ValidUntilTime == nil {
		r.ValidUntilTime = big.NewInt(0)
	}

	hashedData := crypto.Keccak256Hash(r.Data)

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
		{
			Type: uint256Ty,
		},
		// {
		// 	Type: bytesTy,
		// },
	}
	packed, err := args.Pack(
		ForwardRequestTypehash,
		r.From,
		r.To,
		r.Value,
		r.Gas,
		r.Nonce,
		hashedData,
		r.ValidUntilTime,
		//suffixData,
	)
	if err != nil {
		return nil, err
	}

	return packed, nil
}
