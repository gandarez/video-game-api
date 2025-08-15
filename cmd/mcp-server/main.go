package main

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gandarez/video-game-api/internal/model"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type ConsoleSearch struct {
	ID string `json:"id" jsonschema:"required, description=The ID of the console to search"`
}

type GameSearch struct {
	Name string `json:"name" jsonschema:"required, description=The name of the game to search"`
}

func main() {
	ctx := context.Background()

	// Setup logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	if err := server.RegisterTool("console_create", "Create a console",
		func(arg model.ConsoleInsert) (*mcp_golang.ToolResponse, error) {
			url := "http://localhost:17020/consoles"

			data, err := json.Marshal(arg)
			if err != nil {
				logger.Error("failed to marshal console data", slog.Any("error", err))

				return nil, err
			}

			reqbody := strings.NewReader(string(data))

			req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reqbody)
			if err != nil {
				logger.Error("failed to create console request", slog.Any("error", err))

				return nil, err
			}

			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")

			client := http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				logger.Error("failed to send console request", slog.Any("error", err))

				return nil, err
			}

			defer resp.Body.Close() // nolint: errcheck

			resbody, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Error("failed to read console response body", slog.Any("error", err))

				return nil, err
			}

			// var console model.Console
			// if err := json.Unmarshal(resbody, &console); err != nil {
			// 	logger.Error("failed to unmarshal console response", slog.Any("error", err))

			// 	return nil, err
			// }

			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(resbody))), nil
		}); err != nil {
		logger.Error("failed to register console_create tool", slog.Any("error", err))

		os.Exit(1)
	}

	if err := server.RegisterTool("console_search", "Find a console by id",
		func(arg ConsoleSearch) (*mcp_golang.ToolResponse, error) {
			url := "http://localhost:17020/consoles/" + arg.ID

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				logger.Error("failed to create console request", slog.Any("error", err))

				return nil, err
			}

			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")

			client := http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				logger.Error("failed to send console request", slog.Any("error", err))

				return nil, err
			}

			defer resp.Body.Close() // nolint: errcheck

			resbody, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Error("failed to read console response body", slog.Any("error", err))

				return nil, err
			}

			// var console model.Console
			// if err := json.Unmarshal(resbody, &console); err != nil {
			// 	logger.Error("failed to unmarshal console response", slog.Any("error", err))

			// 	return nil, err
			// }

			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(resbody))), nil
		}); err != nil {
		logger.Error("failed to register console_search tool", slog.Any("error", err))

		os.Exit(1)
	}

	if err := server.RegisterTool("game_search", "Find a game by name",
		func(arg GameSearch) (*mcp_golang.ToolResponse, error) {
			url := "http://localhost:17020/games/" + arg.Name

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				logger.Error("failed to create console request", slog.Any("error", err))

				return nil, err
			}

			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")

			client := http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				logger.Error("failed to send console request", slog.Any("error", err))

				return nil, err
			}

			defer resp.Body.Close() // nolint: errcheck

			resbody, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Error("failed to read console response body", slog.Any("error", err))

				return nil, err
			}

			// var console model.Game
			// if err := json.Unmarshal(resbody, &console); err != nil {
			// 	logger.Error("failed to unmarshal console response", slog.Any("error", err))

			// 	return nil, err
			// }

			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(resbody))), nil
		}); err != nil {
		logger.Error("failed to register console_search tool", slog.Any("error", err))

		os.Exit(1)
	}

	if err := server.Serve(); err != nil {
		logger.Error("failed to start server", slog.Any("error", err))

		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
}
