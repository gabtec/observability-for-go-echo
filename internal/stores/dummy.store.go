package stores

import (
	"gabtec/go-echo-obs-app/internal/model"
	"math/rand"
	"net/http"
)

var errorMessages = map[int]string{
	400: "Account not found",
	401: "You are not authorized to create users",
	402: "Payment is required",
	403: "You have no permissions to delete accounts",
	404: "Missing required field: 'X'",
	500: "Unexpected server behavior",
}

var successMessages = map[int]string{
	200: "Object fetched successfully",
	201: "Object created successfully",
	202: "Cronjob started successfully",
	203: "Cached object fetched successfully",
	204: "Object deleted successfully",
	206: "Object part fetched successfully",
}

func getRandomErrorCode(idx int) int {
	possibleCodes := []int{400, 401, 402, 403, 404, 500}
	return possibleCodes[idx]
}

func getRandomSuccessCode(idx int) int {
	possibleCodes := []int{200, 201, 202, 203, 204, 206}
	return possibleCodes[idx]
}

func GetRandomErrorMessage() model.ApiResponse {
	idx := rand.Intn(6)

	code := getRandomErrorCode(idx)

	return model.ApiResponse{
		Status:     "error",
		StatusCode: code,
		StatusText: http.StatusText(code),
		Message:    errorMessages[code],
		Emoji:      "⚠️",
	}
}

func GetRandomSuccessMessage() model.ApiResponse {
	idx := rand.Intn(6)

	code := getRandomSuccessCode(idx)

	return model.ApiResponse{
		Status:     "success",
		StatusCode: code,
		StatusText: http.StatusText(code),
		Message:    successMessages[code],
		Emoji:      "✅",
	}
}
