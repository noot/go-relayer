package gsnforwarder

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	errInvalidForwarderContract = errors.New("given contract address does not contain correct Forwarder code")
)

// CheckForwarderContractCode checks that the trusted forwarder contract used by
// the given swap contract has the expected bytecode.
func CheckForwarderContractCode(
	ctx context.Context,
	ec *ethclient.Client,
	contractAddr ethcommon.Address,
) error {
	code, err := ec.CodeAt(ctx, contractAddr, nil)
	if err != nil {
		return err
	}

	return checkContractBytes(code)
}

func checkContractBytes(contractCode []byte) error {
	expectedCode, err := hex.DecodeString(ExpectedForwarderByteCodeHex)
	if err != nil {
		panic(err) // We are decoding a constant, not reachable
	}

	if len(contractCode) < len(expectedCode) {
		return errInvalidForwarderContract
	}

	if !bytes.Equal(expectedCode, contractCode[0:len(expectedCode)]) {
		return errInvalidForwarderContract
	}

	return nil
}
