package model

type ValidationResult struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
