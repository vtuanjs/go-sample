APP_NAME = server
BINARY_NAME=tmp/main

run:
	go run ./cmd/${APP_NAME}

watch:
	air

clean:
	rm -f ./$(BINARY_NAME)*

wire:
	wire ./internal/wire

test:
	go test ./...

test-cover:
	go test ./... --coverprofile=coverage.out

test-cover-html:
	go tool cover -html=coverage.out -o coverage.html

.PHONY: run air wire test test-cover test-cover-html