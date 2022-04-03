.PHONY: build test clean lint
.SILENT: build test clean lint

BIN=$(CURDIR)/bin
GO=$(shell which go)
APP=docker-color-output

build:
	$(GO) build -o $(BIN)/$(APP) ./cmd/cli

test:
	go test -v ./...

clean:
	rm -f $(BIN)/$(APP)*

lint:
	docker run --rm -w /opt -v $(shell pwd):/opt golangci/golangci-lint golangci-lint run
