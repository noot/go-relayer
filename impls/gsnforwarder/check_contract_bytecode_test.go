package gsnforwarder

import (
	"context"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/go-relayer/tests"
)

func TestCheckForwarderContractCode_mainnet(t *testing.T) {
	ec := getMainnetClient(t)
	err := CheckForwarderContractCode(context.Background(), ec, ethcommon.HexToAddress(MainnetForwarderAddrHex))
	require.NoError(t, err)
}

func TestCheckForwarderContractCode_goerli(t *testing.T) {
	ec := getGoerliClient(t)
	err := CheckForwarderContractCode(context.Background(), ec, ethcommon.HexToAddress(GoerliForwarderAddrHex))
	require.NoError(t, err)
}

func TestCheckForwarderContractCode_sepolia(t *testing.T) {
	ec := getSepoliaClient(t)
	err := CheckForwarderContractCode(context.Background(), ec, ethcommon.HexToAddress(SepoliaForwarderAddrHex))
	require.NoError(t, err)
}

func TestCheckForwarderContractCode_fail(t *testing.T) {
	ec, _ := tests.NewEthClient(t)
	code, addr := getOurForwarderByteCode(t)

	// Fail on the CodeAt error check by using a cancelled context (bad addresses
	// don't seem to fail it).
	canceledContext, cancel := context.WithCancel(context.Background())
	cancel()
	err := CheckForwarderContractCode(canceledContext, ec, addr)
	require.ErrorIs(t, err, context.Canceled)

	// too short
	err = checkContractBytes(make([]byte, 0))
	require.ErrorIs(t, err, errInvalidForwarderContract)

	// bytes don't match
	code[5]++
	err = checkContractBytes(code)

	// too long
	err = checkContractBytes(make([]byte, 2048))
	require.ErrorIs(t, err, errInvalidForwarderContract)
}
