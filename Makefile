GOPATH ?= $(shell go env GOPATH)

.PHONY: lint
lint:
	bash scripts/install-lint.sh
	${GOPATH}/bin/golangci-lint run

test:
	go test ./...

build:
	mkdir bin
	cd cmd && go build -o ../bin/relayer 