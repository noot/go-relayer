package contracts

import (
	"errors"
	"math/big"

	"github.com/AthanorLabs/go-relayer/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ common.Forwarder = &IMinimalForwarderWrapped{}

var errInvalidForwardRequestType = errors.New("invalid ForwardRequest type")

type IMinimalForwarderWrapped struct {
	f *IMinimalForwarder
}

func NewIMinimalForwarderWrapped(f *IMinimalForwarder) *IMinimalForwarderWrapped {
	return &IMinimalForwarderWrapped{
		f: f,
	}
}

func (f *IMinimalForwarderWrapped) GetNonce(opts *bind.CallOpts, from ethcommon.Address) (*big.Int, error) {
	return f.f.GetNonce(opts, from)
}

func (f *IMinimalForwarderWrapped) Verify(opts *bind.CallOpts, req common.ForwardRequest, signature []byte) (bool, error) {
	r, ok := req.(*IMinimalForwarderForwardRequest)
	if !ok {
		return false, errInvalidForwardRequestType
	}

	return f.f.Verify(opts, *r, signature)
}

func (f *IMinimalForwarderWrapped) Execute(opts *bind.TransactOpts, req common.ForwardRequest, signature []byte) (*types.Transaction, error) {
	r, ok := req.(*IMinimalForwarderForwardRequest)
	if !ok {
		return nil, errInvalidForwardRequestType
	}

	return f.f.Execute(opts, *r, signature)
}
