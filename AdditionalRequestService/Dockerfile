FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY AdditionalRequestService/go.mod AdditionalRequestService/go.sum ./

COPY shared/go.mod shared/

RUN go mod tidy
RUN go mod download

COPY AdditionalRequestService/ .
COPY shared/ ./shared

RUN go build -o main .

RUN chmod +x /app/main

EXPOSE 8082

CMD ["./main"]