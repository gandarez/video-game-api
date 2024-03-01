package db

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // nolint:blank-imports
	_ "github.com/golang-migrate/migrate/v4/source/file"       // nolint:blank-imports
)

// Run runs the database migrations.
func Run(connString string, logger *slog.Logger) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %s", err)
	}

	var path string

	dockerized := os.Getenv("DOCKER")
	if dockerized == "true" {
		path = filepath.Join(wd, "db/migrations")
	} else {
		path = filepath.Join(wd, "../../", "db/migrations")
	}

	m, err := migrate.New(
		"file://"+path,
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

	return nil
}

// Purge purges the database.
func Purge(connString string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %s", err)
	}

	var path string

	dockerized := os.Getenv("DOCKER")
	if dockerized == "true" {
		path = filepath.Join(wd, "db/migrations")
	} else {
		path = filepath.Join(wd, "../../", "db/migrations")
	}

	m, err := migrate.New(
		"file://"+path,
		connString,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize database for drop: %s", err)
	}

	if err := m.Drop(); err != nil {
		return fmt.Errorf("failed to drop database: %s", err)
	}

	return nil
}
