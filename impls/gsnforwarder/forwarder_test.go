package gsnforwarder

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/go-relayer/common"
	"github.com/athanorlabs/go-relayer/tests"
)

// expectedCBORMetadataLen is the length of the metadata in the deployed opengsn
// forwarder contracts. We match the solc version and flags. Although we are
// unable to get a 100% match to the source code published on etherscan
// (etherscan doesn't verify metadata), we are producing matching bytecode and
// even a matching metadata size.
const expectedCBORMetadataLen = uint16(51)

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

func getMainnetForwarderByteCode(t *testing.T) []byte {
	ctx := context.Background()
	ec := getMainnetClient(t)

	code, err := ec.CodeAt(ctx, ethcommon.HexToAddress(MainnetForwarderAddrHex), nil)
	require.NoError(t, err)

	return code
}

func getOurForwarderByteCode(t *testing.T) ([]byte, ethcommon.Address) {
	ec, _ := tests.NewEthClient(t)
	pk := tests.GetTestKeyByIndex(t, 0)
	auth := tests.NewTXOpts(t, ec, pk)

	address, tx, _, err := DeployForwarder(auth, ec)
	require.NoError(t, err)
	tests.MineTransaction(t, ec, tx)

	ourByteCode, err := ec.CodeAt(context.Background(), address, nil)
	require.NoError(t, err)

	return ourByteCode, address
}

// Test that our ExpectedForwarderByteCodeHex matches the contract bytecode
// when deploy it.
func TestExpectedForwarderByteCodeHex(t *testing.T) {
	// The metadata appended to the end is CBOR encoded. The length of the data
	// is in 2 bytes at the very end. We need to remove N+2 bytes, where N is
	// the value big endian value in the last 2 bytes.
	byteCode, _ := getOurForwarderByteCode(t)
	cborMetaDataLen := binary.BigEndian.Uint16(byteCode[len(byteCode)-2:])
	require.Equal(t, expectedCBORMetadataLen, cborMetaDataLen)
	byteCodeWithoutMeta := byteCode[0 : len(byteCode)-int(cborMetaDataLen)-2]
	require.Equal(t, ExpectedForwarderByteCodeHex, hex.EncodeToString(byteCodeWithoutMeta))
	encodedMetaData := byteCode[len(byteCodeWithoutMeta):]
	t.Logf("contract metadata can be decoded using https://playground.sourcify.dev/: %x", encodedMetaData)
}

func TestOurCompiledForwarderMatchesMainnet(t *testing.T) {
	ourByteCode, _ := getOurForwarderByteCode(t)
	mainnetByteCode := getMainnetForwarderByteCode(t)

	// The metadata won't match, but the way we've set things
	// up, even the metadata size should match.
	require.Equal(t, len(mainnetByteCode), len(ourByteCode))
	require.Equal(t, expectedCBORMetadataLen, binary.BigEndian.Uint16(ourByteCode[len(mainnetByteCode)-2:]))
	require.Equal(t, expectedCBORMetadataLen, binary.BigEndian.Uint16(mainnetByteCode[len(mainnetByteCode)-2:]))

	lenWithoutMeta := len(mainnetByteCode) - int(expectedCBORMetadataLen) - 2 // 2 bytes for the length
	t.Logf("our metadata:     %x", ourByteCode[lenWithoutMeta:])
	t.Logf("mainnet metadata: %x", mainnetByteCode[lenWithoutMeta:])
	require.Equal(t, mainnetByteCode[:lenWithoutMeta], ourByteCode[:lenWithoutMeta])
}
