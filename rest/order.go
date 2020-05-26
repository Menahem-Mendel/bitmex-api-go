package rest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type OrderService struct {
	request
}

// OrderSnapshot slice of orders
type OrderSnapshot []*models.Order

// OrderAmendFilter amend order filter configuration
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
type OrderAmendFilter struct {
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
func (o OrderService) Amend(f OrderAmendFilter) (*models.Order, error) {
	var out *models.Order

	data, err := json.Marshal(f)
	if err != nil {
		return nil, errors.Wrapf(err, "can't marshal (%v) to json", f)
	}

	bs, err := o.put(order, data)
	if err != nil {
		return nil, errors.Wrapf(err, "can't put order (uri = %q, data = [%s])", order, data)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

// AmendBulk amend multiple orders for the same symbols
func (o OrderService) AmendBulk(orders string) (OrderSnapshot, error) {
	var out OrderSnapshot

	data := bytes.TrimSpace([]byte(orders))

	bs, err := o.put(orderBulk, data)
	if err != nil {
		return nil, errors.Wrapf(err, "can't put orders (uri = %q, data = [%s])", orderBulk, data)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

// OrderCancelFilter cancel order filter configuration
//
// OrderID - Order ID(s).
//
// ClOrdID - Client Order ID(s). See POST /order
//
// Text - Optional cancellation annotation. e.g. 'Spread Exceeded'
type OrderCancelFilter struct {
	OrderID string `url:"orderID,omitempty"`
	ClOrdID string `url:"clOrdID,omitempty"`
	Text    string `url:"text,omitempty"`
}

// Cancel order(s). Send multiple order IDs to cancel in bulk
func (o OrderService) Cancel(f OrderCancelFilter) (OrderSnapshot, error) {
	var out OrderSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, errors.Wrapf(err, "can't parse (%v) to url value", f)
	}

	uri := order + "?" + params.Encode()

	bs, err := o.delete(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "can't delete order (uri = %q)", uri)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

// OrderCancelAllFilter cancel order filter configuration
//
// Symbol - Optional symbol. If provided, only cancels orders for that symbol
//
// Filter - Optional filter for cancellation. Use to only cancel some orders, e.g. {"side": "Buy"}
//
// Text - Optional cancellation annotation. e.g. 'Spread Exceeded
type OrderCancelAllFilter struct {
	Symbol string `url:"symbol,omitempty"`
	Filter string `url:"filter,omitempty"`
	Text   string `url:"text,omitempty"`
}

// CancelAll cancels all of your orders
func (o OrderService) CancelAll(f OrderCancelAllFilter) (OrderSnapshot, error) {
	var out OrderSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, errors.Wrapf(err, "can't parse (%v) to url value", f)
	}

	uri := orderAll + "?" + params.Encode()

	bs, err := o.delete(uri)
	if err != nil {
		return nil, errors.Wrap(err, "can't delete orders")
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

// CancelAllAfter automatically cancel all your orders after a specified timeout
func (o OrderService) CancelAllAfter(timeout float64) (interface{}, error) {
	var out interface{}

	var params url.Values

	if timeout != 0. {
		params = make(url.Values)
		params.Set("timeout", fmt.Sprintf("%f", timeout))
	}

	uri := orderCancelAllAfter + "?" + params.Encode()

	bs, err := o.delete(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "can't delete orders (uri = %q)", uri)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

// OrderFilter get order filter configuration
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
type OrderFilter struct {
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
func (o OrderService) Get(f OrderFilter) (OrderSnapshot, error) {
	var out OrderSnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, errors.Wrapf(err, "can't parse (%v) to url value", f)
	}

	uri := order + "?" + params.Encode()

	bs, err := o.get(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get orders (uri = %q)", uri)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

// OrderNewFilter order filter configuration
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
type OrderNewFilter struct {
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
func (o OrderService) New(f OrderNewFilter) (*models.Order, error) {
	var out *models.Order

	f.ClOrdID = encode(f.ClOrdID)

	data, err := json.Marshal(f)
	if err != nil {
		return nil, errors.Wrapf(err, "can't marshal json")
	}

	bs, err := o.post(order, data)
	if err != nil {
		return nil, errors.Wrapf(err, "can't post %s", order)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

// NewBulk create multiple new orders for the same symbol
// orders - array of orders
func (o OrderService) NewBulk(orders string) (OrderSnapshot, error) {
	var out OrderSnapshot

	data := bytes.TrimSpace([]byte(orders))

	bs, err := o.post(orderBulk, data)
	if err != nil {
		return nil, errors.Wrapf(err, "can't post %s", orderBulk)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrapf(err, "can't unmarshal response body to type %T", out)
	}

	return out, nil
}

func encode(str string) string {
	return str + strings.ReplaceAll(base64.StdEncoding.EncodeToString(uuid.New().NodeID()), "=", "-")
}
