package contracts

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/AthanorLabs/go-relayer/common"
	"github.com/athanorlabs/atomic-swap/ethereum/block"
)

func TestMock_Execute(t *testing.T) {
	auth, conn, pk := setupAuth(t)
	chainID, err := conn.ChainID(context.Background())
	require.NoError(t, err)

	address, tx, contract, err := DeployMinimalForwarder(auth, conn)
	require.NoError(t, err)
	receipt, err := block.WaitForReceipt(context.Background(), conn, tx.Hash())
	require.NoError(t, err)
	t.Logf("gas cost to deploy MinimalForwarder.sol: %d", receipt.GasUsed)

	mockAddress, mockTx, _, err := DeployMock(auth, conn, address)
	require.NoError(t, err)
	receipt, err = block.WaitForReceipt(context.Background(), conn, mockTx.Hash())
	require.NoError(t, err)
	t.Logf("gas cost to deploy Mock.sol: %d", receipt.GasUsed)

	// transfer to Mock.sol
	value := big.NewInt(1000000)
	fee := big.NewInt(10000)

	transferTx := ethtypes.NewTransaction(
		0,
		mockAddress,
		value,
		100000,
		big.NewInt(526456961),
		nil,
	)

	transferTx, err = ethtypes.SignTx(transferTx, ethtypes.LatestSignerForChainID(chainID), pk)
	require.NoError(t, err)
	err = conn.SendTransaction(context.Background(), transferTx)
	require.NoError(t, err)
	_, err = block.WaitForReceipt(context.Background(), conn, transferTx.Hash())
	require.NoError(t, err)

	// generate ForwardRequest and sign it
	key := common.NewKeyFromPrivateKey(pk)

	functionSig := GetFunctionSignature("withdraw(uint256,uint256)")
	require.Equal(t, "441a3e70", hex.EncodeToString(functionSig))

	uint256Ty, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)

	args := &abi.Arguments{
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
	}
	params, err := args.Pack(value, fee)
	require.NoError(t, err)

	req := &MinimalForwarderForwardRequest{
		From:  key.Address(),
		To:    mockAddress,
		Value: big.NewInt(0),
		Gas:   big.NewInt(7000000),
		Nonce: big.NewInt(0),
		Data:  append(functionSig, params...),
	}

	digest, err := GetForwardRequestDigestToSign(req, chainID, address)
	require.NoError(t, err)

	sig, err := key.Sign(digest)
	require.NoError(t, err)

	callOpts := &bind.CallOpts{
		From:    key.Address(),
		Context: context.Background(),
	}

	// verify sig beforehand
	ok, err := contract.Verify(callOpts, *req, sig)
	require.NoError(t, err)
	require.True(t, ok)

	// execute withdraw() via forwarder
	tx, err = contract.Execute(auth, *req, sig)
	require.NoError(t, err)
	receipt, err = block.WaitForReceipt(context.Background(), conn, tx.Hash())
	require.NoError(t, err)
	t.Logf("gas cost to call Mock.withdraw() via MinimalForwarder.execute(): %d", receipt.GasUsed)
	require.Equal(t, uint64(1), receipt.Status)
	require.Equal(t, 1, len(receipt.Logs))

	// check that transfer worked
	mockBalance, err := conn.BalanceAt(context.Background(), mockAddress, nil)
	require.NoError(t, err)
	require.Equal(t, int64(0), mockBalance.Int64())
}
