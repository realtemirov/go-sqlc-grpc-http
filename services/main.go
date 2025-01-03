package services

import (
	"context"

	db "github.com/realtemirov/go-sqlc-grpc-http/db/sqlc"
	my_service "github.com/realtemirov/go-sqlc-grpc-http/generated/service"
)

type services struct {
	db db.Store
	my_service.UnimplementedMyServiceServer
}

func New(store db.Store) my_service.MyServiceServer {
	return &services{
		db: store,
	}
}

func (s *services) ServiceMethod(ctx context.Context, req *my_service.Request) (*my_service.Response, error) {
	return &my_service.Response{
		Message: req.GetMessage(),
	}, nil
}
