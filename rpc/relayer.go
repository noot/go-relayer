package rpc

import (
	"net/http"

	"github.com/athanorlabs/go-relayer/common"
	"github.com/athanorlabs/go-relayer/relayer"
)

// RelayerService represents a relayer JSON-RPC service.
type RelayerService struct {
	r *relayer.Relayer
}

// NewRelayerService returns a new RPC RelayerService.
func NewRelayerService(r *relayer.Relayer) (*RelayerService, error) {
	return &RelayerService{
		r: r,
	}, nil
}

// SubmitTransaction submits a transaction to the relayer.
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
