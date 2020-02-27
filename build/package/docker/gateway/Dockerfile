FROM golang:1.14-alpine3.11 as build

WORKDIR /go/app

COPY . .

RUN set -eux && \
  go build -o grpc-gateway ./cmd/gateway/main.go

FROM alpine:3.11

WORKDIR /app

COPY --from=build /go/app/grpc-gateway .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/grpc-gateway

CMD ["./grpc-gateway"]
