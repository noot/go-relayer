package mock

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/go-relayer/common"
	"github.com/athanorlabs/go-relayer/impls/mforwarder"
	"github.com/athanorlabs/go-relayer/tests"
)

// NODE_OPTIONS="--max_old_space_size=8192" ganache --deterministic --accounts=50
var ganachePrivateKeys = []string{
	"6370fd033278c143179d81c5526140625662b8daa446c22ee2d73db3707e620c",
	"646f1ce2fdad0e6deeeb5c7e8e5543bdde65e86029e2fd9fc169899c440a7913",
}

func TestMock_Execute(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	pk, err := ethcrypto.HexToECDSA(ganachePrivateKeys[0])
	require.NoError(t, err)

	ec, chainID := tests.NewEthClient(t)

	txOpts := tests.NewTXOpts(t, ec, pk)
	address, tx, contract, err := mforwarder.DeployMinimalForwarder(txOpts, ec)
	require.NoError(t, err)
	receipt := tests.MineTransaction(t, ec, tx)
	t.Logf("gas cost to deploy MinimalForwarder.sol: %d", receipt.GasUsed)

	txOpts = tests.NewTXOpts(t, ec, pk)
	mockAddress, mockTx, _, err := DeployMock(txOpts, ec, address)
	require.NoError(t, err)
	receipt = tests.MineTransaction(t, ec, mockTx)
	t.Logf("gas cost to deploy Mock.sol: %d", receipt.GasUsed)

	// transfer to Mock.sol
	value := big.NewInt(1000000)
	fee := big.NewInt(10000)

	gasPrice, err := ec.SuggestGasPrice(ctx)
	require.NoError(t, err)
	t.Logf("suggested gas price: %d", gasPrice)

	fromAddress := ethcrypto.PubkeyToAddress(*pk.Public().(*ecdsa.PublicKey))
	nonce, err := ec.PendingNonceAt(ctx, fromAddress)
	require.NoError(t, err)

	transferTx := ethtypes.NewTx(&ethtypes.LegacyTx{
		Nonce:    nonce,
		To:       &mockAddress,
		Value:    value,
		Gas:      100000,
		GasPrice: gasPrice,
	})
	transferTx, err = ethtypes.SignTx(transferTx, ethtypes.LatestSignerForChainID(chainID), pk)
	require.NoError(t, err)

	err = ec.SendTransaction(ctx, transferTx)
	require.NoError(t, err)

	receipt, err = bind.WaitMined(ctx, ec, transferTx)
	require.NoError(t, err)

	require.Equal(t, ethtypes.ReceiptStatusSuccessful, receipt.Status)
	t.Logf("transfer sent: %s", transferTx.Hash())

	// generate ForwardRequest and sign it
	key := common.NewKeyFromPrivateKey(pk)

	abi, err := MockMetaData.GetAbi()
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
	t.Logf("verified forward request")

	// execute withdraw() via forwarder
	txOpts = tests.NewTXOpts(t, ec, pk)
	tx, err = contract.Execute(txOpts, *req, sig)
	require.NoError(t, err)

	receipt = tests.MineTransaction(t, ec, tx)
	t.Logf("gas cost to call Mock.withdraw() via MinimalForwarder.execute(): %d", receipt.GasUsed)
	require.Equal(t, uint64(1), receipt.Status)
	require.Equal(t, 1, len(receipt.Logs))

	// check that transfer worked
	mockBalance, err := ec.BalanceAt(ctx, mockAddress, nil)
	require.NoError(t, err)
	require.Equal(t, int64(0), mockBalance.Int64())
}
