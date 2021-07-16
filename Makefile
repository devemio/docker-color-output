.PHONY: build lint cleanup

GO=$(shell which go)
BIN=$(CURDIR)/bin

build:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BIN)/dco-darwin-amd64 && \
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BIN)/dco-linux-amd64

lint:
	@docker run --rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint golangci-lint run

cleanup:
	@rm -f $(BIN)/dco-* && echo "\033[0;32mDone.\033[0m"