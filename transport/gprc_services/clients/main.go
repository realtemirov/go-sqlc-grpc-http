package clients

import (
	"github.com/realtemirov/go-sqlc-grpc-http/config"
)

type GrpcClient struct {
	cfg *config.Config
}

func NewGrpcClients(cfg *config.Config) *GrpcClient {
	// authConn, err := grpc.NewClient(
	// 	net.JoinHostPort(cfg.GrpcClient.AuthServiceHost, strconv.Itoa(cfg.GrpcClient.AuthServicePort)),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	zap.L().Fatal("failed to connect to auth service", zap.Error(err))
	// }

	return &GrpcClient{
		cfg: cfg,
		// AuthServiceClient:           auth.NewAuthServiceClient(authConn),
		// UserServiceClient:           users.NewUserServiceClient(authConn),
		// RolesAndPermissionsClient:   users.NewRolesAndPermissionsClient(authConn),
		// PermissionsManagementClient: users.NewPermissionsManagementClient(authConn),
	}
}
