package contracts

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/atomic-swap/common"
	"github.com/athanorlabs/atomic-swap/ethereum/block"
)

// NODE_OPTIONS="--max_old_space_size=8192" ganache --deterministic --accounts=50
var ganachePrivateKeys = []string{
	"4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
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

func TestForwarder_Verify(t *testing.T) {
	auth, conn, _ := setupAuth(t)
	address, tx, contract, err := DeployMinimalForwarder(auth, conn)
	require.NoError(t, err)
	require.NotEqual(t, ethcommon.Address{}, address)
	require.NotNil(t, tx)
	require.NotNil(t, contract)
	receipt, err := block.WaitForReceipt(context.Background(), conn, tx.Hash())
	require.NoError(t, err)
	t.Logf("gas cost to deploy MinimalForwarder.sol: %d", receipt.GasUsed)
}
