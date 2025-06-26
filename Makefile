.PHONY: build test lint install

BINARY_NAME=crossplane-render

build:
	go build -o $(BINARY_NAME) ./cmd/render/main

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/render/main
