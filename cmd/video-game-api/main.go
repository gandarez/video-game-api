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
	redis "github.com/gandarez/video-game-api/internal/cache"
	"github.com/gandarez/video-game-api/internal/client/igdb"
	"github.com/gandarez/video-game-api/internal/client/twitch"
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
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Load configuration from env vars.
	cfg, err := config.Load(ctx, "./.env")
	if err != nil {
		logger.Error("failed to load configuration", slog.Any("error", err))

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
		logger.Error("failed to open database connection", slog.Any("error", err))

		os.Exit(1)
	}

	// Run database migrations
	if err = dbmigration.Run(db.ConnectionString, logger); err != nil {
		logger.Error(err.Error())

		os.Exit(1)
	}

	// create cache client
	cacheClient := redis.NewClient(redis.Configuration{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// create Twitch client
	twitchClient := twitch.NewClient(twitch.Config{
		BaseURL:      cfg.VendorTwitch.Host,
		Cache:        cacheClient,
		CacheKey:     "twitch:token",
		ClientID:     cfg.VendorTwitch.ClientID,
		ClientSecret: cfg.VendorTwitch.ClientSecret,
		Logger:       logger,
	})

	// create IGDB client
	igdbClient := igdb.NewClient(igdb.Config{
		BaseURL:      cfg.VendorIGDB.Host,
		Logger:       logger,
		TwitchClient: twitchClient,
	})

	// setup server
	httpserver := server.New(cfg.Server.Port,
		server.WithRecover(logger),
		server.WithDecompress(),
		server.WithGzip(),
	)

	// add http routes
	httpserver.AddRoute(handlerhttp.SearchConsoleByID(ctx, logger, db))
	httpserver.AddRoute(handlerhttp.CreateConsole(ctx, logger, db))
	httpserver.AddRoute(handlerhttp.SearchGameByName(ctx, logger, igdbClient))

	// Add default routes for health check
	httpserver.AddRoute(server.ReadinessRoute())
	httpserver.AddRoute(server.LivenessRoute())

	// start httpserver
	go func() {
		if err := httpserver.Start(); err != http.ErrServerClosed {
			logger.Error("failed to start server", slog.Any("error", err))
		}
	}()

	// setup gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		logger.Error("failed to start listening: %s", err)

		os.Exit(1)
	}

	opts := []grpc.ServerOption{
		middleware.WithPanicRecovery(logger),
		middleware.WithUnknownServiceHandler(logger),
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
			logger.Error("failed to serve gRPC", slog.Any("error", err))
		}
	}()

	logger.Info("gRPC server listening", slog.String("addr", lis.Addr().String()))
	logger.Info("http server started", slog.Int("port", cfg.Server.Port))

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

	logger.Info("server gracefully stopped")
}
