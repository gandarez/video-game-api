package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gandarez/video-game-api/internal/database"
	"github.com/gandarez/video-game-api/internal/model"
	"github.com/gandarez/video-game-api/internal/repository"
	"github.com/gandarez/video-game-api/internal/server"
	"github.com/gandarez/video-game-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

// CreateConsole creates a new console.
func CreateConsole(ctx context.Context, logger *slog.Logger, db database.Connector) server.Route {
	return server.Route{
		Method: "POST",
		Path:   "/consoles",
		Handler: func(c echo.Context) error {
			logger.Info("create console")

			var body model.ConsoleInsert

			if err := c.Bind(&body); err != nil {
				return c.JSON(errorHandler(logger, err))
			}

			uc := usecase.NewConsoleCreate(
				repository.NewConsole(db, nil),
			)

			console, err := uc.Create(ctx, body)
			if err != nil {
				return c.JSON(errorHandler(logger, err))
			}

			logger.Info("successfully created console", slog.String("id", console.ID))

			return c.JSON(http.StatusCreated, console)
		},
	}
}
