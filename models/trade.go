package models

import (
	"time"
)

// Trade Individual Trades
// GET /trade Trades
type Trade struct {
	ForeignNotional float64   `json:"foreignNotional,omitempty"`
	GrossValue      float32   `json:"grossValue,omitempty"`
	HomeNotional    float64   `json:"homeNotional,omitempty"`
	Price           float64   `json:"price,omitempty"`
	Side            string    `json:"side,omitempty"`
	Size            float32   `json:"size,omitempty"`
	Symbol          string    `json:"symbol"`
	TickDirection   string    `json:"tickDirection,omitempty"`
	Timestamp       time.Time `json:"timestamp"`
	TrdMatchID      string    `json:"trdMatchID,omitempty"`
}

// TradeBucketed Bucketed Trades
// GET /trade/bucketed previous trades in time buckets
type TradeBucketed struct {
	Close           float64   `json:"close,omitempty"`
	ForeignNotional float64   `json:"foreignNotional,omitempty"`
	High            float64   `json:"high,omitempty"`
	HomeNotional    float64   `json:"homeNotional,omitempty"`
	LastSize        float32   `json:"lastSize,omitempty"`
	Low             float64   `json:"low,omitempty"`
	Open            float64   `json:"open,omitempty"`
	Symbol          string    `json:"symbol"`
	Timestamp       time.Time `json:"timestamp"`
	Trades          float32   `json:"trades,omitempty"`
	Turnover        float32   `json:"turnover,omitempty"`
	Volume          float32   `json:"volume,omitempty"`
	Vwap            float64   `json:"vwap,omitempty"`
}
