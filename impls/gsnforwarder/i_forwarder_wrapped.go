package gsnforwarder

import (
	"errors"
	"math/big"

	"github.com/AthanorLabs/go-relayer/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ common.Forwarder = &IForwarderWrapped{}

var errInvalidForwardRequestType = errors.New("invalid ForwardRequest type")

type IForwarderWrapped struct {
	f *IForwarder
}

func NewIForwarderWrapped(f *IForwarder) *IForwarderWrapped {
	return &IForwarderWrapped{
		f: f,
	}
}

func (f *IForwarderWrapped) GetNonce(opts *bind.CallOpts, from ethcommon.Address) (*big.Int, error) {
	return f.f.GetNonce(opts, from)
}

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
