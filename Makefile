GOPATH ?= $(shell go env GOPATH)

.PHONY: build
build:
	mkdir -p bin
	GOBIN="$(CURDIR)/bin/" go install ./cmd/...

.PHONY: lint
lint:
	bash scripts/install-lint.sh
	${GOPATH}/bin/golangci-lint run --timeout 5m

.PHONY: format
format:
	test -x $(GOPATH)/bin/goimports || go install golang.org/x/tools/cmd/goimports@latest
	$(GOPATH)/bin/goimports -local github.com/athanorlabs/go-relayer -w .


.PHONY: test
test:
	go test -count=1 ./... -v -timeout=2m

.PHONY: clean
clean:
	rm -rf bin
