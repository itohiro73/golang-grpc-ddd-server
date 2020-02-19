package infrastructure

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var catMethodList = map[string][]string{
	"/Cat/FindCuteCat": {"FindCuteCat"},
}

type User struct {
	permissions []string
}

// 本来はDB等からユーザーを検索する想定
func findUser(id string) *User {
	// 本番ではDBから取得する。
	switch id {
	case "1":
		return &User{permissions: []string{"FindCuteCat"}}
	case "2":
		return &User{permissions: []string{"CreateCat"}}
	}
	return &User{}
}

func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// ヘルスチェックの場合は無条件で認可する
		if info.FullMethod == "/grpc.health.v1.Health/Check" {
			return handler(ctx, req)
		}

		token, ok := GetAuthToken(ctx)
		if ok == true {
			// 認可処理としてはかなり雑だがサンプルなのでこれでいく
			user := findUser(token.Subject)
			if canAccessToMethod(info.FullMethod, user) {
				return handler(ctx, req)
			}
		}

		return nil, status.Error(
			codes.PermissionDenied,
			"could not access to specified method",
		)
	}
}

func canAccessToMethod(method string, user *User) bool {
	r, ok := catMethodList[method]
	if !ok {
		return false
	}

	permissions := map[string]bool{}
	for _, p := range user.permissions {
		permissions[p] = true
	}

	for _, p := range r {
		if !permissions[p] {
			return false
		}
	}

	return true
}
