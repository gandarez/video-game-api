package grpc

import (
	"log/slog"

	"github.com/gandarez/video-game-api/internal/database"
	"github.com/gandarez/video-game-api/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// handleError handles error and returns a gRPC status error.
func handleError[T any](logger *slog.Logger, err error) (*T, error) {
	var code codes.Code

	logger.Error(err.Error())

	// assert error type.
	switch err.(type) {
	case database.ErrNotFound:
		code = codes.NotFound
	case database.ErrConflict:
		code = codes.AlreadyExists
	case model.ErrorResponse:
		code = codes.InvalidArgument
	default:
		code = codes.Internal
	}

	return nil, status.Error(code, err.Error())
}
