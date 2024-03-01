package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	dbmigration "github.com/gandarez/video-game-api/db"
	"github.com/gandarez/video-game-api/internal/config"
	"github.com/gandarez/video-game-api/internal/database"
	"github.com/gandarez/video-game-api/internal/grpc/middleware"
	handlergrpc "github.com/gandarez/video-game-api/internal/handler/grpc"
	handlerhttp "github.com/gandarez/video-game-api/internal/handler/http"
	"github.com/gandarez/video-game-api/internal/server"
	pbconsole "github.com/gandarez/video-game-api/pkg/grpc/console"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	// Setup logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Load configuration from env vars.
	cfg, err := config.Load(ctx, "./.env")
	if err != nil {
		logger.ErrorContext(ctx, "failed to load configuration", slog.Any("error", err))

		os.Exit(1)
	}

	// Setup database
	db := database.NewClient(database.Configuration{
		DbName:   cfg.Database.Name,
		Host:     cfg.Database.Host,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Port:     cfg.Database.Port,
	})

	// Open database connection
	if err = db.Open(ctx); err != nil {
		logger.ErrorContext(ctx, "failed to open database connection", slog.Any("error", err))

		os.Exit(1)
	}

	// Run database migrations
	if err = dbmigration.Run(db.ConnectionString, logger); err != nil {
		logger.ErrorContext(ctx, err.Error())

		os.Exit(1)
	}

	// setup server
	httpserver := server.New(cfg.Server.Port,
		server.WithRecover(ctx, logger),
		server.WithDecompress(),
		server.WithGzip(),
	)

	// add http routes
	httpserver.AddRoute(handlerhttp.SearchConsoleByID(ctx, logger, db))
	httpserver.AddRoute(handlerhttp.CreateConsole(ctx, logger, db))

	// start httpserver
	go func() {
		if err := httpserver.Start(); err != http.ErrServerClosed {
			logger.Error("failed to start server", slog.Any("error", err))
		}
	}()

	// setup gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		logger.ErrorContext(ctx, "failed to start listening: %s", err)

		os.Exit(1)
	}

	opts := []grpc.ServerOption{
		middleware.WithPanicRecovery(ctx, logger),
		middleware.WithUnknownServiceHandler(ctx, logger),
	}

	// Create a gRPC server object
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	// Setup gRPC handlers
	handlerConsole := handlergrpc.NewConsoleServer(db, logger)

	// Register gRPC services
	pbconsole.RegisterConsoleServiceServer(grpcServer, handlerConsole)

	// start gRPC server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.ErrorContext(ctx, "failed to serve gRPC", slog.Any("error", err))
		}
	}()

	logger.InfoContext(ctx, "gRPC server listening", slog.String("addr", lis.Addr().String()))
	logger.InfoContext(ctx, "http server started", slog.Int("port", cfg.Server.Port))

	// wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// shutdown http server
	if err := httpserver.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown server", slog.Any("error", err))
	}

	// shutdown gRPC server
	grpcServer.GracefulStop()

	// purge database
	if err := dbmigration.Purge(db.ConnectionString); err != nil {
		logger.Warn("failed to purge database", slog.Any("error", err))
	}

	logger.Info("server gracefully stopped")
}
