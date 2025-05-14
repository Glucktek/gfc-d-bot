# Build stage
FROM golang:1.24.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bot .

# Runtime stage
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/bot .
CMD ["./bot"]
