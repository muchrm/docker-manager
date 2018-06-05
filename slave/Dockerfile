FROM golang:alpine3.7 AS builder
ADD . /app
WORKDIR /app
RUN go build -o http

FROM alpine:3.7
WORKDIR /app
EXPOSE 8000
COPY --from=builder /app/http /app
CMD ["/app/http"]