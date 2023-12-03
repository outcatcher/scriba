FROM golang:1.21.4-alpine3.17 AS builder

WORKDIR /opt/build

COPY . .

RUN go build -trimpath -o ./scriba ./cmd/bot/main.go


FROM alpine:3.17.2

WORKDIR /opt/scriba

COPY --from=builder /opt/build/scriba ./scriba
COPY ./config/config.yaml ./config.yaml

CMD ["./scriba", "-config", "./config.yaml"]
