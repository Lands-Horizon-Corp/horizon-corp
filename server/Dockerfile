# Base image for development with Go and Air
FROM golang:1.23-alpine as development

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/main.go

# Expose port 8080
EXPOSE 8080

# Start Air for live reloading
CMD ["air", "-c", ".air.toml"]
