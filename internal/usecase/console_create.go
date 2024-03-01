package usecase

import (
	"context"
	"time"

	"github.com/gandarez/video-game-api/internal/entity"
	"github.com/gandarez/video-game-api/internal/model"
	"github.com/gandarez/video-game-api/internal/repository"

	"github.com/google/uuid"
)

// ConsoleCreate is a use case for creating a console.
type ConsoleCreate struct {
	repo repository.ConsoleRepository
}

// NewConsoleCreate creates a new console create use case.
func NewConsoleCreate(repo repository.ConsoleRepository) *ConsoleCreate {
	return &ConsoleCreate{
		repo: repo,
	}
}

// Create creates a console.
func (c *ConsoleCreate) Create(ctx context.Context, input model.ConsoleInsert) (*model.Console, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	console := parseConsoleInsertModel(input)

	if err := c.repo.Save(ctx, console); err != nil {
		return nil, err
	}

	return parseConsoleEntity(console), nil
}

func parseConsoleInsertModel(input model.ConsoleInsert) *entity.Console {
	var releaseDate time.Time

	if input.ReleaseDate != "" {
		releaseDate, _ = time.Parse(time.DateOnly, input.ReleaseDate)
	}

	return &entity.Console{
		ID:           uuid.New(),
		Name:         input.Name,
		Manufacturer: input.Manufacturer,
		ReleaseDate:  releaseDate,
	}
}
