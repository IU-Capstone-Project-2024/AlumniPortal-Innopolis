FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY ProjectService/go.mod ProjectService/go.sum ./

COPY shared/go.mod shared/

RUN go mod tidy
RUN go mod download

COPY ProjectService/ .
COPY shared/ ./shared

RUN go build -o main .

EXPOSE 8085

CMD ["./main"]