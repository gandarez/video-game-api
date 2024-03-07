package integration_test

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/gandarez/video-game-api/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGameSearch_Http(t *testing.T) {
	apiURL := os.Getenv("VIDEO_GAME_API_URL")
	url := apiURL + "/games/Mario"

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

	var games []model.Game

	err = json.Unmarshal(resbody, &games)
	require.NoError(t, err)

	assert.Len(t, games, 10)

	// check response body
	expectedResponseBody, err := os.ReadFile("testdata/igdb_games_search_response.json")
	require.NoError(t, err)

	assert.JSONEq(t, string(expectedResponseBody), string(resbody))
}
