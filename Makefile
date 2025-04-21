APP_NAME = server
BINARY_NAME=tmp/main

run:
	go run ./cmd/${APP_NAME}

watch:
	air

clean:
	rm -f ./$(BINARY_NAME)*

wiregen:
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

create-migration:
	goose -env .env.local create ${name} sql

up_by_one:
	goose -env .env.local up-by-one

sqlgen:
	sqlc generate

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: run watch clean wire test test-cover test-cover-html upse downse resetse sqlc