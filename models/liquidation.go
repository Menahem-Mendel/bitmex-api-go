package models

// Liquidation Active Liquidations
// GET /liquidation liquidation orders
type Liquidation struct {
	OrderID   string  `json:"orderID"`
	Symbol    string  `json:"symbol,omitempty"`
	Side      string  `json:"side,omitempty"`
	Price     float64 `json:"price,omitempty"`
	LeavesQty float32 `json:"leavesQty,omitempty"`
}
