FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o app ./cmd/server

FROM scratch

COPY .env.prod .env.local

COPY --from=builder /build/app /

EXPOSE 8002

CMD ["/app"]
