package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
)

type BookService struct {
	RequestFactory
	Synchronous
}

// OrderBoolL2Snapshot an order book L2 slice
type OrderBoolL2Snapshot []models.OrderBookL2

// OrderBookL2Conf get order book L2 configuration
//
// Symbol - Instrument symbol. Send a series (e.g. XBT) to get data for the nearest contract in that series
//
// Depth - Orderbook depth per side. Send 0 for full depth
type OrderBookL2Conf struct {
	Symbol string  `url:"symbol,omitempty"`
	Depth  float32 `url:"depth,omitempty"`
}

// Get get current orderbook in vertical format
func (b *BookService) Get(ctx context.Context, f OrderBookL2Conf) (*OrderBoolL2Snapshot, error) {
	var out *OrderBoolL2Snapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := NewRequest(http.MethodGet, OrderBookL2+params.Encode(), nil)

	bs, err := b.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}
