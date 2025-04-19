# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/app ./main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /usr/local/bin/app .

# Copy .env file if needed
COPY .env .

# Expose the port your app runs on
EXPOSE 4000

# Command to run the executable
CMD ["./app"]