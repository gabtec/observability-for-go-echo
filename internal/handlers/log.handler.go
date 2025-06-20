package handlers

import (
	"gabtec/go-echo-obs-app/internal/model"
	"gabtec/go-echo-obs-app/internal/services"

	"github.com/labstack/echo/v4"
)

func LogHandler(c echo.Context) error {
	selectedLogType := c.Param("type")

	var resp model.ApiResponse

	if selectedLogType == "error" {
		resp = services.GenerateErrorLog()
	} else {
		resp = services.GenerateSuccessLog()
	}

	return c.JSON(resp.StatusCode, resp)
}
