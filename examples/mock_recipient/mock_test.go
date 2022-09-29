package mock

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/AthanorLabs/go-relayer/common"
	mforwarder "github.com/AthanorLabs/go-relayer/examples/minimal_forwarder"
	"github.com/athanorlabs/atomic-swap/ethereum/block"
)

// NODE_OPTIONS="--max_old_space_size=8192" ganache --deterministic --accounts=50
var ganachePrivateKeys = []string{
	"4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
	"6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1",
}

func setupAuth(t *testing.T) (*bind.TransactOpts, *ethclient.Client, *ecdsa.PrivateKey) {
	ec, err := ethclient.Dial(common.DefaultEthEndpoint)
	require.NoError(t, err)
	t.Cleanup(func() {
		ec.Close()
	})
	chainID, err := ec.ChainID(context.Background())
	require.NoError(t, err)

	pk, err := ethcrypto.HexToECDSA(ganachePrivateKeys[0])
	require.NoError(t, err)

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	require.NoError(t, err)
	return auth, ec, pk
}

func TestMock_Execute(t *testing.T) {
	auth, conn, pk := setupAuth(t)
	chainID, err := conn.ChainID(context.Background())
	require.NoError(t, err)

	address, tx, contract, err := mforwarder.DeployMinimalForwarder(auth, conn)
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
		big.NewInt(531099458),
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

	abi, err := abi.JSON(strings.NewReader(MockABI))
	require.NoError(t, err)
	calldata, err := abi.Pack("withdraw", value, fee)
	require.NoError(t, err)

	req := &mforwarder.IMinimalForwarderForwardRequest{
		From:  key.Address(),
		To:    mockAddress,
		Value: big.NewInt(0),
		Gas:   big.NewInt(679639582),
		Nonce: big.NewInt(0),
		Data:  calldata,
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

	// in real life, the signer of the digest is the end-user, not the relayer
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
