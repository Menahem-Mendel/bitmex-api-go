package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Menahem-Mendel/bitmex-api-go"
	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
)

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

// GetOrderBookL2 get current orderbook in vertical format
func (c Client) GetOrderBookL2(ctx context.Context, f OrderBookL2Conf) (OrderBoolL2Snapshot, error) {
	var out OrderBoolL2Snapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetOrderBookL2 query: %v", err)
	}

	x, err := c.Request(ctx, http.MethodGet, bitmex.OrderBookL2+"?"+params.Encode(), nil, out)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetOrderBookL2: %v", err)
	}

	switch x.(type) {
	case *OrderBoolL2Snapshot:
		out = *x.(*OrderBoolL2Snapshot)
	default:
		return nil, fmt.Errorf("#Client.GetOrderBookL2 type %T isn't supported", x)
	}

	return out, nil
}
