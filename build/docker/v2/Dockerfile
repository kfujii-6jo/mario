# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /run-api cmd/mario/main.go

FROM alpine:3.20

WORKDIR /root/
COPY --from=builder /run-api /run-api

EXPOSE 8080
CMD ["/run-api"]
