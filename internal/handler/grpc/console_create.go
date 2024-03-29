package grpc

import (
	"context"
	"log/slog"

	"github.com/gandarez/video-game-api/internal/model"
	"github.com/gandarez/video-game-api/internal/repository"
	"github.com/gandarez/video-game-api/internal/usecase"
	grpc "github.com/gandarez/video-game-api/pkg/grpc/console"
)

// CreateConsole creates a new console.
func (s *ConsoleServer) CreateConsole(
	ctx context.Context,
	in *grpc.CreateConsoleRequest) (*grpc.CreateConsoleResponse, error) {
	s.logger.Info("create console")

	uc := usecase.NewConsoleCreate(
		repository.NewConsole(s.db),
	)

	input := model.ConsoleInsert{
		Name:         in.GetName(),
		Manufacturer: in.GetManufacturer(),
		ReleaseDate:  in.GetReleaseDate(),
	}

	console, err := uc.Create(ctx, input)
	if err != nil {
		return handleError[grpc.CreateConsoleResponse](s.logger, err)
	}

	s.logger.Info("successfully created console", slog.String("id", console.ID))

	return &grpc.CreateConsoleResponse{
		Console: &grpc.Console{
			Id:           console.ID,
			Name:         console.Name,
			Manufacturer: console.Manufacturer,
			ReleaseDate:  console.ReleaseDate,
		},
	}, nil
}
