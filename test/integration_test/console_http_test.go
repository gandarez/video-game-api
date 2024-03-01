//go:build integration

package integration_test

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/gandarez/video-game-api/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConsoleSearch_Http(t *testing.T) {
	apiURL := os.Getenv("VIDEO_GAME_API_URL")
	url := apiURL + "/consoles/b171ae30-2d02-4da2-98b4-33ad2c331669" // Xbox 360

	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	resbody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var console model.Console

	err = json.Unmarshal(resbody, &console)
	require.NoError(t, err)

	assert.Equal(t, "b171ae30-2d02-4da2-98b4-33ad2c331669", console.ID)
	assert.Equal(t, "Xbox 360", console.Name)
	assert.Equal(t, "Microsoft", console.Manufacturer)
	assert.Equal(t, "2005-11-22", console.ReleaseDate)
}

func TestConsoleSearch_Http_NotFound(t *testing.T) {
	apiURL := os.Getenv("VIDEO_GAME_API_URL")
	url := apiURL + "/consoles/d5a0469e-5825-4838-b5c7-2ccd3bca1344" // Not Found

	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()

	require.Equal(t, http.StatusNotFound, resp.StatusCode)

	resbody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var errors model.ErrorResponse

	err = json.Unmarshal(resbody, &errors)
	require.NoError(t, err)

	assert.Len(t, errors.Errors, 1)

	assert.Equal(t, "console with id d5a0469e-5825-4838-b5c7-2ccd3bca1344 not found", errors.Errors[0])
}

func TestConsoleCreate_Http(t *testing.T) {
	apiURL := os.Getenv("VIDEO_GAME_API_URL")
	url := apiURL + "/consoles"

	data, err := os.ReadFile("testdata/api_consoles_create_request.json")
	require.NoError(t, err)

	reqbody := strings.NewReader(string(data))

	req, err := http.NewRequest(http.MethodPost, url, reqbody)
	require.NoError(t, err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()

	require.Equal(t, http.StatusCreated, resp.StatusCode)

	resbody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var console model.Console

	err = json.Unmarshal(resbody, &console)
	require.NoError(t, err)

	assert.Equal(t, "Super Nintendo", console.Name)
	assert.Equal(t, "Nintendo", console.Manufacturer)
	assert.Equal(t, "1990-11-21", console.ReleaseDate)

	_, err = uuid.Parse(console.ID)
	assert.NoError(t, err)
}

func TestConsoleCreate_Http_AlreadyExists(t *testing.T) {
	apiURL := os.Getenv("VIDEO_GAME_API_URL")
	url := apiURL + "/consoles"

	data, err := os.ReadFile("testdata/api_consoles_create_duplicated_request.json")
	require.NoError(t, err)

	reqbody := strings.NewReader(string(data))

	req, err := http.NewRequest(http.MethodPost, url, reqbody)
	require.NoError(t, err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()

	require.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	resbody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var errors model.ErrorResponse

	err = json.Unmarshal(resbody, &errors)
	require.NoError(t, err)

	assert.Len(t, errors.Errors, 1)

	assert.Equal(t, "console already exists", errors.Errors[0])
}
