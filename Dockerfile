FROM golang:1.21

RUN apt-get update && apt-get install -y postgresql-client

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

EXPOSE 8080

CMD ["/app/main"]