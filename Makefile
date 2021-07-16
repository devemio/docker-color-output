GO = $(shell which go)
GOBIN = $(shell $(GO) env GOPATH)/bin
BINPATH = $(CURDIR)/bin

.PHONY: build

build:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BINPATH)/dco-darwin-amd64 && \
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BINPATH)/dco-linux-amd64

deps_lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(GOBIN) v1.41.1

lint:
	golangci-lint run
