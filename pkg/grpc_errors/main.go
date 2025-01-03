package grpcerrors

import (
	"errors"
	"runtime"

	xerrors "github.com/realtemirov/go-sqlc-grpc-http/config/x_errors"
	"github.com/realtemirov/go-sqlc-grpc-http/generated/general"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCStatusError interface {
	GRPCStatus() *status.Status
}

func GetGrpcError(platformError error) error {
	var grpcErr gRPCStatusError
	if errors.As(platformError, &grpcErr) {
		return platformError
	}

	rootErr := parseRootError(platformError)

	if rootErr.notFound {
		zap.L().Error("error message for the method is not found", zap.Error(rootErr))
		rootErr.Function = xerrors.InternalError
	}

	errorCodes := xerrors.GlobalErrors[rootErr.Function]
	grpcErrorDefinition := errorCodes[rootErr.Code]

	mainSt, err := status.
		New(grpcErrorDefinition.Code, grpcErrorDefinition.Message).
		WithDetails(&general.ErrorInfo{
			ErrorCode: grpcErrorDefinition.ErrorCode,
			Message:   platformError.Error(),
			Label: &general.ErrorInfo_Label{
				Uz: grpcErrorDefinition.Labels.Uz,
				Ru: grpcErrorDefinition.Labels.Ru,
				En: grpcErrorDefinition.Labels.En,
			},
		})
	if err != nil {
		return status.Errorf(codes.Internal, "failed to create grpc error: %v", err)
	}

	return mainSt.Err()
}

func parseRootError(err error) Error {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()

	var rootError Error
	if matches := errors.As(err, &rootError); !matches {
		return Error{
			Code:     xerrors.NotFound,
			Function: funcName,
			Err:      err,
			notFound: true,
		}
	}

	return rootError
}
