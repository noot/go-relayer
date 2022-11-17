package mforwarder

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/go-relayer/common"
	"github.com/athanorlabs/go-relayer/tests"
)

func TestMinimalForwarder_Verify(t *testing.T) {
	ec, chainID := tests.NewEthClient(t)
	pk := tests.GetTestKeyByIndex(t, 0)
	auth := tests.NewTXOpts(t, ec, pk)

	address, tx, contract, err := DeployMinimalForwarder(auth, ec)
	require.NoError(t, err)
	require.NotEqual(t, ethcommon.Address{}, address)
	require.NotNil(t, tx)
	require.NotNil(t, contract)
	receipt, err := bind.WaitMined(context.Background(), ec, tx)
	require.NoError(t, err)
	t.Logf("gas cost to deploy MinimalForwarder.sol: %d", receipt.GasUsed)

	key := common.NewKeyFromPrivateKey(pk)

	req := &IMinimalForwarderForwardRequest{
		From:  key.Address(),
		To:    ethcommon.Address{2}, // arbitrary
		Value: big.NewInt(0),
		Gas:   big.NewInt(7000000),
		Nonce: big.NewInt(0),
		Data:  []byte{},
	}

	name := "MinimalForwarder"
	version := "0.0.1"

	domainSeparator, err := common.GetEIP712DomainSeparator(name, version, chainID, address)
	require.NoError(t, err)

	digest, err := common.GetForwardRequestDigestToSign(
		req,
		domainSeparator,
		nil,
	)
	require.NoError(t, err)

	sig, err := key.Sign(digest)
	require.NoError(t, err)

	callOpts := &bind.CallOpts{
		From:    key.Address(),
		Context: context.Background(),
	}

	ok, err := contract.Verify(callOpts, *req, sig)
	require.NoError(t, err)
	require.True(t, ok)
}

func TestMinimalForwarder_IsIMinimalForwarder(t *testing.T) {
	ec, _ := tests.NewEthClient(t)
	pk := tests.GetTestKeyByIndex(t, 0)
	auth := tests.NewTXOpts(t, ec, pk)

	address, tx, _, err := DeployMinimalForwarder(auth, ec)
	require.NoError(t, err)
	_, err = bind.WaitMined(context.Background(), ec, tx)
	require.NoError(t, err)

	_, err = NewIMinimalForwarder(address, ec)
	require.NoError(t, err)
}
