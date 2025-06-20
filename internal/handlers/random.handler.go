package handlers

import (
	"gabtec/go-echo-obs-app/internal/model"
	"gabtec/go-echo-obs-app/internal/services"
	"math/rand"

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

	return c.JSON(resp.StatusCode, resp)
}
