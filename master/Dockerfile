FROM golang:alpine3.7 AS builder
RUN apk add --no-cache --virtual .run-deps \
    ca-certificates bash wget git openssl \
    && update-ca-certificates

RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/muchrm/docker-manager/master
ADD . .
RUN dep ensure
RUN go build -o master

FROM alpine:3.7
WORKDIR /app
COPY --from=builder /go/src/github.com/muchrm/docker-manager/master /app
CMD ["/app/manager"]