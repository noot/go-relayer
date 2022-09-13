package relayer

import (
	"context"
	"errors"
	"math/big"

	"github.com/AthanorLabs/go-relayer/common"
	//"github.com/AthanorLabs/go-relayer/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	errFailedToVerify = errors.New("failed to verify forward request signature")
)

type NewForwarderRequestFunc[T any] func(ethcommon.Address, ethcommon.Address, *big.Int, *big.Int, *big.Int, []byte, *big.Int) T

// Relayer represents a transaction relayer.
// It contains an Ethereum client and a private key, allowing it to forward transactions.
type Relayer[T any] struct {
	ctx                     context.Context
	ec                      *ethclient.Client
	forwarder               Forwarder[T]
	callOpts                *bind.CallOpts
	txOpts                  *bind.TransactOpts
	newForwarderRequestFunc NewForwarderRequestFunc[T]
}

// Config ...
type Config[T any] struct {
	Ctx                     context.Context
	EthClient               *ethclient.Client
	Forwarder               Forwarder[T]
	Key                     *common.Key
	ChainID                 *big.Int
	NewForwarderRequestFunc NewForwarderRequestFunc[T]
}

func NewRelayer[T any](cfg *Config[T]) (*Relayer[T], error) {
	callOpts := &bind.CallOpts{
		From:    cfg.Key.Address(),
		Context: cfg.Ctx,
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(cfg.Key.PrivateKey(), cfg.ChainID)
	if err != nil {
		return nil, err
	}

	return &Relayer[T]{
		ctx:                     cfg.Ctx,
		ec:                      cfg.EthClient,
		forwarder:               cfg.Forwarder,
		callOpts:                callOpts,
		txOpts:                  txOpts,
		newForwarderRequestFunc: cfg.NewForwarderRequestFunc,
	}, nil
}

// TODO: get rid of this, make SubmitTransaction take T
// actually, we might need this for RPC requests (json tags)
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

func (s *Relayer[T]) SubmitTransaction(req *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	fwdReq := s.newForwarderRequestFunc(req.From, req.To, req.Value, req.Gas, req.Nonce, req.Data, nil)

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

	return &SubmitTransactionResponse{
		TxHash: tx.Hash(),
	}, nil
}
