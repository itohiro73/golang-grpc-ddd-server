package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/itohiro73/golang-grpc-ddd-server/internal/infrastructure"
	"github.com/itohiro73/golang-grpc-ddd-server/internal/infrastructure/persistence"
	"github.com/itohiro73/golang-grpc-ddd-server/internal/interfaces/grpc/handler"
	"github.com/itohiro73/golang-grpc-ddd-server/internal/usecase"
	"github.com/itohiro73/golang-grpc-ddd-server/pkg/pb/cat"
	"github.com/itohiro73/golang-grpc-ddd-server/pkg/pb/dog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	health "github.com/itohiro73/golang-grpc-ddd-server/google.golang.org/grpc/health/grpc_health_v1"
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
			infrastructure.AccessLogUnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(infrastructure.Authentication),
			infrastructure.AuthorizationUnaryServerInterceptor(),
		)),
	)

	catService := &infrastructure.CatService{}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	dogPersistence := persistence.NewDogPersistence(db)
	dogUseCase := usecase.NewDogUseCase(dogPersistence)
	dogHandler := handler.NewDogHandler(dogUseCase)
	// 実行したい実処理をseverに登録する
	cat.RegisterCatServer(server, catService)
	dog.RegisterDogServer(server, dogHandler)

	// ヘルスチェック用のメソッド
	healthCheckService := &infrastructure.SkipAuthHealthServer{}
	health.RegisterHealthServer(server, healthCheckService)

	// gRPCサーバのサービスの内容を公開
	reflection.Register(server)

	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
