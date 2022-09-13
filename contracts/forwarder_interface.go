package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ ForwarderInterface[IMinimalForwarderForwardRequest] = &IMinimalForwarder{}

type ForwarderInterface[T any] interface {
	GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error)
	Verify(opts *bind.CallOpts, req T, signature []byte) (bool, error)
	Execute(opts *bind.TransactOpts, req T, signature []byte) (*types.Transaction, error)
}
