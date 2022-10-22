#!/bin/bash

# Use the project root (one directory above this script) as the current working directory:
PROJECT_ROOT="$(dirname "$(dirname "$(readlink -f "$0")")")"
cd "${PROJECT_ROOT}" || exit 1

ABIGEN="$(go env GOPATH)/bin/abigen"

if [[ -z "${SOLC_BIN}" ]]; then
	SOLC_BIN=solc
fi

"${SOLC_BIN}" --abi contracts/MinimalForwarder.sol -o contracts/abi/ --overwrite
"${SOLC_BIN}" --bin contracts/MinimalForwarder.sol -o contracts/bin/ --overwrite

"${ABIGEN}" \
	--abi contracts/abi/MinimalForwarder.abi \
	--bin contracts/bin/MinimalForwarder.bin \
	--pkg mforwarder \
	--type MinimalForwarder \
	--out tmp_forwarder.go

"${SOLC_BIN}" --abi contracts/IMinimalForwarder.sol -o contracts/abi/ --overwrite
"${SOLC_BIN}" --bin contracts/IMinimalForwarder.sol -o contracts/bin/ --overwrite

"${ABIGEN}" \
	--abi contracts/abi/IMinimalForwarder.abi \
	--bin contracts/bin/IMinimalForwarder.bin \
	--pkg mforwarder \
	--type IMinimalForwarder \
	--out i_minimal_forwarder.go

sed '31,40d' tmp_forwarder.go > minimal_forwarder.go
rm tmp_forwarder.go