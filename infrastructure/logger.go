package infrastructure

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"strings"
)

func CreateLogger() *zap.Logger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	zapConfig := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ := zapConfig.Build()

	return logger
}

func AccessLogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		clientIP := "unknown"
		if p, ok := peer.FromContext(ctx); ok {
			clientIP = p.Addr.String()
		}

		ua := ""
		authToken := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if u, ok := md["user-agent"]; ok {
				ua = strings.Join(u, ",")
			}

			if token, ok := md["authorization"]; ok {
				authToken = token[0]
			}
		}

		ctxzap.AddFields(
			ctx,
			zap.String("access.clientip", clientIP),
			zap.String("access.useragent", ua),
			zap.String("access.authToken", authToken),
		)

		return handler(ctx, req)
	}
}
