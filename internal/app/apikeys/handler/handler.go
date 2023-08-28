package handler

import (
	"apikeys/internal/app/apikeys/service"
	"database/sql"
	"encoding/json"
	"net/http"
)

type APIKeyHandler struct {
	DB *sql.DB
}

// GenerateAPIKeyHandler handles the request for generating new API keys
func (h *APIKeyHandler) GenerateAPIKeyHandler(w http.ResponseWriter, r *http.Request) {
	service := &service.APIKeyService{DB: h.DB}

	// Generate new API key
	apiKeyString, err := service.GenerateAPIKey()
	if err != nil {
		http.Error(w, "Failed to serialize API key", http.StatusInternalServerError)
		return
	}

	// Prepare JSON response
	response, err := json.Marshal(apiKeyString)
	if err != nil {
		http.Error(w, "Failed to serialize API key", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
