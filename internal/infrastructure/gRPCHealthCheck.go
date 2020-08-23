package infrastructure

import (
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	health "github.com/itohiro73/golang-grpc-ddd-server/google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SkipAuthHealthServer struct {
	health.HealthServer
}

var _ grpc_auth.ServiceAuthFuncOverride = (*SkipAuthHealthServer)(nil)

// AuthFuncOverride SkipAuthHealthServer構造体がServiceAuthFuncOverrideインターフェースを実装する
func (*SkipAuthHealthServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (h *SkipAuthHealthServer) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	// ここまで処理が来ればヘルスチェックとしては成功となる
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (h *SkipAuthHealthServer) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "service watch is not implemented current version.")
}
