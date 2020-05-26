package models

import (
	"time"
)

// Funding Swap Funding History
// GET /funding funding history
type Funding struct {
	Timestamp        time.Time `json:"timestamp"`
	Symbol           string    `json:"symbol"`
	FundingInterval  time.Time `json:"fundingInterval,omitempty"`
	FundingRate      float64   `json:"fundingRate,omitempty"`
	FundingRateDaily float64   `json:"fundingRateDaily,omitempty"`
}
