package models

import "time"

// Stats Exchange Statistics
// GET /stats exchange-wide and per-series turnover and volume statistics
type Stats struct {
	RootSymbol   string  `json:"rootSymbol"`
	Currency     string  `json:"currency,omitempty"`
	Volume24h    float32 `json:"volume24h,omitempty"`
	Turnover24h  float32 `json:"turnover24h,omitempty"`
	OpenInterest float32 `json:"openInterest,omitempty"`
	OpenValue    float32 `json:"openValue,omitempty"`
}

// StatsUSD Exchange Statistics
// GET /stats/historyUSD a summary of exchange statistics in USD
type StatsUSD struct {
	RootSymbol   string  `json:"rootSymbol"`
	Currency     string  `json:"currency,omitempty"`
	Turnover24h  float32 `json:"turnover24h,omitempty"`
	Turnover30d  float32 `json:"turnover30d,omitempty"`
	Turnover365d float32 `json:"turnover365d,omitempty"`
	Turnover     float32 `json:"turnover,omitempty"`
}

// StatsHistory Exchange Statistics
// GET /stats/history historical exchange-wide and per-series turnover and volume statistics
type StatsHistory struct {
	Date       time.Time `json:"date"`
	RootSymbol string    `json:"rootSymbol"`
	Currency   string    `json:"currency,omitempty"`
	Volume     float32   `json:"volume,omitempty"`
	Turnover   float32   `json:"turnover,omitempty"`
}
