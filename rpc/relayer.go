package rpc

import (
	"context"
	"errors"
	"math/big"
	"net/http"

	"github.com/AthanorLabs/go-relayer/common"
	"github.com/AthanorLabs/go-relayer/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	errFailedToVerify = errors.New("failed to verify forward request signature")
)

type RelayerService struct {
	ec        *ethclient.Client
	forwarder *contracts.MinimalForwarder // TODO: make this an interface
	callOpts  *bind.CallOpts
	txOpts    *bind.TransactOpts
}

func NewRelayerService(
	ctx context.Context,
	ec *ethclient.Client,
	forwarder *contracts.MinimalForwarder,
	key *common.Key,
	chainID *big.Int,
) (*RelayerService, error) {
	callOpts := &bind.CallOpts{
		From:    key.Address(),
		Context: ctx,
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey(), chainID)
	if err != nil {
		return nil, err
	}

	return &RelayerService{
		ec:        ec,
		forwarder: forwarder,
		callOpts:  callOpts,
		txOpts:    txOpts,
	}, nil
}

type SubmitTransactionRequest struct {
	From      ethcommon.Address
	To        ethcommon.Address
	Value     *big.Int
	Gas       *big.Int
	Nonce     *big.Int
	Data      []byte
	Signature []byte
}

type SubmitTransactionResponse struct {
	TxHash ethcommon.Hash `json:"transactionHash"`
}

func (s *RelayerService) SubmitTransaction(_ *http.Request, req *SubmitTransactionRequest, resp *SubmitTransactionResponse) error {
	fwdReq := &contracts.MinimalForwarderForwardRequest{
		From:  req.From,
		To:    req.To,
		Value: req.Value,
		Gas:   req.Gas,
		Nonce: req.Nonce,
		Data:  req.Data,
	}

	// verify sig beforehand
	ok, err := s.forwarder.Verify(s.callOpts, *fwdReq, req.Signature)
	if !ok {
		return errFailedToVerify
	}

	if err != nil {
		return err
	}

	tx, err := s.forwarder.Execute(s.txOpts, *fwdReq, req.Signature)
	if err != nil {
		return err
	}

	resp.TxHash = tx.Hash()
	return nil
}
