package rest

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

// endpoints
const (
	Announcement       = "announcement"        // GET site announcement
	AnnouncementUrgent = "announcement/urgent" // GET urgent (banner) announcement
	APIKey             = "apiKey"              // GET api key
	Chat               = "chat"                //
	// GET chat messages
	// POST Send a chat message
	ChatChannels               = "chat/channels"               // GET available channels
	ChatConnected              = "chat/connected"              // GET connected users
	Execution                  = "execution"                   // GET all raw execution for your account
	ExecutionTradeHistory      = "exection/tradeHistory"       // GET all balance-affecting executions. This includes each trade, insurance charge, and settlement
	Funding                    = "funding"                     // GET funding history
	GlobalNotification         = "globalNotification"          // GET your current GlobalNotifications
	Instrument                 = "instrument"                  // GET instruments
	InstrumentActive           = "instrument/active"           // GET all active instruments and instruments that have expired in <24hrs
	InstrumentActiveAndIndices = "instrument/activeAndIndices" // GET Helper method. Gets all active instruments and all indices. This is a join of the result of /indices and /active
	InstrumentActiveIntervals  = "instrument/activeIntervals"  // GET Return all active contract series and interval pairs
	InstrumentCompositeIndex   = "instrument/compositeIndex"   // GET Show constituent parts of an index
	InstrumentIndices          = "instrument/indices"          // GET all price indices
	Insurance                  = "insurance"                   // GET insurance fund history
	Leaderboard                = "leaderboard"                 // GET current leaderboard
	LeaderboardName            = "leaderboard/name"            // GET your alias on the leaderboard
	Liquidation                = "liquidation"                 // GET liquidation orders
	Order                      = "order"                       //
	// GET your orders
	// POST Create a new order
	// PUT Amend the quantity or price of an open order
	// DELETE Cancel order(s). Send multiple order IDs to cancel in bulk
	OrderAll  = "order/all"  // DELETE all your orders
	OrderBulk = "order/bulk" //
	// POST Create multiple new orders for the same symbol
	// PUT Amend multiple orders for the same symbol
	OrderCancelAllAfter    = "order/cancelAllAfter"    // POST Automatically cancel all your orders after a specified timeout
	OrderClosePosition     = "order/closePosition"     // POST Close a position. [Deprecated, use POST /order with execInst: 'Close']
	OrderBookL2            = "orderBook/L2"            // GET current orderbook in vertical format
	Position               = "position"                // GET your positions
	PositionIsolate        = "position/isolate"        // POST Enable isolated margin or cross margin per-position
	PositionLeverage       = "position/leverage"       // POST Choose leverage for a position
	PositionRiskLimit      = "position/riskLimit"      // POST Update your risk limit
	PositionTransferMargin = "position/transferMargin" // POST Transfer equity in or out of a position
	Schema                 = "schema"                  // GET model schemata for data objects returned by this API
	SchemaWebsocketHelp    = "schema/websocketHelp"    // GET Returns help text & subject list for websocket usage
	Settlement             = "settlement"              // GET settlement history
	Stats                  = "stats"                   // GET exchange-wide and per-series turnover and volume statistics
	StatsHistory           = "stats/history"           // GET historical exchange-wide and per-series turnover and volume statistics
	StatsHistoryUSD        = "stats/historyUSD"        // GET a summary of exchange statistics in USD
	Trade                  = "trade"                   // GET Trades
	TradeBucketed          = "trade/bucketed"          // GET previous trades in time buckets
	User                   = "user"                    // GET your user model
	UserAffiliateStatus    = "user/affiliateStatus"    // GET your current affiliate/referral status
	UserCancelWithdrawal   = "user/cancelWithdrawal"   // POST Cancel a withdrawal
	UserCheckReferralCode  = "user/checkReferralCode"  // GET Check if a referral code is valid
	UserCommission         = "user/commission"         // GET your account's commission status
	UserCommunicationToken = "user/communicationToken" // POST Register your communication token for mobile clients
	UserConfirmEmail       = "user/confirmEmail"       // POST Confirm your email address with a token
	UserConfirmWithdrawal  = "user/confirmWithdrawal"  // POST Confirm a withdrawal
	UserDepositAddress     = "user/depositAddress"     // GET a deposit address
	UserExecutionHistory   = "user/executionHistory"   // GET the execution history by day
	UserLogout             = "user/logout"             // POST Log out of BitMEX
	UserMargin             = "user/margin"             // GET your account's margin status. Send a currency of "all" to receive an array of all supported currencies
	UserMinWithdrawalFee   = "user/minWithdrawalFee"   // GET the minimum withdrawal fee for a currency
	UserPreferences        = "user/preferences"        // POST Save user preferences
	UserQuoteFillRatio     = "user/quoteFillRatio"     // GET 7 days worth of Quote Fill Ratio statistics
	UserRequestWithdrawal  = "user/requestWithdrawal"  // POST Request a withdrawal to an external wallet
	UserWallet             = "user/wallet"             // GET your current wallet information
	UserWalletHistory      = "user/walletHistory"      // GET a history of all of your wallet transactions (deposits, withdrawals, PNL)
	UserWalletSummary      = "user/walletSummary"      // GET a summary of all of your wallet transactions (deposits, withdrawals, PNL)
	UserEvent              = "userEvent"               // GET your user events
)

// headers
const (
	limitRemaining = "x-ratelimit-remaining"
	limitLimit     = "x-ratelimit-limit"
	limitReset     = "x-ratelimit-reset"
)

type Request struct {
	ctx     context.Context
	Method  string
	URI     string
	Data    []byte
	Headers map[string]string
}

type RequestFactory interface {
	NewRequest(ctx context.Context, method, uri string, data []byte) *Request
}

type Synchronous interface {
	Exec(*Request) ([]byte, error)
}

func NewRequest(method, uri string, data []byte) *Request {
	req := Request{
		Method:  method,
		URI:     uri,
		Data:    data,
		Headers: make(map[string]string),
	}

	req.Headers["content-type"] = "application/json"
	req.Headers["accept"] = "application/json"

	return &req
}

type Response struct {
	Response  *http.Response
	Body      []byte
	ErrorBody []byte
}

func NewResponse(resp *http.Response) *Response {
	var body []byte
	var Err []byte

	bs := bufio.NewScanner(resp.Body)
	for bs.Scan() {
		if i, err := strconv.ParseInt(resp.Header.Get(limitRemaining), 10, 64); err != nil {

		} else if i < 1 {
			t, err := strconv.ParseInt(resp.Header.Get(limitReset), 10, 64)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Until(time.Unix(t, 0)))
		}

		if resp.StatusCode != http.StatusOK {
			Err = bs.Bytes()
			body = []byte("{}")
			break
		}
		body = bs.Bytes()
	}
	if err := bs.Err(); err != nil {
		Err = []byte(`"error":"can't read the response body:` + err.Error() + `"}`)
	}

	return &Response{
		Response:  resp,
		Body:      body,
		ErrorBody: Err,
	}
}

func (r *Response) String() string {
	return string(r.Body)
}
