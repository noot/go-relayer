package gsnforwarder

import (
	"errors"
	"math/big"

	"github.com/athanorlabs/go-relayer/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ common.Forwarder = &IForwarderWrapped{}

var errInvalidForwardRequestType = errors.New("invalid ForwardRequest type")

// IForwarderWrapped represents a wrapped GSN IForwarder.
type IForwarderWrapped struct {
	f *IForwarder
}

// NewIForwarderWrapped returns a new IForwarderWrapped.
func NewIForwarderWrapped(f *IForwarder) *IForwarderWrapped {
	return &IForwarderWrapped{
		f: f,
	}
}

// GetNonce returns an account's nonce from the forwarder contract.
func (f *IForwarderWrapped) GetNonce(opts *bind.CallOpts, from ethcommon.Address) (*big.Int, error) {
	return f.f.GetNonce(opts, from)
}

// Verify verifies a forward request using the forwarder contract.
func (f *IForwarderWrapped) Verify(
	opts *bind.CallOpts,
	req common.ForwardRequest,
	domainSeparator,
	requestTypeHash [32]byte,
	suffixData,
	signature []byte,
) (bool, error) {
	r, ok := req.(*IForwarderForwardRequest)
	if !ok {
		return false, errInvalidForwardRequestType
	}

	err := f.f.Verify(opts, *r, domainSeparator, requestTypeHash, suffixData, signature)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Execute executes a forward request using the forwarder contract.
func (f *IForwarderWrapped) Execute(
	opts *bind.TransactOpts,
	req common.ForwardRequest,
	domainSeparator,
	requestTypeHash [32]byte,
	suffixData,
	signature []byte,
) (*types.Transaction, error) {
	r, ok := req.(*IForwarderForwardRequest)
	if !ok {
		return nil, errInvalidForwardRequestType
	}

	return f.f.Execute(opts, *r, domainSeparator, requestTypeHash, suffixData, signature)
}

// NewEmptyForwardRequest returns a new empty *IForwarderForwardRequest.
func (f *IForwarderWrapped) NewEmptyForwardRequest() common.ForwardRequest {
	return &IForwarderForwardRequest{}
}
