.PHONY: build

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/dco-darwin-amd64 && \
	GOOS=linux GOARCH=amd64 go build -o bin/dco-linux-amd64