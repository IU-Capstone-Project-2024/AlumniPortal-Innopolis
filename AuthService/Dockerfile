FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY AuthService/go.mod AuthService/go.sum ./

COPY shared/go.mod shared/

RUN go mod tidy
RUN go mod download

COPY AuthService/ .
COPY shared/ ./shared
COPY certs ./certs

RUN go build -o main .

RUN chmod +x /app/main

EXPOSE 8081

CMD ["./main"]
