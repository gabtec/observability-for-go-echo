package services

import (
	"gabtec/go-echo-obs-app/internal/model"
	"gabtec/go-echo-obs-app/internal/stores"
)

func GenerateSuccessLog() model.ApiResponse {
	return stores.GetRandomSuccessMessage()
}

func GenerateErrorLog() model.ApiResponse {
	return stores.GetRandomErrorMessage()
}
