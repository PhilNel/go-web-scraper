# Project settings
CMD_DIR=cmd/parser
BIN_NAME=go-web-scraper
OUTPUT_BIN=./bin/$(BIN_NAME)

.PHONY: run
run:
	go run $(CMD_DIR)/main.go

.PHONY: build
build:
	mkdir -p ./bin
	go build -o $(OUTPUT_BIN) $(CMD_DIR)/main.go

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run --config .golangci.yml
