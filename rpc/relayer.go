package rpc

import (
	"net/http"
)

type RelayerService struct{}

func NewRelayerService() *RelayerService {
	return &RelayerService{}
}

type SubmitTransactionRequest struct{}

type SubmitTransactionResponse struct{}

func (s *RelayerService) SubmitTransaction(_ *http.Request, req *SubmitTransactionRequest, resp *SubmitTransactionResponse) error {
	return nil
}
