# Use an official Golang runtime as a parent image
FROM golang:1.22.4-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal Docker image for the runtime
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Copy configuration files
COPY ./routes.yaml .
COPY ./api-keys.yaml .
COPY ./redis.yaml .

# Expose the application port (if applicable, change accordingly)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
