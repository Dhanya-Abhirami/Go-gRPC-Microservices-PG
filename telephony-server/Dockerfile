FROM golang

RUN apt-get update \
        apt-get install protobuf-compiler

WORKDIR /app

ADD . .

RUN GO111MODULE=on \
        go get github.com/golang/protobuf/proto \
        github.com/golang/protobuf/protoc-gen-go \
        google.golang.org/grpc