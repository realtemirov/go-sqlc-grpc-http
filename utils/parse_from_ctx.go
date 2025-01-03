package utils

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

const (
	IPAddress           = "x-forwarded-for"
	UserAgent           = "grpcgateway-user-agent"
	Platform            = "x-platform"
	APIKey              = "x-api-key"
	GrpcGatewayHTTPPath = "grpcgateway-http-path"
	GrpcGaewayMethod    = "grpcgateway-method"
	AuthorizationHeader = "Authorization"
)

func ParseKeyFromCtx(ctx context.Context, key string) string {
	md, _ := metadata.FromIncomingContext(ctx)
	keyValue := md.Get(key)

	if len(keyValue) == 0 {
		zap.L().Warn("Key not found in context", zap.String("key", key))
		return ""
	}

	return keyValue[0]
}
