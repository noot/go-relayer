package mforwarder

import (
	"errors"
	"math/big"

	"github.com/athanorlabs/go-relayer/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ common.Forwarder = &IMinimalForwarderWrapped{}

var errInvalidForwardRequestType = errors.New("invalid ForwardRequest type")

// IMinimalForwarderWrapped represents a wrapped MinimalForwarder.
type IMinimalForwarderWrapped struct {
	f *IMinimalForwarder
}

// NewIMinimalForwarderWrapped returns a new IMinimalForwarderWrapped.
func NewIMinimalForwarderWrapped(f *IMinimalForwarder) *IMinimalForwarderWrapped {
	return &IMinimalForwarderWrapped{
		f: f,
	}
}

// GetNonce returns an account's nonce from the forwarder contract.
func (f *IMinimalForwarderWrapped) GetNonce(opts *bind.CallOpts, from ethcommon.Address) (*big.Int, error) {
	return f.f.GetNonce(opts, from)
}

// Verify verifies a forward request using the forwarder contract.
func (f *IMinimalForwarderWrapped) Verify(
	opts *bind.CallOpts,
	req common.ForwardRequest,
	_,
	_ [32]byte,
	_,
	signature []byte,
) (bool, error) {
	r, ok := req.(*IMinimalForwarderForwardRequest)
	if !ok {
		return false, errInvalidForwardRequestType
	}

	return f.f.Verify(opts, *r, signature)
}

// Execute executes a forward request using the forwarder contract.
func (f *IMinimalForwarderWrapped) Execute(
	opts *bind.TransactOpts,
	req common.ForwardRequest,
	_,
	_ [32]byte,
	_,
	signature []byte,
) (*types.Transaction, error) {
	r, ok := req.(*IMinimalForwarderForwardRequest)
	if !ok {
		return nil, errInvalidForwardRequestType
	}

	return f.f.Execute(opts, *r, signature)
}

// NewEmptyForwardRequest returns a new empty *IMinimalForwarderForwardRequest.
func (f *IMinimalForwarderWrapped) NewEmptyForwardRequest() common.ForwardRequest {
	return &IMinimalForwarderForwardRequest{}
}
