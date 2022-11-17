package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"runtime"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/go-relayer/common"
)

/*
 * Golang packages can be tested in parallel, but in our case, they are sharing a
 * common ganache Ethereum simulator. To avoid transaction conflicts, we allocate
 * each package listed in `testPackageNames` one or more keys not shared with the
 * tests of any other package.
 */

// testPackageNames is the list of packages with _test.go code that requires access to
// one or more prefunded ganache wallets.
var testPackages = []struct {
	name    string
	numKeys int
}{
	{"impls/gsnforwarder", 1},
	{"impls/mforwarder", 1},
}

const (
	repoName = "github.com/athanorlabs/go-relayer/"
)

// `ganache --deterministic` provides the following keys with 1000 ETH on startup.
var ganacheTestKeys = []string{
	"4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d", // ganache key #0
	"6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1", // ganache key #1
	"6370fd033278c143179d81c5526140625662b8daa446c22ee2d73db3707e620c", // ganache key #2
	"646f1ce2fdad0e6deeeb5c7e8e5543bdde65e86029e2fd9fc169899c440a7913", // ganache key #3
	"add53f9a7e588d003326d1cbf9e4a43c061aadd9bc938c843a79e7b4fd2ad743", // ganache key #4
	"395df67f0c2d2d9fe1ad08d1bc8b6627011959b79c53d7dd6a3536a33ab8a4fd", // ganache key #5
	"e485d098507f54e7733a205420dfddbe58db035fa577fc294ebd14db90767a52", // ganache key #6
	"a453611d9419d0e56f499079478fd72c37b251a94bfde4d19872c44cf65386e3", // ganache key #7
	"829e924fdf021ba3dbbc4225edfece9aca04b929d6e75613329ca6f1d31c0bb4", // ganache key #8
	"b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773", // ganache key #9
}

func init() {
	totalKeys := 0
	for _, pkg := range testPackages {
		totalKeys += pkg.numKeys
	}
	if totalKeys > len(ganacheTestKeys) {
		panic("Insufficient ganache test keys")
	}
}

// minPackageName takes a long-form package+function name (example:
// "github.com/athanorlabs/atomic-swap/protocol/xmrtaker.newBackend") and returns
// just the package name without the repository prefix ("protocol/xmrtaker").
func minPackageName(t *testing.T, pkgAndFunc string) string {
	minPkgAndFunc := strings.TrimPrefix(pkgAndFunc, repoName)
	if minPkgAndFunc == pkgAndFunc {
		t.Fatalf("%q does not have the repo prefix %q", pkgAndFunc, repoName)
	}
	// with the domain name gone, the minimal package is everything before the first period.
	return strings.Split(minPkgAndFunc, ".")[0]
}

func getCallingPackageName(t *testing.T) string {
	// Determine the test package that requested the key from the call stack. We skip 2 callers
	// (1) this function and (2) the public function from this package that invoked it.
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		t.Fatalf("Failed to get caller info")
	}
	fullPackageName := runtime.FuncForPC(pc).Name()
	return minPackageName(t, fullPackageName)
}

func getPackageKeys(t *testing.T, packageName string) []string {
	startIndex := 0
	for _, pkg := range testPackages {
		if pkg.name == packageName {
			return ganacheTestKeys[startIndex : startIndex+pkg.numKeys]
		}
		startIndex += pkg.numKeys
	}
	t.Fatalf("Package %q does not have reserved test keys", packageName)
	panic("unreachable code")
}

func getPackageTestKey(t *testing.T, pkgName string, index int) *ecdsa.PrivateKey {
	keys := getPackageKeys(t, pkgName)
	require.Lessf(t, index, len(keys), "insufficient keys allocated to package %q", pkgName)
	pk, err := ethcrypto.HexToECDSA(keys[index])
	require.NoError(t, err)
	return pk
}

// GetTestKeyByIndex returns the ganache test key allocated to a package by index
func GetTestKeyByIndex(t *testing.T, index int) *ecdsa.PrivateKey {
	pkgName := getCallingPackageName(t)
	return getPackageTestKey(t, pkgName, index)
}

// NewEthClient returns a connection to the local ganache instance. The connection is
// automatically closed when the test completes.
func NewEthClient(t *testing.T) (*ethclient.Client, *big.Int) {
	ec, err := ethclient.Dial(common.DefaultEthEndpoint)
	require.NoError(t, err)
	t.Cleanup(func() {
		ec.Close()
	})
	chainID, err := ec.ChainID(context.Background())
	require.NoError(t, err)
	return ec, chainID
}

// NewTXOpts creates a new *bind.TransactOpts for tests
func NewTXOpts(t *testing.T, ec *ethclient.Client, privkey *ecdsa.PrivateKey) *bind.TransactOpts {
	chainID, err := ec.ChainID(context.Background())
	require.NoError(t, err)
	txOpts, err := bind.NewKeyedTransactorWithChainID(privkey, chainID)
	require.NoError(t, err)
	return txOpts
}

// MineTransaction blocks until the transaction is included in a block and returns the
// receipt. Errors are checked including the status.
func MineTransaction(t *testing.T, ec bind.DeployBackend, tx *types.Transaction) *types.Receipt {
	ctx := context.Background() // Create a MineTransactionWithCtx if a future test needs a custom context
	receipt, err := bind.WaitMined(ctx, ec, tx)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status) // Make sure the transaction was not reverted
	return receipt
}
