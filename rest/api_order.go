package rest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
)

type OrderSnapshot []models.Order

// OrderAmendConf
// OrderID - Order ID
//
// OrigClOrdID - Client Order ID. See POST /order
//
// ClOrdID - Optional new Client Order ID, requires origClOrdID
//
// OrderQty - Optional order quantity in units of the instrument (i.e. contracts)
//
// LeavesQty - Optional leaves quantity in units of the instrument (i.e. contracts). Useful for amending partially filled orders
//
// Price - Optional leaves quantity in units of the instrument (i.e. contracts). Useful for amending partially filled orders
//
// StopPx - Optional trigger price for 'Stop', 'StopLimit', 'MarketIfTouched', and 'LimitIfTouched' orders.
// Use a price below the current price for stop-sell orders and buy-if-touched orders
//
// PegOffsetValue - Optional trailing offset from the current price for 'Stop', 'StopLimit', 'MarketIfTouched', and 'LimitIfTouched' orders;
// use a negative offset for stop-sell orders and buy-if-touched orders. Optional offset from the peg price for 'Pegged' orders
//
// Text - Optional amend annotation. e.g. 'Adjust skew'
type OrderAmendConf struct {
	OrderID        string  `url:"orderID"`
	OrigClOrdID    string  `url:"origClOrdID,omitempty"`
	ClOrdID        string  `url:"clOrdID,omitempty"`
	OrderQty       float32 `url:"orderQty,omitempty"`
	LeavesQty      float32 `url:"leavesQty,omitempty"`
	Price          float64 `url:"price,omitempty"`
	StopPx         float64 `url:"stopPx,omitempty"`
	PegOffsetValue float64 `url:"pegOffsetValue,omitempty"`
	Text           string  `url:"text,omitempty"`
}

func (c Client) AmendOrder(ctx context.Context, f OrderAmendConf) (*models.Order, error) {
	var out models.Order

	return &out, nil
}

type OrderAmendBulkConf struct {
	Orders string `url:"orders,omitempty"`
}

func (c Client) AmendOrderBulk(ctx context.Context, f OrderAmendBulkConf) (OrderSnapshot, error) {
	var out OrderSnapshot

	return out, nil
}

type OrderCancelConf struct {
	OrderID string `url:"orderID,omitempty"`
	ClOrdID string `url:"clOrdID,omitempty"`
	Text    string `url:"text,omitempty"`
}

func (c Client) CancelOrder(ctx context.Context, f OrderCancelConf) (OrderSnapshot, error) {
	var out OrderSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetTrades query: %v", err)
	}

	bs, err := c.Do(ctx, http.MethodDelete, bitmex.Order+"?"+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetTrades get: %v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("#Client.GetTrades unmarshal: %v", err)
	}

	return out, nil
}

type OrderCancelAllConf struct {
	Symbol string `url:"symbol,omitempty"`
	Filter string `url:"filter,omitempty"`
	Text   string `url:"text,omitempty"`
}

func (c Client) CancelOrderAll(ctx context.Context, f OrderCancelAllConf) (OrderSnapshot, error) {
	var out OrderSnapshot

	return out, nil
}

func (c Client) CancelOrderAllAfter(ctx context.Context, timeout float64) (interface{}, error) {
	var out interface{}

	return out, nil
}

// OrderConf
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
type OrderConf struct {
	Columns   string    `url:"columns,omitempty"`
	Count     float32   `url:"count,omitempty"`
	EndTime   time.Time `url:"endTime,omitempty"`
	Filter    string    `url:"filter,omitempty"`
	Reverse   bool      `url:"reverse,omitempty"`
	Start     float32   `url:"start,omitempty"`
	StartTime time.Time `url:"startTime,omitempty"`
	Symbol    string    `url:"symbol,omitempty"`
}

func (c Client) GetOrders(ctx context.Context, f OrderConf) (OrderSnapshot, error) {
	var out OrderSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetTrades query: %v", err)
	}

	bs, err := c.Do(ctx, http.MethodGet, bitmex.Order+"?"+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetTrades get: %v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("#Client.GetTrades unmarshal: %v", err)
	}

	return out, nil
}

// OrderNewConf
//
// ClOrdID - Optional Client Order ID. This clOrdID will come back on the order and any related executions
//
// DisplayQty - Optional quantity to display in the book. Use 0 for a fully hidden order.
//
// ExecInst - Optional execution instructions. Valid options: ParticipateDoNotInitiate, AllOrNone, MarkPrice, IndexPrice, LastPrice, Close, ReduceOnly, Fixed.
// 'AllOrNone' instruction requires displayQty to be 0.
// 'MarkPrice', 'IndexPrice' or 'LastPrice' instruction valid for 'Stop', 'StopLimit', 'MarketIfTouched', and 'LimitIfTouched' orders
//
// OrderQty - Order quantity in units of the instrument (i.e. contracts
//
// OrdType - Order type. Valid options: Market, Limit, Stop, StopLimit, MarketIfTouched, LimitIfTouched, Pegged.
// Defaults to 'Limit' when price is specified. Defaults to 'Stop' when stopPx is specified. Defaults to 'StopLimit' when price and stopPx are specified
//
// PegOffsetValue - Optional trailing offset from the current price for 'Stop', 'StopLimit', 'MarketIfTouched', and 'LimitIfTouched' orders;
// use a negative offset for stop-sell orders and buy-if-touched orders. Optional offset from the peg price for 'Pegged' orders
//
// PegPriceType - Optional peg price type. Valid options: LastPeg, MidPricePeg, MarketPeg, PrimaryPeg, TrailingStopPeg
//
// Price - Optional limit price for 'Limit', 'StopLimit', and 'LimitIfTouched' orders
//
// Side - Order side. Valid options: Buy, Sell. Defaults to 'Buy' unless orderQty is negative
//
// Symbol - (MUST) Instrument symbol. e.g. 'XBTUSD'
//
// StopPx - Optional trigger price for 'Stop', 'StopLimit', 'MarketIfTouched', and 'LimitIfTouched' orders.
// Use a price below the current price for stop-sell orders and buy-if-touched orders.
// Use execInst of 'MarkPrice' or 'LastPrice' to define the current price used for triggering
//
// Text - Optional order annotation. e.g. 'Take profit'
//
// TimeInForce - Time in force. Valid options: Day, GoodTillCancel, ImmediateOrCancel, FillOrKill.
// Defaults to 'GoodTillCancel' for 'Limit', 'StopLimit', and 'LimitIfTouched' orders
type OrderNewConf struct {
	ClOrdID        string  `json:"clOrdID,omitempty"`
	DisplayQty     float32 `json:"displayQty,omitempty"`
	ExecInst       string  `json:"execInst,omitempty"`
	OrderQty       float32 `json:"orderQty,omitempty"`
	OrdType        string  `json:"ordType,omitempty"`
	PegOffsetValue float64 `json:"pegOffsetValue,omitempty"`
	PegPriceType   string  `json:"pegPriceType,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Side           string  `json:"side,omitempty"`
	StopPx         float64 `json:"stopPx,omitempty"`
	Symbol         string  `json:"symbol,omitempty"`
	Text           string  `json:"text,omitempty"`
	TimeInForce    string  `json:"timeInForce,omitempty"`
}

func (c Client) NewOrder(ctx context.Context, f OrderNewConf) (*models.Order, error) {
	var out models.Order

	f.ClOrdID += base64.StdEncoding.EncodeToString(uuid.New().NodeID())
	fmt.Println("ORDER ID:", f.ClOrdID)

	data, err := json.Marshal(f)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetTrades query: %v", err)
	}

	bs, err := c.Do(ctx, http.MethodPost, bitmex.Order, data)
	if err != nil {
		return nil, fmt.Errorf("#Client.NewOrder get: %v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("#Client.NewOrder unmarshal: %v", err)
	}

	return &out, nil
}

type OrderNewBulkConf struct {
	Orders string `url:"orders,omitempty"`
}

func (c Client) NewOrderBulk(ctx context.Context, f OrderNewBulkConf) (OrderSnapshot, error) {
	var out OrderSnapshot

	return out, nil
}
