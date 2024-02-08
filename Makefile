BIN=$(CURDIR)/bin
GO=$(shell which go)
APP=docker-color-output

.PHONY: build
build:
	@$(GO) build -o $(BIN)/$(APP) ./cmd/cli

.PHONY: publish
publish: clean
	@GOOS=darwin GOARCH=amd64 $(GO) build -o $(BIN)/$(APP)-darwin-amd64 ./cmd/cli && \
		GOOS=linux GOARCH=amd64 $(GO) build -o $(BIN)/$(APP)-linux-amd64 ./cmd/cli && \
		GOOS=windows GOARCH=amd64 $(GO) build -o $(BIN)/$(APP)-windows-amd64.exe ./cmd/cli && \
		mv ./bin/* ~/Downloads

.PHONY: test
test:
	@go test ./...

.PHONY: clean
clean:
	@rm -f $(BIN)/$(APP)*

.PHONY: lint
lint:
	@golangci-lint run -v --fix
