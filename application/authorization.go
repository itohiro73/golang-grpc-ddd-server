package application

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		token, ok := GetAuthToken(ctx)
		if ok == true {
			// 認可処理としてはかなり雑だがサンプルなのでこれでいく
			// 本来は info.FullMethod の値などを見て、対象のリクエストを許可して良いか判断する事になる
			log.Print(token)
			log.Println(info.FullMethod)
			return handler(ctx, req)
		}

		return nil, status.Error(
			codes.PermissionDenied,
			"could not access to specified method",
		)
	}
}
