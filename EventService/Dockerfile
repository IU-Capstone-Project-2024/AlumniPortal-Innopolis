FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY EventService/go.mod EventService/go.sum ./

COPY shared/go.mod shared/

RUN go mod tidy
RUN go mod download

COPY EventService/ .
COPY shared/ ./shared

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

RUN chmod +x /app/main

EXPOSE 8084

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]