package model

type ApiResponse struct {
	Emoji      string `json:"emoji"`
	Status     string `json:"status"` // TODO enum: success or error
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
}
