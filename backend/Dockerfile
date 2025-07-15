# Multi-stage build for Go
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod ./
COPY go.su[m] ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]