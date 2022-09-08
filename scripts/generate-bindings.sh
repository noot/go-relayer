#!/bin/bash

# Use the project root (one directory above this script) as the current working directory:
PROJECT_ROOT="$(dirname "$(dirname "$(readlink -f "$0")")")"
cd "${PROJECT_ROOT}" || exit 1

ABIGEN="$(go env GOPATH)/bin/abigen"

if [[ -z "${SOLC_BIN}" ]]; then
	SOLC_BIN=solc
fi

"${SOLC_BIN}" --abi contracts/contracts/Mock.sol -o contracts/abi/ --overwrite
"${SOLC_BIN}" --bin contracts/contracts/Mock.sol -o contracts/bin/ --overwrite

"${ABIGEN}" \
	--abi contracts/abi/Mock.abi \
	--bin contracts/bin/Mock.bin \
	--pkg contracts \
	--type Mock \
	--out contracts/mock.go

"${SOLC_BIN}" --abi contracts/contracts/MinimalForwarder.sol -o contracts/abi/ --overwrite
"${SOLC_BIN}" --bin contracts/contracts/MinimalForwarder.sol -o contracts/bin/ --overwrite

"${ABIGEN}" \
	--abi contracts/abi/MinimalForwarder.abi \
	--bin contracts/bin/MinimalForwarder.bin \
	--pkg contracts \
	--type MinimalForwarder \
	--out contracts/forwarder.go
