package repository

import (
	"context"
	"fmt"

	"github.com/gandarez/video-game-api/internal/database"
	"github.com/gandarez/video-game-api/internal/entity"
)

//go:generate mockery --name ConsoleRepository --structname MockConsoleRepository --inpackage --case snake
type (
	// ConsoleRepository represents the repository for console.
	ConsoleRepository interface {
		FindByID(ctx context.Context, id string) (*entity.Console, error)
		Save(ctx context.Context, console *entity.Console) error
	}

	// Console is the repository for console.
	Console struct {
		db database.QueryExecutor
		tx database.Transactioner
	}
)

var _ ConsoleRepository = (*Console)(nil)

// NewConsole creates a new console repository.
func NewConsole(db database.QueryExecutor, tx database.Transactioner) ConsoleRepository {
	return &Console{
		db: db,
		tx: tx,
	}
}

// FindByID returns a console by id.
func (c *Console) FindByID(ctx context.Context, id string) (*entity.Console, error) {
	var console entity.Console

	err := c.db.QueryRow(
		ctx,
		"SELECT id, name, manufacturer, release_date FROM console WHERE id = $1", id).
		Scan(&console.ID, &console.Name, &console.Manufacturer, &console.ReleaseDate)
	if err != nil {
		return nil, handleFindError(err, "console", id)
	}

	return &console, nil
}

// Save saves a console.
func (c *Console) Save(ctx context.Context, console *entity.Console) error {
	affected, err := c.db.Exec(
		ctx,
		"INSERT INTO console (id, name, manufacturer, release_date) VALUES ($1, $2, $3, $4)",
		console.ID,
		console.Name,
		console.Manufacturer,
		console.ReleaseDate,
	)
	if err != nil {
		return handleSaveError(err, "console")
	}

	if affected != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", affected)
	}

	return nil
}
