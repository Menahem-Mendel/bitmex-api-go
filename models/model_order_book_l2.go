package models

// OrderBookL2 Level 2 Book Data
// GET orderBook/L2 current orderbook in vertical format
type OrderBookL2 struct {
	Symbol string  `json:"symbol"`
	ID     float32 `json:"id"`
	Side   string  `json:"side"`
	Size   float32 `json:"size,omitempty"`
	Price  float64 `json:"price,omitempty"`
}
