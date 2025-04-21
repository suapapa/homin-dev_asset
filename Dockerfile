FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY *.go .

RUN go build -o ./app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY asset ./asset

CMD ["./app"]