package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
)

type TradeService struct {
	RequestFactory
	Synchronous
}

// TradeSnapshot snapshot of trades
type TradeSnapshot []models.Trade

// TradeConf query parameters for filtering the Trades
//
// Columns - Array of column names to fetch. If omitted, will return all columns.
// Note that this method will always return item keys, even when not specified, so you may receive more columns that you expect
//
// Count - Number of results to fetch
//
// EndTime - Ending date filter for results
//
// Filter - Generic table filter. Send JSON key/value pairs, such as {"key": "value"}.
// You can key on individual fields, and do more advanced querying on timestamps. See the Timestamp Docs for more details.
//
// Reverse - If true, will sort results newest first
//
// Start - Starting point for results
//
// StartTime - Starting date filter for results
//
// Symbol - Instrument symbol. Send a bare series (e.g. XBT) to get data for the nearest expiring contract in that series.
// You can also send a timeframe, e.g. XBT:quarterly. Timeframes are nearest, daily, weekly, monthly, quarterly, biquarterly, and perpetual
type TradeConf struct {
	Columns   string    `url:"columns,omitempty"`
	Count     float32   `url:"count,omitempty"`
	EndTime   time.Time `url:"endTime,omitempty"`
	Filter    string    `url:"filter,omitempty"`
	Reverse   bool      `url:"reverse,omitempty"`
	Start     float32   `url:"start,omitempty"`
	StartTime time.Time `url:"startTime,omitempty"`
	Symbol    string    `url:"symbol,omitempty"`
}

// Get get trades
func (t *TradeService) Get(f TradeConf) (*TradeSnapshot, error) {
	var out *TradeSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := NewRequest(http.MethodGet, Trade+"?"+params.Encode(), nil)

	bs, err := t.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}

// TradeBucketedSnapshot snapshot of trade bucketeds
type TradeBucketedSnapshot []models.TradeBucketed

// TradeBucketedConf query parameters for filtering the TradeBucketeds
//
// BinSize - Time interval to bucket by. Available options: [1m,5m,1h,1d]
//
// Columns - Array of column names to fetch. If omitted, will return all columns.
// Note that this method will always return item keys, even when not specified, so you may receive more columns that you expect
//
// Count - Number of results to fetch
//
// EndTime - Ending date filter for results
//
// Filter - Generic table filter. Send JSON key/value pairs, such as {"key": "value"}.
// You can key on individual fields, and do more advanced querying on timestamps. See the Timestamp Docs for more details.
//
// Partial - If true, will send in-progress (incomplete) bins for the current time period
//
// Reverse - If true, will sort results newest first
//
// Start - Starting point for results
//
// StartTime - Starting date filter for results
//
// Symbol - Instrument symbol. Send a bare series (e.g. XBT) to get data for the nearest expiring contract in that series.
// You can also send a timeframe, e.g. XBT:quarterly. Timeframes are nearest, daily, weekly, monthly, quarterly, biquarterly, and perpetual
type TradeBucketedConf struct {
	BinSize   string    `url:"binSize,omitempty"`
	Columns   string    `url:"columns,omitempty"`
	Count     float32   `url:"count,omitempty"`
	EndTime   time.Time `url:"endTime,omitempty"`
	Filter    string    `url:"filter,omitempty"`
	Partial   bool      `url:"partial,omitempty"`
	Reverse   bool      `url:"reverse,omitempty"`
	Start     float32   `url:"start,omitempty"`
	StartTime time.Time `url:"startTime,omitempty"`
	Symbol    string    `url:"symbol,omitempty"`
}

// GetBucketeds get previous trades in time buckets
func (t *TradeService) GetBucketeds(f TradeBucketedConf) (*TradeBucketedSnapshot, error) {
	var out *TradeBucketedSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, err
	}

	req := NewRequest(http.MethodGet, TradeBucketed+"?"+params.Encode(), nil)

	bs, err := t.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, err
	}

	return out, nil
}
