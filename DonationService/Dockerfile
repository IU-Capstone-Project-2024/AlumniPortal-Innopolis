FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY DonationService/go.mod DonationService/go.sum ./

COPY shared/go.mod shared/

RUN go mod tidy
RUN go mod download

COPY DonationService/ .
COPY shared/ ./shared

RUN go build -o main .

RUN chmod +x /app/main

EXPOSE 8083

CMD ["./main"]