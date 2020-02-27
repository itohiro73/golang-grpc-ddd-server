FROM golang:1.14-alpine3.11 as build

WORKDIR /tmp

ENV PROTOBUF_VERSION 3.11.4
ENV PROTOBUF_URL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/protobuf-cpp-$PROTOBUF_VERSION.tar.gz

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl build-base autoconf automake libtool && \
  curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL && \
  tar xvzf protobuf.tar.gz

WORKDIR /tmp/protobuf-$PROTOBUF_VERSION

RUN set -eux && \
  ./autogen.sh && \
  ./configure && \
  make -j 3 && \
  make install && \
  go get -u github.com/golang/protobuf/protoc-gen-go && \
  go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc && \
  go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
  go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

WORKDIR /go/app

COPY . .

RUN set -eux && \
  go build -o golang-grpc-server ./cmd/grpc/main.go && \
  go get gopkg.in/urfave/cli.v2@master && \
  go get github.com/oxequa/realize && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

FROM alpine:3.11

WORKDIR /app

COPY --from=build /go/app/golang-grpc-server .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/golang-grpc-server

CMD ["./golang-grpc-server"]
