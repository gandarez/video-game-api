package grpc

import (
	"context"
	"log/slog"

	"github.com/gandarez/video-game-api/internal/repository"
	"github.com/gandarez/video-game-api/internal/usecase"
	grpc "github.com/gandarez/video-game-api/pkg/grpc/console"
)

// GetConsole returns a console by id.
func (s *ConsoleServer) GetConsole(
	ctx context.Context,
	in *grpc.GetConsoleRequest) (*grpc.GetConsoleResponse, error) {
	s.logger.Info("search console by id")

	uc := usecase.NewConsoleSearch(
		repository.NewConsole(s.db, nil),
	)

	console, err := uc.Search(ctx, in.GetId())
	if err != nil {
		return handleError[grpc.GetConsoleResponse](s.logger, err)
	}

	s.logger.Info("console found", slog.String("id", console.ID))

	return &grpc.GetConsoleResponse{
		Console: &grpc.Console{
			Id:           console.ID,
			Name:         console.Name,
			Manufacturer: console.Manufacturer,
			ReleaseDate:  console.ReleaseDate,
		},
	}, nil
}
