package usecase

import (
	"context"
	"fmt"

	"github.com/gandarez/video-game-api/internal/entity"
	"github.com/gandarez/video-game-api/internal/model"
)

type (
	// ConsoleRepositoryLister is an interface for listing consoles from the repository.
	ConsoleRepositoryLister interface {
		FindAll(ctx context.Context, page, rows int) ([]*entity.Console, error)
	}

	// ConsoleList is a use case for listing consoles.
	ConsoleList struct {
		repo ConsoleRepositoryLister
	}
)

// NewConsoleList creates a new console list use case.
func NewConsoleList(repo ConsoleRepositoryLister) *ConsoleList {
	return &ConsoleList{
		repo: repo,
	}
}

// List lists consoles.
func (c *ConsoleList) List(ctx context.Context, page, rows int) ([]model.Console, error) {
	consoles, err := c.repo.FindAll(ctx, page, rows)
	if err != nil {
		return nil, fmt.Errorf("failed to list consoles: %w", err)
	}

	return parseConsoleEntities(consoles), nil
}

func parseConsoleEntities(input []*entity.Console) []model.Console {
	output := make([]model.Console, 0, len(input))

	for _, console := range input {
		output = append(output, *parseConsoleEntity(console))
	}

	return output
}
