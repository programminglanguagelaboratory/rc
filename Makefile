all: build test

build:
	@go build -o build/ ./...

test:
	@go test ./...

.PHONY: build test
