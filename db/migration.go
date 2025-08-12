package db

import (
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // nolint:revive
	_ "github.com/golang-migrate/migrate/v4/source/file"       // nolint:revive
)

// Run runs the database migrations.
func Run(migrations string, connString string, logger *slog.Logger) error {
	logger.Debug("starting database migrations")

	m, err := migrate.New(
		"file://"+migrations,
		connString,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize database for migration: %s", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Debug("no migrations to run")

			return nil
		}

		return fmt.Errorf("failed to run migrations: %s", err)
	}

	logger.Debug("database migrations completed")

	return nil
}
