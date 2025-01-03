package testcontainer

import (
	"context"
	"fmt"
	"net"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	testcontainers.Container
	MappedPort   string
	Host         string
	databaseName string
}

func (c PostgresContainer) GetDSN() string {
	return fmt.Sprintf(
		"postgres://postgres:postgres@%s/%s?sslmode=disable",
		net.JoinHostPort(c.Host, c.MappedPort),
		c.databaseName,
	)
}

func NewPostgresContainer(ctx context.Context, databaseName string) (*PostgresContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15.5-alpine3.18",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
			"POSTGRES_DB":       databaseName,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		return nil, err
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		Container:    container,
		MappedPort:   mappedPort.Port(),
		Host:         host,
		databaseName: databaseName,
	}, nil
}
