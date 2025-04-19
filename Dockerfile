FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o app ./cmd/server

FROM scratch

COPY ./config/production.yaml /config/local.yaml

COPY --from=builder /build/app /

EXPOSE 8002

CMD ["/app"]
