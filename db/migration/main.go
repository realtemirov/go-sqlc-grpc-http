package migration

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // sql library for postgres
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations
var Migrations embed.FS

func Migrate(dsn string, path fs.FS) error {
	d, err := iofs.New(path, "migrations")
	if err != nil {
		return fmt.Errorf("failed to create iofs: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	err = m.Up()
	if err != nil {
		return fmt.Errorf("failed to migrate up: %w", err)
	}

	return nil
}
