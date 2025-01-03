package handlers

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/realtemirov/go-sqlc-grpc-http/config"
	my_service "github.com/realtemirov/go-sqlc-grpc-http/generated/service"
	"github.com/realtemirov/go-sqlc-grpc-http/pkg/libraries/wrapper"
	"github.com/realtemirov/go-sqlc-grpc-http/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

func New(ctx context.Context, cfg *config.Config) *runtime.ServeMux {
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	gwMux := runtime.NewServeMux(
		jsonOption,
		runtime.WithMetadata(func(_ context.Context, req *http.Request) metadata.MD {
			return metadata.New(map[string]string{
				utils.GrpcGatewayHTTPPath: req.URL.Path,
				utils.GrpcGaewayMethod:    req.Method,
				utils.APIKey:              req.Header.Get(utils.APIKey),
			})
		}),
		runtime.WithIncomingHeaderMatcher(wrapper.CustomMatcher),
	)

	grpcServerEndpoint := net.JoinHostPort(cfg.Grpc.Host, cfg.Grpc.Port)
	conn, err := grpc.NewClient(grpcServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}

	if err = my_service.RegisterMyServiceHandler(ctx, gwMux, conn); err != nil {
		log.Fatalf("failed to register my_service handler: %v", err)
	}

	return gwMux
}
