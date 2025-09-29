package server

import (
	"context"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/labstack/echo/v4"
)

type dbChecker interface {
	Check(ctx context.Context) error
}

// ReadinessRoute checks if the service and its dependencies are healthy.
func ReadinessRoute(db dbChecker) Route {
	return Route{
		Method: "GET",
		Path:   "/readiness",
		Handler: func(c echo.Context) error {
			if err := db.Check(c.Request().Context()); err != nil {
				return c.String(http.StatusServiceUnavailable, "Unavailable")
			}

			return c.String(http.StatusOK, "Ok")
		},
	}
}

// LivenessRoute checks if the service is alive.
func LivenessRoute() Route {
	return Route{
		Method: "GET",
		Path:   "/liveness",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{
				"status":     "ok",
				"name":       os.Getenv("KUBERNETES_NAME"),
				"pod_ip":     os.Getenv("KUBERNETES_POD_IP"),
				"node":       os.Getenv("KUBERNETES_NODE_NAME"),
				"namespace":  os.Getenv("KUBERNETES_NAMESPACE"),
				"GOMAXPROCS": strconv.Itoa(runtime.GOMAXPROCS(0)),
			})
		},
	}
}
