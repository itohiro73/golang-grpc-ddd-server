package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/keitakn/golang-grpc-server/infrastructure"
	pb "github.com/keitakn/golang-grpc-server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":9998")

	if err != nil {
		log.Fatalf("failed to listen port: %v", err)
	}

	zapLogger := infrastructure.CreateLogger()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpczap.UnaryServerInterceptor(zapLogger),
			grpc_auth.UnaryServerInterceptor(infrastructure.Authentication),
			infrastructure.AuthorizationUnaryServerInterceptor(),
		)),
	)

	catService := &infrastructure.CatService{}

	// 実行したい実処理をseverに登録する
	pb.RegisterCatServer(server, catService)

	// gRPCサーバのサービスの内容を公開
	reflection.Register(server)

	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
