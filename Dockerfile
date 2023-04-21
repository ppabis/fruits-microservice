FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . /app/

RUN go build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/fruits_microservice /usr/bin/fruits_microservice

EXPOSE 8081

ENTRYPOINT [ "/usr/bin/fruits_microservice" ]