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

upse:
	goose -env .env.local up

downse:
	goose -env .env.local down

resetse:
	goose -env .env.local reset

sqlc:
	sqlc generate

.PHONY: run watch clean wire test test-cover test-cover-html upse downse resetse sqlc