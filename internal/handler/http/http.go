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

// SearchConsoleByID returns a console by id.
func SearchConsoleByID(ctx context.Context, logger *slog.Logger, db database.Connector) server.Route {
	return server.Route{
		Method: "GET",
		Path:   "/consoles/:id",
		Handler: func(c echo.Context) error {
			logger.InfoContext(ctx, "search console by id")

			id := c.Param("id")

			uc := usecase.NewConsoleSearch(
				repository.NewConsole(db, nil),
			)

			console, err := uc.Search(ctx, id)
			if err != nil {
				return c.JSON(errorHandler(ctx, logger, err))
			}

			logger.InfoContext(ctx, "console found", slog.String("id", console.ID))

			return c.JSON(http.StatusOK, console)
		},
	}
}

// CreateConsole creates a new console.
func CreateConsole(ctx context.Context, logger *slog.Logger, db database.Connector) server.Route {
	return server.Route{
		Method: "POST",
		Path:   "/consoles",
		Handler: func(c echo.Context) error {
			logger.InfoContext(ctx, "create console")

			var body model.ConsoleInsert

			if err := c.Bind(&body); err != nil {
				return c.JSON(errorHandler(ctx, logger, err))
			}

			uc := usecase.NewConsoleCreate(
				repository.NewConsole(db, nil),
			)

			console, err := uc.Create(ctx, body)
			if err != nil {
				return c.JSON(errorHandler(ctx, logger, err))
			}

			logger.InfoContext(ctx, "successfully created console", slog.String("id", console.ID))

			return c.JSON(http.StatusCreated, console)
		},
	}
}
