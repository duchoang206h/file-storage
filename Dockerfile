FROM golang:1.21-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/static/ ./static


COPY .env .

EXPOSE 4000

CMD ["./app"]