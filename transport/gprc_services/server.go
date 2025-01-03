package gprcservices

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/realtemirov/go-sqlc-grpc-http/config"
	db "github.com/realtemirov/go-sqlc-grpc-http/db/sqlc"
	my_service "github.com/realtemirov/go-sqlc-grpc-http/generated/service"
	"github.com/realtemirov/go-sqlc-grpc-http/services"
	"github.com/realtemirov/go-sqlc-grpc-http/transport/gprc_services/clients"
	"github.com/realtemirov/go-sqlc-grpc-http/transport/gprc_services/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServerParams struct {
	Cfg     *config.Config
	DB      db.Store
	Clients *clients.GrpcClient
}

func New(params GrpcServerParams) *grpc.Server {
	// authMiddleware := middleware.NewAuthMiddleware(params.Clients.AuthServiceClient)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// authMiddleware.Auth,
				middleware.GrpcLoggerMiddleware,
				middleware.GrpcErrorMiddleware,
			),
		),
	)

	reflection.Register(grpcServer)

	my_service.RegisterMyServiceServer(grpcServer, services.New(params.DB))

	return grpcServer
}
