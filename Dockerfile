# Stage 1: Build the Go application
FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod  .
COPY go.sum  .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

#Stage 2: use minimal resource to run binary
FROM alpine:latest

WORKDIR /app
RUN touch .env
COPY --from=builder /app/myapp .


CMD ["./myapp"]