.PHONY: build lint clean

GO=$(shell which go)
BIN=$(CURDIR)/bin
APP=dco

build:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BIN)/$(APP)-darwin-amd64 && \
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BIN)/$(APP)-linux-amd64 && \
	echo "\033[0;32mDone.\033[0m"

lint:
	@docker run --rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint golangci-lint run

clean:
	@rm -f $(BIN)/$(APP)-*
