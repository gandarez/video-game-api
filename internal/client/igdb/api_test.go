package igdb_test

import (
	"net/http"
	"net/http/httptest"
)

func setupTestServer() (string, *http.ServeMux, string, *http.ServeMux, func()) {
	router := http.NewServeMux()
	srv := httptest.NewServer(router)

	routerTwitch := http.NewServeMux()
	srvTwitch := httptest.NewServer(routerTwitch)

	return srv.URL, router, srvTwitch.URL, routerTwitch, func() {
		srv.Close()
		srvTwitch.Close()
	}
}
