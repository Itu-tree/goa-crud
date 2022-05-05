FROM golang:1.18

RUN apt-get update  \
    && apt-get install unzip \
    && apt-get install -y protobuf-compiler


WORKDIR /go/src/todo

COPY go.mod go.sum ./
RUN go mod download

ENV GO111MODULE on
