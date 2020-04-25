package models

type ErrorResponse struct {
	Status  string `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message"`
}