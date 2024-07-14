FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY VolunteerService/go.mod VolunteerService/go.sum ./

COPY shared/go.mod shared/

RUN go mod tidy
RUN go mod download

COPY VolunteerService/ .
COPY shared/ ./shared

RUN go build -o main .

RUN chmod +x /app/main

EXPOSE 8086

CMD ["./main"]