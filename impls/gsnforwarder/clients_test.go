package gsnforwarder

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// These file is for client test helper functions

const (
	fallbackMainnetEndpoint = "https://rpc.flashbots.net"
	fallbackGoerliEndpoint  = "https://rpc-goerli.flashbots.net"
	fallbackSepoliaEndpoint = "https://rpc.sepolia.org/"
)

func getMainnetClient(t *testing.T) *ethclient.Client {
	endpoint := os.Getenv("ETH_MAINNET_ENDPOINT")
	if endpoint == "" {
		endpoint = fallbackMainnetEndpoint
	}

	ec, err := ethclient.Dial(endpoint)
	require.NoError(t, err)
	t.Cleanup(func() {
		ec.Close()
	})

	return ec
}

func getGoerliClient(t *testing.T) *ethclient.Client {
	endpoint := os.Getenv("ETH_GOERLI_ENDPOINT")
	if endpoint == "" {
		endpoint = fallbackGoerliEndpoint
	}

	ec, err := ethclient.Dial(endpoint)
	require.NoError(t, err)
	t.Cleanup(func() {
		ec.Close()
	})

	return ec
}

func getSepoliaClient(t *testing.T) *ethclient.Client {
	endpoint := os.Getenv("ETH_SEPOLIA_ENDPOINT")
	if endpoint == "" {
		endpoint = fallbackSepoliaEndpoint
	}

	ec, err := ethclient.Dial(endpoint)
	require.NoError(t, err)
	t.Cleanup(func() {
		ec.Close()
	})

	return ec
}
