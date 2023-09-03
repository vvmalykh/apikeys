package model

type InvalidationResult struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
