# Build stage
FROM golang:alpine AS builder

WORKDIR /app

# Install git and other dependencies if needed (useful for Go modules)
RUN apk add --no-cache git

# Copy go mod and sum files before copying the rest to leverage caching
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./main.go

# Final stage
FROM alpine

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app runs on
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./main"]
