FROM golang:1.21.1-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod verify

RUN go build -o build/backend-develop cmd/server.go

EXPOSE 3333

CMD ["build/backend-develop"]
