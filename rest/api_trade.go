package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
)

// TradeSnapshot snapshot of Thrades
type TradeSnapshot []models.Trade

// TradeConf query parameters for filtering the Trades
type TradeConf struct {
	Reverse   bool      `url:"reverse,omitempty"`
	Count     float32   `url:"count,omitempty"`
	Start     float32   `url:"start,omitempty"`
	Symbol    string    `url:"symbol,omitempty"`
	Filter    string    `url:"filter,omitempty"`
	Columns   string    `url:"columns,omitempty"`
	StartTime time.Time `url:"startTime,omitempty"`
	EndTime   time.Time `url:"endTime,omitempty"`
}

// GetTrades returns snapshot of Trades
func (c Client) GetTrades(ctx context.Context, f TradeConf) (TradeSnapshot, error) {
	var out TradeSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("#c.GetTrades query: %v", err)
	}

	path, err := c.Base.Parse(trade + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("#GetTrades path: %v", err)
	}

	bs, err := c.get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("#GetTrades get: %v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("#GetTrades unmarshal: %v", err)
	}
	return out, nil
}

// TradeBinSnapshot snapshot of ThradeBins
type TradeBucketedSnapshot []models.TradeBucketed

// TradeBinConf query parameters for filtering the TradeBins
type TradeBucketedConf struct {
	BinSize   string    `url:"binSize,omitempty"` // must
	Partial   bool      `url:"partial,omitempty"`
	Reverse   bool      `url:"reverse,omitempty"`
	Count     float32   `url:"count,omitempty"`
	Start     float32   `url:"start,omitempty"`
	Columns   string    `url:"columns,omitempty"`
	Filter    string    `url:"filter,omitempty"`
	Symbol    string    `url:"symbol,omitempty"`
	EndTime   time.Time `url:"endTime,omitempty"`
	StartTime time.Time `url:"startTime,omitempty"`
}

// GetTradeBins returns snapshot of TradeBins
func (c Client) GetTradeBucketeds(ctx context.Context, f TradeBucketedConf) (TradeBucketedSnapshot, error) {
	var out TradeBucketedSnapshot

	if f.BinSize == "" {
		f.BinSize = Minute
	} else if f.Count > MAXCount {
		f.Count = MAXCount
	}

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("#c.GetTrades query: %v", err)
	}

	path, err := c.Base.Parse(tradeBucketed + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("#GetTrades path: %v", err)
	}

	bs, err := c.get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("#GetTrades get: %v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("#GetTrades unmarshal: %v", err)
	}
	return out, nil
}
