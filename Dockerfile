FROM golang:1.20-alpine

COPY . /app

WORKDIR /app

EXPOSE 5000

CMD go run server.go