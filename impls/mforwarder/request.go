package mforwarder

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/athanorlabs/go-relayer/common"
)

var _ common.ForwardRequest = &IMinimalForwarderForwardRequest{}

var (
	forwardRequestTypehash = crypto.Keccak256Hash(
		[]byte("ForwardRequest(address from,address to,uint256 value,uint256 gas,uint256 nonce,bytes data)"),
	)
)

// FromSubmitTransactionRequest returns a new SubmitTransactionRequest from the
// given IMinimalForwarderForwardRequest.
func (r *IMinimalForwarderForwardRequest) FromSubmitTransactionRequest(
	req *common.SubmitTransactionRequest,
) {
	r.From = req.From
	r.To = req.To
	r.Value = req.Value
	r.Gas = req.Gas
	r.Nonce = req.Nonce
	r.Data = req.Data
}

// Pack packs the IForwarderForwardRequest data into an ABI-encoded format.
func (r *IMinimalForwarderForwardRequest) Pack(_ []byte) ([]byte, error) {
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
	}
	packed, err := args.Pack(
		forwardRequestTypehash,
		r.From,
		r.To,
		r.Value,
		r.Gas,
		r.Nonce,
		hashedData,
	)
	if err != nil {
		return nil, err
	}

	return packed, nil
}
