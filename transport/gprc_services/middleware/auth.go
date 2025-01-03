package middleware

// import (
// 	"context"
// 	"errors"
// 	"slices"
// 	"strings"
// 	"time"

// 	"github.com/realtemirov/auth_service/generated/auth"
// 	"github.com/realtemirov/screener_core_service/config/constants"
// 	"github.com/realtemirov/screener_core_service/pkg/libraries/token"
// 	utils "github.com/realtemirov/screener_core_service/util"
// 	"go.uber.org/zap"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/metadata"
// 	"google.golang.org/grpc/status"
// )

// type authMiddleware struct {
// 	authClient      auth.AuthServiceClient
// 	publicEndpoints map[string][]string
// }

// type AuthMiddleware interface {
// 	Auth(
// 		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
// 		handler grpc.UnaryHandler,
// 	) (interface{}, error)
// }

// const (
// 	authorizationForm = "bearer"
// )

// func NewAuthMiddleware(authClient auth.AuthServiceClient) AuthMiddleware {
// 	return &authMiddleware{
// 		authClient: authClient,
// 		publicEndpoints: map[string][]string{
// 			"/v1/auth/login":                  {"POST"},
// 			"/v1/auth/send-confirmation-code": {"POST"},
// 			"/v1/recommendation":              {"GET"},
// 			"/v1/todos":                       {"GET", "POST", "PUT", "DELETE"},
// 		},
// 	}
// }

// func (m *authMiddleware) Auth(
// 	ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo,
// 	handler grpc.UnaryHandler) (
// 	interface{}, error,
// ) {
// 	httpPath := utils.ParseKeyFromCtx(ctx, utils.GrpcGatewayHTTPPath)
// 	httpMethod := utils.ParseKeyFromCtx(ctx, utils.GrpcGaewayMethod)
// 	ipAddress := utils.ParseKeyFromCtx(ctx, utils.IPAddress)
// 	userAgent := utils.ParseKeyFromCtx(ctx, utils.UserAgent)
// 	platform := utils.ParseKeyFromCtx(ctx, utils.Platform)
// 	apiKey := utils.ParseKeyFromCtx(ctx, utils.APIKey)

// 	newCtx := metadata.NewOutgoingContext(ctx,
// 		metadata.Pairs(
// 			utils.GrpcGatewayHTTPPath,
// 			httpPath,
// 			utils.GrpcGaewayMethod,
// 			httpMethod,
// 			utils.IPAddress,
// 			ipAddress,
// 			utils.UserAgent,
// 			userAgent,
// 			utils.Platform,
// 			platform,
// 		),
// 	)

// 	if method, ok := m.publicEndpoints[httpPath]; ok && slices.Contains(method, httpMethod) {
// 		return handler(newCtx, req)
// 	}

// 	if apiKey != "" {
// 		payload, errToken := token.ParseToken(apiKey)
// 		if errToken != nil {
// 			zap.L().Error("failed to parse api-key token", zap.Error(errToken))
// 			return nil, errToken
// 		}

// 		if time.Now().After(payload.ExpiresAt) {
// 			return nil, errors.New("your company api-key expired")
// 		}

// 		if payload.Type != token.Company {
// 			return nil, errors.New("wrong api-key")
// 		}

// 		// not checked company from DB.
// 		return handler(newCtx, req)
// 	}

// 	accessToken := utils.ParseKeyFromCtx(ctx, utils.AuthorizationHeader)
// 	fields := strings.Fields(accessToken)
// 	if len(fields) < 2 || strings.ToLower(fields[0]) != authorizationForm {
// 		return nil, status.Errorf(codes.Unauthenticated, "invalid authorization header")
// 	}

// 	newCtx = metadata.NewOutgoingContext(newCtx,
// 		metadata.Pairs(
// 			utils.AuthorizationHeader, accessToken,
// 			utils.GrpcGatewayHTTPPath, httpPath,
// 			utils.GrpcGaewayMethod, httpMethod,
// 		),
// 	)

// 	accessToken = fields[1]

// 	tokenResp, err := m.authClient.ParseToken(newCtx,
// 		&auth.ParseTokenRequest{
// 			AccessToken: accessToken,
// 			HttpPath:    httpPath,
// 			HttpMethod:  httpMethod,
// 		},
// 	)
// 	if err != nil {
// 		zap.L().Error("failed to parse token", zap.Error(err))
// 		return nil, err
// 	}

// 	ctx = context.WithValue(ctx, constants.UserID, tokenResp.GetId())
// 	ctx = context.WithValue(ctx, constants.Role, tokenResp.GetRole())
// 	ctx = context.WithValue(ctx, constants.SessionID, tokenResp.GetSessionId())
// 	ctx = metadata.NewOutgoingContext(
// 		ctx,
// 		metadata.Pairs(
// 			utils.AuthorizationHeader,
// 			"Bearer "+accessToken,
// 			utils.GrpcGatewayHTTPPath,
// 			httpPath,
// 			utils.GrpcGaewayMethod,
// 			httpMethod,
// 		),
// 	)

// 	return handler(ctx, req)
// }
