# Use a more compatible Go runtime as a parent image
FROM golang:1.23.1 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./bin/server.exe

# Start a new stage from scratch
FROM debian:bullseye-slim

# Install necessary dependencies for running the app, including newer glibc
RUN apt-get update && apt-get install -y \
  ca-certificates \
  libc6 \
  && rm -rf /var/lib/apt/lists/*

# Set environment variables
ENV PORT=8080
ENV CONN_STR=postgres://postgres:log123@host.docker.internal:9696/clockbakers?sslmode=disable
ENV SECRET_KEY=manojgandu69420@yahoo.com

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/bin/server.exe .

# Copy the swagger file
COPY ./swagger.json /root/swagger.json


# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./server.exe"]