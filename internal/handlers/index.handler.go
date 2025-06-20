package handlers

import (
	u "gabtec/go-echo-obs-app/internal/utils"
	"gabtec/go-echo-obs-app/version"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {

	endpoints := map[string]string{
		"/":          "this page",
		"/log/:type": "returns a success (type=ok), or error (type=error), response",
		"/random":    "returns a random response",
		"/demo":      "same as /random (for back compatibility)",
	}

	resp := map[string]interface{}{
		"endpointsList": endpoints,
		"version":       version.Version(),
		"observability": "not implemented",
	}

	return u.JSONOK(c, resp)
}
