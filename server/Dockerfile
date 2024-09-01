# Use the official Golang image
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go build -o main ./cmd/server/main.go

FROM golang:1.23

WORKDIR /app

# Copy the binary and other required files from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/config/config.yaml ./config/config.yaml
COPY --from=builder /app/tmp ./tmp

COPY --from=builder /go/bin/air /usr/local/bin/air

EXPOSE 8080

CMD ["air"]
