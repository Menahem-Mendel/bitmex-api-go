package models

import (
	"time"
)

// Insurance Insurance Fund Data
// GET /insurance insurance fund history
type Insurance struct {
	Currency      string    `json:"currency"`
	Timestamp     time.Time `json:"timestamp"`
	WalletBalance float32   `json:"walletBalance,omitempty"`
}
