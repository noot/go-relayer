#!/bin/bash

# Change to the gsnforwarder dir, one directory up from this script
GSNFORWARDER_DIR="$(dirname "$(dirname "$(readlink -f "$0")")")"
cd "${GSNFORWARDER_DIR}" || exit 1

ABIGEN="$(go env GOPATH)/bin/abigen"

PKG_NAME="gsnforwarder"

# We need the solc version to match the contract deployed here:
# https://docs.opengsn.org/networks/ethereum/mainnet.html
# https://etherscan.io/address/0xB2b5841DBeF766d4b521221732F9B618fCf34A87#code
SOLC_VERSION="0.8.7"

SOLC_EXEC=(
		docker run --rm \
		--volume "${GSNFORWARDER_DIR}/contracts:/contracts" \
		-u "$(id -u):$(id -g)" \
		"ethereum/solc:${SOLC_VERSION}"
	)

set -x
"${SOLC_EXEC[@]}" --abi contracts/Forwarder.sol -o contracts/abi/ --overwrite
"${SOLC_EXEC[@]}" --bin contracts/Forwarder.sol -o contracts/bin/ --overwrite

"${ABIGEN}" \
	--abi contracts/abi/Forwarder.abi \
	--bin contracts/bin/Forwarder.bin \
	--pkg ${PKG_NAME} \
	--type Forwarder \
	--out forwarder.go

# IForwarderForwardRequest gets generated in both forwarder.go and i_forwarder.go
# with an identical definition.
sed -i 's/^type IForwarderForwardRequest /type _duplicate_IForwarderForwardRequest /' forwarder.go
