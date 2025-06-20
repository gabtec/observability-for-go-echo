package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	indent2Spaces = "  "
	indent4Spaces = "    "
	indentTab     = "\t"
)

// JSONOK - returns a c.JSON response, already with http.StatusOK, set
func JSONOK(c echo.Context, resp interface{}) error {
	return c.JSON(http.StatusOK, resp)
}

// JSONiOK - returns a c.JSONPretty response, already with http.StatusOK, set
func JSONiOK(c echo.Context, resp interface{}) error {
	return c.JSONPretty(http.StatusOK, resp, indent2Spaces)
}
