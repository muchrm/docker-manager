FROM golang:alpine3.7 AS builder

RUN apk add --no-cache --virtual .run-deps \
    ca-certificates bash wget git openssl \
    && update-ca-certificates
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/muchrm/docker-manager/manager
ADD . .
RUN dep ensure
RUN go build -o manager

FROM alpine:3.7
WORKDIR /app
COPY --from=builder /go/src/github.com/muchrm/docker-manager/manager /app
CMD ["/app/manager"]