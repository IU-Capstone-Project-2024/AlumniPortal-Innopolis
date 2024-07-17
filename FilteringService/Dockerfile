FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY FilteringService/go.mod FilteringService/go.sum ./

COPY shared/go.mod shared/

RUN go mod tidy
RUN go mod download

COPY FilteringService/ .
COPY shared/ ./shared

RUN go build -o main .

RUN chmod +x /app/main

EXPOSE 50051

CMD ["./main"]