package relayer

import (
	"context"
	"errors"
	"math/big"

	"github.com/AthanorLabs/go-relayer/common"
	"github.com/AthanorLabs/go-relayer/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	errFailedToVerify = errors.New("failed to verify forward request signature")
)

// Relayer represents a transaction relayer.
// It contains an Ethereum client and a private key, allowing it to forward transactions.
type Relayer struct {
	ctx       context.Context
	ec        *ethclient.Client
	forwarder *contracts.IMinimalForwarder
	callOpts  *bind.CallOpts
	txOpts    *bind.TransactOpts
}

// Config ...
type Config struct {
	Ctx       context.Context
	EthClient *ethclient.Client
	Forwarder *contracts.IMinimalForwarder
	Key       *common.Key
	ChainID   *big.Int
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
		ctx:       cfg.Ctx,
		ec:        cfg.EthClient,
		forwarder: cfg.Forwarder,
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

func (s *Relayer) SubmitTransaction(req *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	fwdReq := &contracts.IMinimalForwarderForwardRequest{
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
		return nil, errFailedToVerify
	}

	if err != nil {
		return nil, err
	}

	tx, err := s.forwarder.Execute(s.txOpts, *fwdReq, req.Signature)
	if err != nil {
		return nil, err
	}

	return &SubmitTransactionResponse{
		TxHash: tx.Hash(),
	}, nil
}
