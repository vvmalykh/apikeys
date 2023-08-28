package service

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"time"

	// Assume you have an import alias for your model package
	m "apikeys/internal/app/apikeys/model"
)

type APIKeyService struct {
	DB *sql.DB
}

const CurrentHashVersion = 1

const HoursInDay = 24
const DefaultApiKeyLifeTimeDays = 30
const DefaultApiKeyLifeTimeHours = DefaultApiKeyLifeTimeDays * HoursInDay

// GenerateAPIKey creates a new API key and associated metadata.
func (h *APIKeyService) GenerateAPIKey() (string, error) {
	apiKeyString, err := generateRandomString(32)
	if err != nil {
		return "", err
	}

	hashedKey, err := generateAPIKeyHash(apiKeyString)
	if err != nil {
		return "", err
	}

	dateCreated := time.Now()
	dateExpire := dateCreated.Add(DefaultApiKeyLifeTimeHours * time.Hour)

	// Set API key metadata
	apiKey := m.APIKey{
		Hash:          hashedKey,
		CreatedAt:     dateCreated,
		ExpireAt:      dateExpire,
		UpdatedAt:     dateCreated,
		IsActive:      true,
		HashVersionId: CurrentHashVersion,
	}

	// Persist API key to database
	_, err = h.DB.Exec(`INSERT INTO api_keys (hash, created_at, expire_at, updated_at, is_active, hash_version_id) VALUES ($1, $2, $3, $4, $5, $6)`,
		apiKey.Hash, apiKey.CreatedAt, apiKey.ExpireAt, apiKey.UpdatedAt, apiKey.IsActive, apiKey.HashVersionId)

	if err != nil {
		return "", err
	}

	return apiKeyString, nil
}

func generateRandomString(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func generateAPIKeyHash(apiKeyString string) (string, error) {
	hashedArray := sha256.Sum256([]byte(apiKeyString))
	hashedKey := hex.EncodeToString(hashedArray[:])

	return hashedKey, nil
}
