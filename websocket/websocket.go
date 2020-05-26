package websocket

const (
	Announcement        = "announcement"        // Site announcements
	Chat                = "chat"                // Trollbox chat
	Connected           = "connected"           // Statistics of connected users/bots
	Funding             = "funding"             // Updates of swap funding rates. Sent every funding interval (usually 8hrs)
	Instrument          = "instrument"          // Instrument updates including turnover and bid/ask
	Insurance           = "insurance"           // Daily Insurance Fund updates
	Liquidation         = "liquidation"         // Liquidation orders as they're entered into the book
	OrderBookL2_25      = "orderBookL2_25"      // Top 25 levels of level 2 order book
	OrderBookL2         = "orderBookL2"         // Full level 2 order book
	OrderBook10         = "orderBook10"         // Top 10 levels using traditional full book push
	PublicNotifications = "publicNotifications" // System-wide notifications (used for short-lived messages)
	Quote               = "quote"               // Top level of the book
	QuoteBin1m          = "quoteBin1m"          // 1-minute quote bins
	QuoteBin5m          = "quoteBin5m"          // 5-minute quote bins
	QuoteBin1h          = "quoteBin1h"          // 1-hour quote bins
	QuoteBin1d          = "quoteBin1d"          // 1-day quote bins
	Settlement          = "settlement"          // Settlements
	Trade               = "trade"               // Live trades
	TradeBin1m          = "tradeBin1m"          // 1-minute trade bins
	TradeBin5m          = "tradeBin5m"          // 5-minute trade bins
	TradeBin1h          = "tradeBin1h"          // 1-hour trade bins
	TradeBin1d          = "tradeBin1d"          // 1-day trade bins

	Affiliate            = "affiliate"            // Affiliate status, such as total referred users & payout %
	Execution            = "execution"            // Individual executions; can be multiple per order
	Order                = "order"                // Live updates on your orders
	Margin               = "margin"               // Updates on your current account balance and margin requirements
	Position             = "position"             // Updates on your positions
	PrivateNotifications = "privateNotifications" // Individual notifications - currently not used
	Transact             = "transact"             // Deposit/Withdrawal updates
	Wallet               = "wallet"               // Bitcoin address balance data, including total deposits & withdrawals
)

type WSEvent struct {
	Op   string        `json:"op"`
	Args []interface{} `json:"args"`
}

type WSResponse struct {
	Table       string            `json:"table"`
	Action      string            `json:"action"`
	Data        []interface{}     `json:"data"`
	Keys        []string          `json:"keys,omitempty"`
	ForeignKeys map[string]string `json:"foreignKeys,omitempty"`
	Types       map[string]string `json:"types,omitempty"`
	Filter      struct {
		Account int    `json:"account,omitempty"`
		Symbol  string `json:"symbol,omitempty"`
	} `json:"filter,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type WSError struct {
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}
