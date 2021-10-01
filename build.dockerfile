FROM golang:1.17-alpine
WORKDIR /root
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN apk add musl-dev gcc