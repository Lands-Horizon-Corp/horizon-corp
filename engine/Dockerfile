# Stage 1: Development Environment
FROM golang:1.23-alpine AS development

# Set the working directory
WORKDIR /app

# Install dependencies and tools
RUN apk update && apk add --no-cache \
  bash \
  git \
  curl \
  ffmpeg \
  build-base

# Install Air for live reload and Swag for Swagger documentation
RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go.mod and go.sum first to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary (optional for dev, required for initial setup)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Expose the application port
EXPOSE 8080

# Command to run Air with the specified configuration
CMD ["air", "-c", ".air.toml"]
