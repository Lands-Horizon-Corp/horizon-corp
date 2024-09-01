FROM golang:1.23-alpine as development

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

# Build the binary into the tmp directory within the Docker image
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tmp/main cmd/server/main.go

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
