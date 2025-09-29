package http

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gandarez/video-game-api/internal/repository"
	"github.com/gandarez/video-game-api/internal/server"
	"github.com/gandarez/video-game-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

// ListConsoles returns a list of consoles.
func ListConsoles(ctx context.Context, logger *slog.Logger, db repository.DatabaseQueryExecutor) server.Route {
	return server.Route{
		Method: "GET",
		Path:   "/consoles",
		Handler: func(c echo.Context) error {
			logger.Info("list consoles")

			pageStr := c.QueryParam("page")
			rowsStr := c.QueryParam("rows")

			page, rows := parsePaginationParams(pageStr, rowsStr)

			uc := usecase.NewConsoleList(
				repository.NewConsole(db),
			)

			consoles, err := uc.List(ctx, page, rows)
			if err != nil {
				return c.JSON(errorHandler(logger, err))
			}

			logger.Info("consoles listed", slog.Int("page", page), slog.Int("rows", rows), slog.Int("count", len(consoles)))

			return c.JSON(http.StatusOK, consoles)
		},
	}
}

func parsePaginationParams(pageStr, rowsStr string) (int, int) {
	page := 1
	rows := 100

	if p, err := parsePositiveInt(pageStr); err == nil && p > 0 {
		page = p
	}

	if r, err := parsePositiveInt(rowsStr); err == nil && r > 0 && r <= 100 {
		rows = r
	}

	return page, rows
}

func parsePositiveInt(s string) (int, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if val < 0 {
		return 1, nil
	}

	return val, nil
}
