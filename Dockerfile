# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o northstar .

# Runtime Stage
FROM alpine:latest

WORKDIR /app

# Copy executable
COPY --from=builder /app/northstar .

# Copy environment file (optional for development)
COPY .env .

EXPOSE 8080

CMD ["./northstar"]