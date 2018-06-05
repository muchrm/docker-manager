FROM golang:alpine3.7 AS builder

WORKDIR /app
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o manager

FROM alpine:3.7
WORKDIR /app
COPY --from=builder /app/manager /app
CMD ["/app/manager"]