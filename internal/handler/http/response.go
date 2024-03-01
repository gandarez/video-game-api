package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gandarez/video-game-api/internal/database"
	"github.com/gandarez/video-game-api/internal/model"
)

// errorHandler handles error and returns a http status code and a response.
func errorHandler(ctx context.Context, logger *slog.Logger, err error) (int, any) {
	var httpcode int

	res := model.ErrorResponse{
		Errors: []string{err.Error()},
	}

	logger.ErrorContext(ctx, err.Error())

	// assert error type.
	switch errType := err.(type) {
	case ErrBind:
		httpcode = http.StatusBadRequest
	case database.ErrNotFound:
		httpcode = http.StatusNotFound
	case database.ErrConflict:
		httpcode = http.StatusUnprocessableEntity
	case model.ErrorResponse:
		httpcode = http.StatusBadRequest

		res = errType
	default:
		httpcode = http.StatusInternalServerError
	}

	return httpcode, res
}
