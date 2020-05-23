package rest

import (
	"encoding/json"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type BookService struct {
	request
}

// OrderBoolL2Snapshot an order book L2 slice
type OrderBoolL2Snapshot []*models.OrderBookL2

// OrderBookL2Filter get order book L2 configuration
//
// Symbol - Instrument symbol. Send a series (e.g. XBT) to get data for the nearest contract in that series
//
// Depth - Orderbook depth per side. Send 0 for full depth
type OrderBookL2Filter struct {
	Symbol string  `url:"symbol,omitempty"`
	Depth  float32 `url:"depth,omitempty"`
}

// Get get current orderbook in vertical format
func (b BookService) Get(f OrderBookL2Filter) (OrderBoolL2Snapshot, error) {
	var out OrderBoolL2Snapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, errors.Wrapf(err, "can't parse (%v) to url value", f)
	}

	uri := orderBookL2 + params.Encode()

	bs, err := b.get(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get order book (uri = %q)", uri)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}
