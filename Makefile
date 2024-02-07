BIN=$(CURDIR)/bin
GO=$(shell which go)
APP=docker-color-output

.PHONY: build
build:
	@$(GO) build -o $(BIN)/$(APP) ./cmd/cli

.PHONY: test
test:
	@go test ./...

.PHONY: clean
clean:
	@rm -f $(BIN)/$(APP)*

.PHONY: lint
lint:
	@golangci-lint run -v --fix
