APP_NAME = server

run:
	go run ./cmd/${APP_NAME}

air:
	air

wire:
	wire ./internal/wire

test:
	go test ./...

test-cover:
	go test ./... --coverprofile=coverage.out

test-cover-html:
	go tool cover -html=coverage.out -o coverage.html