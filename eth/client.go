package eth

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	ec     *ethclient.Client
	txOpts *bind.TransactOpts
}

type Config struct {
	EthClient  *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
	ChainID    *big.Int
}

func NewClient(cfg *Config) (*Client, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(cfg.PrivateKey, cfg.ChainID)
	if err != nil {
		return nil, err
	}

	return &Client{
		ec:     cfg.EthClient,
		txOpts: txOpts,
	}, nil
}

func (c *Client) SubmitTransaction() {}
