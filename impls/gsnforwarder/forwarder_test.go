package gsnforwarder

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

func TestForwarder_Verify(t *testing.T) {
	ec, chainID := tests.NewEthClient(t)
	pk := tests.GetTestKeyByIndex(t, 0)
	auth := tests.NewTXOpts(t, ec, pk)

	address, tx, contract, err := DeployForwarder(auth, ec)
	require.NoError(t, err)
	require.NotEqual(t, ethcommon.Address{}, address)
	require.NotNil(t, tx)
	require.NotNil(t, contract)
	receipt := tests.MineTransaction(t, ec, tx)
	t.Logf("gas cost to deploy MinimalForwarder.sol: %d", receipt.GasUsed)

	key := common.NewKeyFromPrivateKey(pk)

	auth = tests.NewTXOpts(t, ec, pk)
	tx, err = contract.RegisterDomainSeparator(auth, DefaultName, DefaultVersion)
	require.NoError(t, err)
	receipt = tests.MineTransaction(t, ec, tx)
	t.Logf("gas cost to call RegisterDomainSeparator: %d", receipt.GasUsed)

	req := &IForwarderForwardRequest{
		From:           key.Address(),
		To:             ethcommon.Address{2}, // arbitrary
		Value:          big.NewInt(0),
		Gas:            big.NewInt(7000000),
		Nonce:          big.NewInt(0),
		Data:           []byte{},
		ValidUntilTime: big.NewInt(0),
	}

	// TODO: fix suffixData encoding - geth's abi.Pack isn't working :(
	suffixData := []byte("suffixData")

	domainSeparator, err := common.GetEIP712DomainSeparator(DefaultName, DefaultVersion, chainID, address)
	require.NoError(t, err)

	digest, err := common.GetForwardRequestDigestToSign(req, domainSeparator, suffixData)
	require.NoError(t, err)

	sig, err := key.Sign(digest)
	require.NoError(t, err)

	callOpts := &bind.CallOpts{
		From:    key.Address(),
		Context: context.Background(),
	}

	err = contract.Verify(
		callOpts,
		*req,
		domainSeparator,
		ForwardRequestTypehash,
		nil, //suffixData,
		sig,
	)
	require.NoError(t, err)
}

func TestForwarder_IsIForwarder(t *testing.T) {
	ec, _ := tests.NewEthClient(t)
	pk := tests.GetTestKeyByIndex(t, 0)
	auth := tests.NewTXOpts(t, ec, pk)

	address, tx, _, err := DeployForwarder(auth, ec)
	require.NoError(t, err)
	tests.MineTransaction(t, ec, tx)

	_, err = NewIForwarder(address, ec)
	require.NoError(t, err)
}
