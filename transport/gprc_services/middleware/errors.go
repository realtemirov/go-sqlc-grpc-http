package middleware

import (
	"context"

	grpcerrors "github.com/realtemirov/go-sqlc-grpc-http/pkg/grpc_errors"
	"google.golang.org/grpc"
)

func GrpcErrorMiddleware(
	ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (
	interface{}, error,
) {
	resp, err := handler(ctx, req)
	if err != nil {
		return resp, grpcerrors.GetGrpcError(err)
	}

	return resp, nil
}
