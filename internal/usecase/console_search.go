package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gandarez/video-game-api/internal/entity"
	"github.com/gandarez/video-game-api/internal/model"
	"github.com/gandarez/video-game-api/internal/repository"

	"github.com/google/uuid"
)

// ConsoleSearch is a use case for searching a console.
type ConsoleSearch struct {
	repo repository.ConsoleRepository
}

// NewConsoleSearch creates a new console search use case.
func NewConsoleSearch(repo repository.ConsoleRepository) *ConsoleSearch {
	return &ConsoleSearch{
		repo: repo,
	}
}

// Search searches a console.
func (c *ConsoleSearch) Search(ctx context.Context, id string) (*model.Console, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %s", id)
	}

	console, err := c.repo.FindByID(ctx, uid.String())
	if err != nil {
		return nil, err
	}

	return parseConsoleEntity(console), nil
}

func parseConsoleEntity(input *entity.Console) *model.Console {
	return &model.Console{
		ID:           input.ID.String(),
		Name:         input.Name,
		Manufacturer: input.Manufacturer,
		ReleaseDate:  input.ReleaseDate.Format(time.DateOnly),
	}
}
