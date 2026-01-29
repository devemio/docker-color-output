BIN=$(CURDIR)/bin
GO=$(shell which go)
APP=docker-color-output

.PHONY: build
build:
	@CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o $(BIN)/$(APP) ./cmd/docker-color-output

.PHONY: publish
publish: clean
	# Mac
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o $(BIN)/$(APP)-darwin-amd64 ./cmd/docker-color-output
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o $(BIN)/$(APP)-darwin-arm64 ./cmd/docker-color-output
	# Linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o $(BIN)/$(APP)-linux-amd64 ./cmd/docker-color-output
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o $(BIN)/$(APP)-linux-arm64 ./cmd/docker-color-output
	# Windows
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o $(BIN)/$(APP)-windows-amd64.exe ./cmd/docker-color-output
	# Cleanup
	mv ./bin/* ~/Downloads

.PHONY: test
test:
	@go test ./...

.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

.PHONY: clean
clean:
	@rm -f $(BIN)/$(APP)*

.PHONY: lint
lint:
	@golangci-lint run -v --fix
