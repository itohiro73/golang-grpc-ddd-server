package infrastructure

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const AuthTokenKey = "auth-token"

func Authentication(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")

	if err != nil {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"could not read auth token: %v",
			err,
		)
	}

	// サンプルなので署名検証はスキップ
	// payloadにsubが含まれていれば問題なし
	parser := new(jwt.Parser)
	parsedToken, _, err := parser.ParseUnverified(token, &jwt.StandardClaims{})
	if err != nil {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"could not parsed auth token: %v",
			err,
		)
	}

	return setAuthToken(ctx, parsedToken.Claims.(*jwt.StandardClaims)), nil
}

func setAuthToken(ctx context.Context, token *jwt.StandardClaims) context.Context {
	return context.WithValue(ctx, AuthTokenKey, token)
}

func GetAuthToken(ctx context.Context) (*jwt.StandardClaims, bool) {
	val, ok := ctx.Value(AuthTokenKey).(*jwt.StandardClaims)
	return val, ok
}
