package rpc

import (
	"net/http"

	"github.com/AthanorLabs/go-relayer/common"
	"github.com/AthanorLabs/go-relayer/relayer"
)

type RelayerService struct {
	r *relayer.Relayer
}

func NewRelayerService(r *relayer.Relayer) (*RelayerService, error) {
	return &RelayerService{
		r: r,
	}, nil
}

func (s *RelayerService) SubmitTransaction(
	_ *http.Request,
	req *common.SubmitTransactionRequest,
	resp *common.SubmitTransactionResponse,
) error {
	rresp, err := s.r.SubmitTransaction(req)
	if err != nil {
		return err
	}

	resp.TxHash = rresp.TxHash
	return nil
}
