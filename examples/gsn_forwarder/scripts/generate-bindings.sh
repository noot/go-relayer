#!/bin/bash

# Use the project root (one directory above this script) as the current working directory:
PROJECT_ROOT="$(dirname "$(dirname "$(readlink -f "$0")")")"
cd "${PROJECT_ROOT}" || exit 1

ABIGEN="$(go env GOPATH)/bin/abigen"

PKG_NAME="gsnforwarder"

if [[ -z "${SOLC_BIN}" ]]; then
	SOLC_BIN=solc
fi

"${SOLC_BIN}" --abi contracts/Forwarder.sol -o contracts/abi/ --overwrite
"${SOLC_BIN}" --bin contracts/Forwarder.sol -o contracts/bin/ --overwrite

"${ABIGEN}" \
	--abi contracts/abi/Forwarder.abi \
	--bin contracts/bin/Forwarder.bin \
	--pkg ${PKG_NAME} \
	--type Forwarder \
	--out forwarder.go

"${SOLC_BIN}" --abi contracts/IForwarder.sol -o contracts/abi/ --overwrite
"${SOLC_BIN}" --bin contracts/IForwarder.sol -o contracts/bin/ --overwrite

"${ABIGEN}" \
	--abi contracts/abi/IForwarder.abi \
	--bin contracts/bin/IForwarder.bin \
	--pkg ${PKG_NAME} \
	--type IForwarder \
	--out i_forwarder.go
