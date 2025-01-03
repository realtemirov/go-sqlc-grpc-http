package middleware

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func GrpcLoggerMiddleware(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (
	interface{}, error,
) {
	defer func() {
		if errSync := zap.L().Sync(); errSync != nil {
			zap.L().Error("failed zap.Log.Sync", zap.Error(errSync))
		}
	}()

	mappedBody, err := convertToMap(req)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	redactSensitiveInfo(mappedBody)

	zap.L().Info(
		"-----> Service: ",
		zap.String("method", info.FullMethod),
		zap.Any("body", mappedBody),
	)

	resp, err := handler(ctx, req)

	if err != nil {
		zap.L().Error("failed to make gRPC call: ", zap.Error(err))
		return resp, err
	}

	zap.L().Info("<----- Service:")

	return resp, nil
}
