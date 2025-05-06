# Project settings
CMD_DIR=cmd/parser
BIN_NAME=go-web-scraper
OUTPUT_BIN=./bin/$(BIN_NAME)

# Default target
.PHONY: all
all: tidy vendor build

## Run the parser from source
.PHONY: run
run:
	go run $(CMD_DIR)/main.go

## Build the binary
.PHONY: build
build:
	mkdir -p ./bin
	go build -o $(OUTPUT_BIN) $(CMD_DIR)/main.go

## Vendor dependencies
.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

## Clean binaries and vendor
.PHONY: clean
clean:
	rm -rf ./bin
	rm -rf ./vendor

## Run tests
.PHONY: test
test:
	go test ./...

## Format code
.PHONY: fmt
fmt:
	go fmt ./...
