package rest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
)

type OrderService struct {
	RequestFactory
	Synchronous
}

// OrderSnapshot slice of orders
type OrderSnapshot []models.Order

// OrderAmendConf amend order configuration
//
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
	OrderID        string  `json:"orderID"`
	OrigClOrdID    string  `json:"origClOrdID,omitempty"`
	ClOrdID        string  `json:"clOrdID,omitempty"`
	OrderQty       float32 `json:"orderQty,omitempty"`
	LeavesQty      float32 `json:"leavesQty,omitempty"`
	Price          float64 `json:"price,omitempty"`
	StopPx         float64 `json:"stopPx,omitempty"`
	PegOffsetValue float64 `json:"pegOffsetValue,omitempty"`
	Text           string  `json:"text,omitempty"`
}

// Amend amend the quantity or price of an open order
func (o *OrderService) Amend(ctx context.Context, f OrderAmendConf) (*models.Order, error) {
	var out *models.Order

	data, err := json.Marshal(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := o.NewRequest(ctx, http.MethodPut, Order, data)

	bs, err := o.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}

// OrderAmendBulkConf an array of orders
type OrderAmendBulkConf struct {
	Orders string `json:"orders,omitempty"`
}

// AmendBulk amend multiple orders for the same symbols
func (o *OrderService) AmendBulk(ctx context.Context, f OrderAmendBulkConf) (*OrderSnapshot, error) {
	var out *OrderSnapshot

	data, err := json.Marshal(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := o.NewRequest(ctx, http.MethodPut, OrderBulk, data)

	bs, err := o.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}

// OrderCancelConf cancel order configuration
//
// OrderID - Order ID(s).
//
// ClOrdID - Client Order ID(s). See POST /order
//
// Text - Optional cancellation annotation. e.g. 'Spread Exceeded'
type OrderCancelConf struct {
	OrderID string `url:"orderID,omitempty"`
	ClOrdID string `url:"clOrdID,omitempty"`
	Text    string `url:"text,omitempty"`
}

// Cancel order(s). Send multiple order IDs to cancel in bulk
func (o *OrderService) Cancel(ctx context.Context, f OrderCancelConf) (*OrderSnapshot, error) {
	var out *OrderSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := o.NewRequest(ctx, http.MethodDelete, Order+params.Encode(), nil)

	bs, err := o.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}

// OrderCancelAllConf cancel order configuration
//
// Symbol - Optional symbol. If provided, only cancels orders for that symbol
//
// Filter - Optional filter for cancellation. Use to only cancel some orders, e.g. {"side": "Buy"}
//
// Text - Optional cancellation annotation. e.g. 'Spread Exceeded
type OrderCancelAllConf struct {
	Symbol string `url:"symbol,omitempty"`
	Filter string `url:"filter,omitempty"`
	Text   string `url:"text,omitempty"`
}

// CancelAll cancels all of your orders
func (o *OrderService) CancelAll(ctx context.Context, f OrderCancelAllConf) (*OrderSnapshot, error) {
	var out *OrderSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := o.NewRequest(ctx, http.MethodDelete, OrderAll+params.Encode(), nil)

	bs, err := o.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}

// CancelAllAfter automatically cancel all your orders after a specified timeout
func (o *OrderService) CancelAllAfter(ctx context.Context, timeout float64) (interface{}, error) {
	var out interface{}

	params, err := query.Values(timeout)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := o.NewRequest(ctx, http.MethodDelete, OrderCancelAllAfter+params.Encode(), nil)

	bs, err := o.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}

// OrderConf get order configuration
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

// Get get your orders
func (o *OrderService) Get(ctx context.Context, f OrderConf) (*OrderSnapshot, error) {
	var out *OrderSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := o.NewRequest(ctx, http.MethodGet, Order+params.Encode(), nil)

	bs, err := o.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}

// OrderNewConf order configuration
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

// New create a new order
func (o *OrderService) New(ctx context.Context, f OrderNewConf) (*models.Order, error) {
	var out *models.Order

	f.ClOrdID += base64.StdEncoding.EncodeToString(uuid.New().NodeID())

	data, err := json.Marshal(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := o.NewRequest(ctx, http.MethodPost, Order, data)

	bs, err := o.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}

// OrderNewBulkConf an array of orders
type OrderNewBulkConf struct {
	Orders string `json:"orders,omitempty"`
}

// NewBulk create multiple new orders for the same symbol
func (o *OrderService) NewBulk(ctx context.Context, f OrderNewBulkConf) (OrderSnapshot, error) {
	var out OrderSnapshot

	data, err := json.Marshal(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := o.NewRequest(ctx, http.MethodPost, Order, data)

	bs, err := o.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}
