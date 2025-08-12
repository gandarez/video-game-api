//go:build integration

package integration_test

import (
	"context"
	"os"
	"testing"

	pbconsole "github.com/gandarez/video-game-api/pkg/grpc/console"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestConsoleSearch_Grpc(t *testing.T) {
	grpcHost := os.Getenv("VIDEO_GAME_GRPC_HOST")

	conn, err := grpc.NewClient(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	grpcClient := pbconsole.NewConsoleServiceClient(conn)

	res, err := grpcClient.GetConsole(context.Background(), &pbconsole.GetConsoleRequest{
		Id: "b171ae30-2d02-4da2-98b4-33ad2c331669",
	})
	require.NoError(t, err)

	require.NotNil(t, res)

	assert.Equal(t, "b171ae30-2d02-4da2-98b4-33ad2c331669", res.GetConsole().GetId())
	assert.Equal(t, "Xbox 360", res.GetConsole().GetName())
	assert.Equal(t, "Microsoft", res.GetConsole().GetManufacturer())
	assert.Equal(t, "2005-11-22", res.GetConsole().GetReleaseDate())
}

func TestConsoleSearch_Grpc_NotFound(t *testing.T) {
	grpcHost := os.Getenv("VIDEO_GAME_GRPC_HOST")

	conn, err := grpc.NewClient(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	grpcClient := pbconsole.NewConsoleServiceClient(conn)

	res, err := grpcClient.GetConsole(context.Background(), &pbconsole.GetConsoleRequest{
		Id: "d5a0469e-5825-4838-b5c7-2ccd3bca1344",
	})
	require.Error(t, err)

	require.Nil(t, res)

	status, ok := status.FromError(err)
	require.True(t, ok)

	assert.Equal(t, codes.NotFound, status.Code())
	assert.Equal(t, "console with id d5a0469e-5825-4838-b5c7-2ccd3bca1344 not found", status.Message())
}

func TestConsoleCreate_Grpc(t *testing.T) {
	grpcHost := os.Getenv("VIDEO_GAME_GRPC_HOST")

	conn, err := grpc.NewClient(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	grpcClient := pbconsole.NewConsoleServiceClient(conn)

	res, err := grpcClient.CreateConsole(context.Background(), &pbconsole.CreateConsoleRequest{
		Name:         "Atari 2600",
		Manufacturer: "Atari, Inc.",
		ReleaseDate:  "1977-09-11",
	})
	require.NoError(t, err)

	require.NotNil(t, res)

	assert.Equal(t, "Atari 2600", res.GetConsole().GetName())
	assert.Equal(t, "Atari, Inc.", res.GetConsole().GetManufacturer())
	assert.Equal(t, "1977-09-11", res.GetConsole().GetReleaseDate())

	_, err = uuid.Parse(res.GetConsole().GetId())
	assert.NoError(t, err)
}

func TestConsoleCreate_Grpc_AlreadyExists(t *testing.T) {
	grpcHost := os.Getenv("VIDEO_GAME_GRPC_HOST")

	conn, err := grpc.NewClient(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	grpcClient := pbconsole.NewConsoleServiceClient(conn)

	res, err := grpcClient.CreateConsole(context.Background(), &pbconsole.CreateConsoleRequest{
		Name:         "Xbox 360",
		Manufacturer: "Microsoft",
		ReleaseDate:  "2005-11-22",
	})
	require.Error(t, err)

	require.Nil(t, res)

	status, ok := status.FromError(err)
	require.True(t, ok)

	assert.Equal(t, codes.AlreadyExists, status.Code())
	assert.Equal(t, "console already exists", status.Message())
}
