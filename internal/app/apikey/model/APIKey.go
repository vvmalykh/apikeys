package model

import (
	"time"
)

// APIKey represents an API key in the database.
type APIKey struct {
	ID            int       `json:"id" db:"id"`
	Hash          string    `json:"key" db:"key"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	ExpireAt      time.Time `json:"expire_at" db:"expire_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	IsActive      bool      `json:"is_active" db:"is_active"`
	HashVersionId int       `json:"hash_version_id" db:"hash_version_id"`
}
