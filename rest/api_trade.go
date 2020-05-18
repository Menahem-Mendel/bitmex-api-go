package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
)

type TradeService []models.Trade

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

func (c Client) GetTrades(ctx context.Context, f TradeConf) (TradeService, error) {
	var out TradeService

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

type TradeBinService []models.TradeBin

type TradeBinConf struct {
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

func (c Client) GetTradeBins(ctx context.Context, f TradeBinConf) (TradeBinService, error) {
	var out TradeBinService

	if f.BinSize == "" {
		f.BinSize = Minute
	}
	if f.Count > MAXCount {
		f.Count = MAXCount
	}

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("#c.GetTrades query: %v", err)
	}

	path, err := c.Base.Parse(tradeBin + "?" + params.Encode())
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
