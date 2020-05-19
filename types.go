package bitmex

// endpoints
const (
	announcement       = "announcement"        // GET site announcement
	announcementUrgent = "announcement/urgent" // GET urgent (banner) announcement
	apiKey             = "apiKey"              // GET api key
	chat               = "chat"                //
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

// ContextKey type for context key value
type contextKey string

// ContextAPIKey key for secret value context
const ContextAPIKey contextKey = "apiKey"

// limits
const (
	MAXCount float32 = 1000
)

// headers
const (
	SignatureAPI = "api-signature"
	KeyAPI       = "api-key"
	ExpiresAPI   = "api-expires"

	RemainingRequestNum = "x-ratelimit-remaining"
	LimitRequestNum     = "x-ratelimit-limit"
	LimitReset          = "x-ratelimit-reset"
)

// exchange url
const (
	BitmexHost        = "www.bitmex.com"
	BitmexHostTestnet = "testnet.bitmex.com"
	BitmexAPIPath     = "api/v1/"
)

// timeframe
const (
	Minute         = "1m"
	ThreeMinutes   = "3m"
	FiveMinutes    = "5m"
	FifteenMinutes = "15m"
	ThirtyMinutes  = "30m"
	Hour           = "1h"
	TwoHours       = "2h"
	ThreeHours     = "3h"
	FourHours      = "4h"
	SixHours       = "6h"
	TwelveHours    = "12h"
	Day            = "1D"
	ThreeDays      = "3D"
	Week           = "1W"
	TwoWeeks       = "2W"
	Month          = "1M"
)

// pairs
const (
	XBTUSD    = "XBTUSD"
	XBTM20    = "XBTM20"
	XBTU20    = "XBTU20"
	ADAM20    = "ADAM20"
	BCHM20    = "BCHM20"
	EOSM20    = "EOSM20"
	ETHUSD    = "ETHUSD"
	ETHUSDM20 = "ETHUSDM20"
	ETHM20    = "ETHM20"
	LTCM20    = "LTCM20"
	TRXM20    = "TRXM20"
	XRPUSD    = "XRPUSD"
	XPRM20    = "XPRM20"
)

// indices
const (
	BXBT       = ".BXBT"
	BXBT30M    = ".BXBT30M"
	BXBTNEXT   = ".BXBT_NEXT"
	XBTBON     = ".XBTBON"
	XBTBON2H   = ".XBTBON2H"
	XBTBON8H   = ".XBTBON8H"
	XBTUSDPI   = ".XBTUSDPI"
	XBTUSDPI2H = ".XBTUSDPI2H"
	XBTUSDPI8H = ".XBTUSDPI8H"

	BADAXBT     = ".BADAXBT"
	BADAXBT30M  = ".BADAXBT30M"
	BADAXBTNEXT = ".BADAXBT_NEXT"

	BBCHXBT     = ".BBCHXBT"
	BBCHXBT30M  = ".BBCHXBT30M"
	BBCHXBTNEXT = ".BBCHXBT_NEXT"

	BVOL    = ".BVOL"
	BVOL24H = ".BVOL24H"
	BVOL7D  = ".BVOL7D"

	BEOSXBT     = ".BEOSXBT"
	BEOSXBT30M  = ".BEOSXBT30M"
	BEOSXBTNEXT = ".BEOSXBT_NEXT"

	BETH        = ".BETH"
	BETH30M     = ".BETH30M"
	BETHXBT     = ".BETHXBT"
	BETHXBT30M  = ".BETHXBT30M"
	BETHXBTNEXT = ".BETHXBT_NEXT"
	BETHNEXT    = ".BETH_NEXT"
	ETHBON      = ".ETHBON"
	ETHBON2H    = ".ETHBON2H"
	ETHBON8H    = ".ETHBON8H"
	ETHUSDPI    = ".ETHUSDPI"
	ETHUSDPI2H  = ".ETHUSDPI2H"
	ETHUSDPI8H  = ".ETHUSDPI8H"

	EVOL7D = ".EVOL7D"

	BLTCXBT     = ".BLTCXBT"
	BLTCXBT30M  = ".BLTCXBT30M"
	BLTCXBTNEXT = ".BLTCXBT_NEXT"

	BTRXXBT     = ".BTRXXBT"
	BTRXXBT30M  = ".BTRXXBT30M"
	BTRXXBTNEXT = ".BTRXXBT_NEXT"

	BXRP        = ".BXRP"
	BXRPXBT     = ".BXRPXBT"
	BXRPXBT30M  = ".BXRPXBT30M"
	BXRPXBTNEXT = ".BXRPXBT_NEXT"
	BXRPNEXT    = ".BXRP_NEXT"
	XRPBON      = ".XRPBON"
	XRPBON2H    = ".XRPBON2H"
	XRPBON8H    = ".XRPBON8H"
	XRPUSDPI    = ".XRPUSDPI"
	XRPUSDPI2H  = ".XRPUSDPI2H"
	XRPUSDPI8H  = ".XRPUSDPI8H"

	USDBON   = ".USDBON"
	USDBON2H = ".USDBON2H"
	USDBON8H = ".USDBON8H"
)
