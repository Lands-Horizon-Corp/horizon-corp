FROM golang:1.23-alpine AS development

WORKDIR /app

RUN apk update && apk add --no-cache \
  bash \
  git \
  curl \
  ffmpeg

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
