package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Store interface {
	Querier
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store.
func NewStore(ctx context.Context, psqlURI string) Store {
	zap.L().Info("connecting to psql...")

	dbConn, err := pgxpool.New(ctx, psqlURI)
	if err != nil {
		zap.L().Fatal("failed to connecto to psql", zap.Error(err))
	}

	zap.L().Info("psql connected")

	queries := New(dbConn)

	return &SQLStore{
		connPool: dbConn,
		Queries:  queries,
	}
}
