package rpc

import (
	"net/http"

	"github.com/AthanorLabs/go-relayer/relayer"
)

type RelayerService[T any] struct {
	r *relayer.Relayer[T]
}

func NewRelayerService[T any](r *relayer.Relayer[T]) (*RelayerService[T], error) {
	return &RelayerService[T]{
		r: r,
	}, nil
}

func (s *RelayerService[T]) SubmitTransaction(
	_ *http.Request,
	req *relayer.SubmitTransactionRequest,
	resp *relayer.SubmitTransactionResponse,
) error {
	rresp, err := s.r.SubmitTransaction(req)
	if err != nil {
		return err
	}

	resp.TxHash = rresp.TxHash
	return nil
}
