FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY EventService/go.mod EventService/go.sum ./

COPY shared/go.mod shared/

RUN go mod tidy
RUN go mod download

COPY EventService/ .
COPY shared/ ./shared

RUN go build -o main .

EXPOSE 8084

CMD ["./main"]