package handler

import (
	m "apikeys/internal/app/apikeys/model"
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

// ValidateApiKeyHandler handles the request for validating ApiKey
func (h *APIKeyHandler) ValidateAPIKeyHandler(w http.ResponseWriter, r *http.Request) {
	service := &service.APIKeyService{DB: h.DB}
	apiKey := r.Header.Get("Authorization")

	// Validate the API key
	isValid, err := service.ValidateAPIKey(apiKey)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	validationResult := m.ValidationResult{
		Status:  isValid,
		Message: "API key is valid",
	}

	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		validationResult.Message = "API key is invalid or expired"
	} else {
		w.WriteHeader(http.StatusOK)
	}

	responseJSON, err := json.Marshal(validationResult)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseJSON)
	if err != nil {
		// Handle the error (optional, based on how you want to handle it)
	}
}
