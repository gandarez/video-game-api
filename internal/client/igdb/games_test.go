package igdb_test

import (
	// "context"
	// "io"
	// "net/http"
	// "os"
	"testing"
	// "time"
	// "github.com/gandarez/video-game-api/internal/api"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

func TestClient_Games(t *testing.T) {
	// testServerURL, router, testTwitchURL, routerTwitch, tearDown := setupTestServer()
	// defer tearDown()

	// var numCalls int

	// router.HandleFunc(
	// 	"/Games/Mario", func(w http.ResponseWriter, req *http.Request) {
	// 		numCalls++

	// 		// check request
	// 		assert.Equal(t, http.MethodPost, req.Method)
	// 		assert.Equal(t, []string{"application/json"}, req.Header["Accept"])
	// 		assert.Equal(t, []string{"application/json"}, req.Header["Content-Type"])

	// 		// write response
	// 		f, err := os.Open("testdata/api_goals_id_response.json")
	// 		require.NoError(t, err)

	// 		w.WriteHeader(http.StatusOK)
	// 		_, err = io.Copy(w, f)
	// 		require.NoError(t, err)
	// 	})

	// c := api.NewClient(api.Config{BaseURL: testServerURL})
	// goal, err := c.Games(context.Background(), "Mario")

	// require.NoError(t, err)

	// assert.Equal(t, "3 hrs 23 mins", goal.Data.ChartData[len(goal.Data.ChartData)-1].ActualSecondsText)

	// assert.Eventually(t, func() bool { return numCalls == 1 }, time.Second, 50*time.Millisecond)
}
