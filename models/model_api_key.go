package models

import (
	"time"
)

// APIKey Persistent API Keys for Developers
// GET /apiKey api key
type APIKey struct {
	ID          string        `json:"id"`
	Secret      string        `json:"secret"`
	Name        string        `json:"name"`
	Nonce       float32       `json:"nonce"`
	Cidr        string        `json:"cidr,omitempty"`
	Permissions []interface{} `json:"permissions,omitempty"`
	Enabled     bool          `json:"enabled,omitempty"`
	UserID      float32       `json:"userId"`
	Created     time.Time     `json:"created,omitempty"`
}
