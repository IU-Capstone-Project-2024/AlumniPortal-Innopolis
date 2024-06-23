FROM golang:1.17-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

VOLUME ../shared:/shared

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
