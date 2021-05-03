FROM golang:1.16-alpine as builder
WORKDIR /go/src/github.com/moov-io/iso20022
RUN apk add -U make
COPY . .
RUN make build
