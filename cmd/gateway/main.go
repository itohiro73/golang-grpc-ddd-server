package main

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/keitakn/golang-grpc-server/google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc"
	"net/http"
	"os"
)

func run() error {
	grpcServerEndpoint := "localhost:9998"
	if os.Getenv("DEPLOY_STAGE") == "local" || os.Getenv("DEPLOY_STAGE") == "" {
		// ローカル環境の場合は docker-compose.yml に書いてあるgRPCサーバーのコンテナ名を指定する
		grpcServerEndpoint = "grpc-server:9998"
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHealthHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
