.PHONY: build

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/dco_darwin_amd64 && \
	GOOS=linux GOARCH=amd64 go build -o bin/dco_linux_amd64