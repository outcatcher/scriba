FROM golang:1.20-alpine3.17 AS builder

WORKDIR /opt/build

COPY . .

RUN go build -trimpath -o ./migrate ./cmd/migrate/main.go

FROM alpine:3.17.2

WORKDIR /opt/migrate

COPY --from=builder /opt/build/migrate ./migrate
COPY migrations ./migrations
COPY ./config.yaml ./config.yaml

CMD ["./migrate", "-config", "./config.yaml"]
