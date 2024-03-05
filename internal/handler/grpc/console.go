package grpc

import (
	"log/slog"

	"github.com/gandarez/video-game-api/internal/database"
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
