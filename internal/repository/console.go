package repository

import (
	"context"
	"fmt"

	"github.com/gandarez/video-game-api/internal/entity"

	"github.com/jackc/pgx/v5"
)

type (
	// DatabaseQueryExecutor is the interface for executing database queries.
	DatabaseQueryExecutor interface {
		Exec(ctx context.Context, sql string, args ...any) (int64, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	}

	// Console is the repository for console.
	Console struct {
		db DatabaseQueryExecutor
	}
)

// NewConsole creates a new console repository.
func NewConsole(db DatabaseQueryExecutor) *Console {
	return &Console{
		db: db,
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

func (c *Console) FindAll(ctx context.Context, page, rows int) ([]*entity.Console, error) {
	if page < 1 {
		page = 1
	}

	if rows < 1 || rows > 100 {
		rows = 100
	}

	offset := (page - 1) * rows

	query := "SELECT id, name, manufacturer, release_date FROM console ORDER BY name LIMIT $1 OFFSET $2"

	rowsResult, err := c.db.Query(ctx, query, rows, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rowsResult.Close()

	var consoles []*entity.Console
	for rowsResult.Next() {
		var console entity.Console

		if err := rowsResult.Scan(&console.ID, &console.Name, &console.Manufacturer, &console.ReleaseDate); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		consoles = append(consoles, &console)
	}

	if err := rowsResult.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return consoles, nil
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
