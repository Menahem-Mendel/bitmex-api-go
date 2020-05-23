package rest

import (
	"encoding/json"
	"time"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type TradeService struct {
	request
}

// TradeSnapshot snapshot of trades
type TradeSnapshot []*models.Trade

// TradeFilter query parameters for filtering the Trades
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
type TradeFilter struct {
	Columns   string    `url:"columns,omitempty"`
	Count     float32   `url:"count,omitempty"`
	EndTime   time.Time `url:"endTime,omitempty"`
	Filter    string    `url:"filter,omitempty"`
	Reverse   bool      `url:"reverse,omitempty"`
	Start     float32   `url:"start,omitempty"`
	StartTime time.Time `url:"startTime,omitempty"`
	Symbol    string    `url:"symbol,omitempty"`
}

func (t TradeService) Get(f TradeFilter) (TradeSnapshot, error) {
	var out TradeSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse url value")
	}

	uri := trade + "?" + params.Encode()

	bs, err := t.get(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get %s", trade)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

// TradeBucketedSnapshot snapshot of trade bucketeds
type TradeBucketedSnapshot []models.TradeBucketed

// TradeBucketedFilter query parameters for filtering the TradeBucketeds
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
type TradeBucketedFilter struct {
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

func (t TradeService) GetBucketed(f TradeBucketedFilter) (TradeBucketedSnapshot, error) {
	var out TradeBucketedSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse url value")
	}

	uri := tradeBucketed + "?" + params.Encode()

	bs, err := t.get(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get %s", tradeBucketed)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}
