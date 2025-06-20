package handlers

import (
	"gabtec/go-echo-obs-app/internal/model"
	"gabtec/go-echo-obs-app/internal/services"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RandomHandler(c echo.Context) error {
	selectedLogType := rand.Intn(2)

	var resp model.ApiResponse

	if selectedLogType == 0 {
		resp = services.GenerateErrorLog()
	} else {
		resp = services.GenerateSuccessLog()
	}

	if resp.StatusCode == http.StatusNoContent {
		return c.NoContent(resp.StatusCode)
	}

	return c.JSON(resp.StatusCode, resp)
}
