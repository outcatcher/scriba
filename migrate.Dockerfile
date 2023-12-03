FROM golang:1.21.4-alpine3.17 AS builder

WORKDIR /opt/build

COPY . .

RUN go build -trimpath -o ./migrate ./cmd/migrate/main.go

FROM alpine:3.17.2

WORKDIR /opt/migrate

COPY --from=builder /opt/build/migrate ./migrate
COPY internal/migrations ./internal/migrations
COPY ./config/config.yaml ./config.yaml

CMD ["./migrate", "-config", "./config.yaml"]
