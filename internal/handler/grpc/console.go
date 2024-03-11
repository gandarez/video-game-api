package grpc

import (
	"log/slog"

	"github.com/gandarez/video-game-api/internal/repository"
)

// ConsoleServer is the gRPC server for console service.
type ConsoleServer struct {
	db     repository.DatabaseQueryExecutor
	logger *slog.Logger
}

// NewConsoleServer creates a new console server.
func NewConsoleServer(db repository.DatabaseQueryExecutor, logger *slog.Logger) *ConsoleServer {
	return &ConsoleServer{
		db:     db,
		logger: logger,
	}
}
