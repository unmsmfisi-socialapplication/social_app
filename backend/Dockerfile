
WORKDIR /app

COPY . .

RUN go mod download

RUN go mod verify

RUN go build -o build/backend-develop cmd/server.go

EXPOSE 3333

CMD ["build/backend-develop"]
