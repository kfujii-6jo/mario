# syntax=docker/dockerfile:1
FROM golang:1.23
WORKDIR /app

COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /run-api cmd/mario/main.go

EXPOSE 8080

CMD ["/run-api"]
