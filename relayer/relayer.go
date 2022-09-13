package relayer

import (
	"context"
	"errors"
	"math/big"

	"github.com/AthanorLabs/go-relayer/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	errFailedToVerify = errors.New("failed to verify forward request signature")
)

type NewForwardRequestFunc func() common.ForwardRequest

// Relayer represents a transaction relayer.
// It contains an Ethereum client and a private key, allowing it to forward transactions.
type Relayer struct {
	ctx                   context.Context
	ec                    *ethclient.Client
	forwarder             common.Forwarder
	callOpts              *bind.CallOpts
	txOpts                *bind.TransactOpts
	newForwardRequestFunc NewForwardRequestFunc
}

// Config ...
type Config struct {
	Ctx                   context.Context
	EthClient             *ethclient.Client
	Forwarder             common.Forwarder
	Key                   *common.Key
	ChainID               *big.Int
	NewForwardRequestFunc NewForwardRequestFunc
}

func NewRelayer(cfg *Config) (*Relayer, error) {
	callOpts := &bind.CallOpts{
		From:    cfg.Key.Address(),
		Context: cfg.Ctx,
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(cfg.Key.PrivateKey(), cfg.ChainID)
	if err != nil {
		return nil, err
	}

	return &Relayer{
		ctx:                   cfg.Ctx,
		ec:                    cfg.EthClient,
		forwarder:             cfg.Forwarder,
		callOpts:              callOpts,
		txOpts:                txOpts,
		newForwardRequestFunc: cfg.NewForwardRequestFunc,
	}, nil
}

func (s *Relayer) SubmitTransaction(req *common.SubmitTransactionRequest) (*common.SubmitTransactionResponse, error) {
	fwdReq := s.newForwardRequestFunc()
	fwdReq.FromSubmitTransactionRequest(req)

	// verify sig beforehand
	ok, err := s.forwarder.Verify(s.callOpts, fwdReq, req.Signature)
	if !ok {
		return nil, errFailedToVerify
	}

	if err != nil {
		return nil, err
	}

	tx, err := s.forwarder.Execute(s.txOpts, fwdReq, req.Signature)
	if err != nil {
		return nil, err
	}

	return &common.SubmitTransactionResponse{
		TxHash: tx.Hash(),
	}, nil
}
