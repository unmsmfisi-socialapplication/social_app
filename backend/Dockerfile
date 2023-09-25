FROM golang:1.21.1-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN go build -o build/main cmd/server.go

CMD ["build/main"]
