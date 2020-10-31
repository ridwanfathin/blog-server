FROM golang:alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make build

FROM alpine:latest

RUN apk update && apk upgrade

WORKDIR /app

EXPOSE 8000