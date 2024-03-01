package grpc

import (
	"context"
	"log/slog"

	"github.com/gandarez/video-game-api/internal/database"
	"github.com/gandarez/video-game-api/internal/model"
	"github.com/gandarez/video-game-api/internal/repository"
	"github.com/gandarez/video-game-api/internal/usecase"
	grpc "github.com/gandarez/video-game-api/pkg/grpc/console"
)

// ConsoleServer is the gRPC server for console service.
type ConsoleServer struct {
	db     database.Connector
	logger *slog.Logger
}

// NewConsoleServer creates a new console server.
func NewConsoleServer(db database.Connector, logger *slog.Logger) *ConsoleServer {
	return &ConsoleServer{
		db:     db,
		logger: logger,
	}
}

// GetConsole returns a console by id.
func (s *ConsoleServer) GetConsole(
	ctx context.Context,
	in *grpc.GetConsoleRequest) (*grpc.GetConsoleResponse, error) {
	s.logger.InfoContext(ctx, "search console by id")

	uc := usecase.NewConsoleSearch(
		repository.NewConsole(s.db, nil),
	)

	console, err := uc.Search(ctx, in.GetId())
	if err != nil {
		return handleError[grpc.GetConsoleResponse](ctx, s.logger, err)
	}

	s.logger.InfoContext(ctx, "console found", slog.String("id", console.ID))

	return &grpc.GetConsoleResponse{
		Console: &grpc.Console{
			Id:           console.ID,
			Name:         console.Name,
			Manufacturer: console.Manufacturer,
			ReleaseDate:  console.ReleaseDate,
		},
	}, nil
}

// CreateConsole creates a new console.
func (s *ConsoleServer) CreateConsole(
	ctx context.Context,
	in *grpc.CreateConsoleRequest) (*grpc.CreateConsoleResponse, error) {
	s.logger.InfoContext(ctx, "create console")

	uc := usecase.NewConsoleCreate(
		repository.NewConsole(s.db, nil),
	)

	input := model.ConsoleInsert{
		Name:         in.GetName(),
		Manufacturer: in.GetManufacturer(),
		ReleaseDate:  in.GetReleaseDate(),
	}

	console, err := uc.Create(ctx, input)
	if err != nil {
		return handleError[grpc.CreateConsoleResponse](ctx, s.logger, err)
	}

	s.logger.InfoContext(ctx, "successfully created console", slog.String("id", console.ID))

	return &grpc.CreateConsoleResponse{
		Console: &grpc.Console{
			Id:           console.ID,
			Name:         console.Name,
			Manufacturer: console.Manufacturer,
			ReleaseDate:  console.ReleaseDate,
		},
	}, nil
}
