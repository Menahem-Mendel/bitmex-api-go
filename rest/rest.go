package rest

// endpoints
const (
	announcement       = "announcement"        // GET site announcement
	announcementUrgent = "announcement/urgent" // GET urgent (banner) announcement
	apiKey             = "apiKey"              // GET api key
	chat               = "chat"                //
	// GET chat messages
	// POST Send a chat message
	chatChannels               = "chat/channels"               // GET available channels
	chatConnected              = "chat/connected"              // GET connected users
	execution                  = "execution"                   // GET all raw execution for your account
	executionTradeHistory      = "exection/tradeHistory"       // GET all balance-affecting executions. This includes each trade, insurance charge, and settlement
	funding                    = "funding"                     // GET funding history
	globalNotification         = "globalNotification"          // GET your current GlobalNotifications
	instrument                 = "instrument"                  // GET instruments
	instrumentActive           = "instrument/active"           // GET all active instruments and instruments that have expired in <24hrs
	instrumentActiveAndIndices = "instrument/activeAndIndices" // GET Helper method. Gets all active instruments and all indices. This is a join of the result of /indices and /active
	instrumentActiveIntervals  = "instrument/activeIntervals"  // GET Return all active contract series and interval pairs
	instrumentCompositeIndex   = "instrument/compositeIndex"   // GET Show constituent parts of an index
	instrumentIndices          = "instrument/indices"          // GET all price indices
	insurance                  = "insurance"                   // GET insurance fund history
	leaderboard                = "leaderboard"                 // GET current leaderboard
	leaderboardName            = "leaderboard/name"            // GET your alias on the leaderboard
	liquidation                = "liquidation"                 // GET liquidation orders
	order                      = "order"                       //
	// GET your orders
	// POST Create a new order
	// PUT Amend the quantity or price of an open order
	// DELETE Cancel order(s). Send multiple order IDs to cancel in bulk
	orderAll  = "order/all"  // DELETE all your orders
	orderBulk = "order/bulk" //
	// POST Create multiple new orders for the same symbol
	// PUT Amend multiple orders for the same symbol
	orderCancelAllAfter    = "order/cancelAllAfter"    // POST Automatically cancel all your orders after a specified timeout
	orderClosePosition     = "order/closePosition"     // POST Close a position. [Deprecated, use POST /order with execInst: 'Close']
	orderBookL2            = "orderBook/L2"            // GET current orderbook in vertical format
	position               = "position"                // GET your positions
	positionIsolate        = "position/isolate"        // POST Enable isolated margin or cross margin per-position
	positionLeverage       = "position/leverage"       // POST Choose leverage for a position
	positionRiskLimit      = "position/riskLimit"      // POST Update your risk limit
	positionTransferMargin = "position/transferMargin" // POST Transfer equity in or out of a position
	schema                 = "schema"                  // GET model schemata for data objects returned by this API
	schemaWebsocketHelp    = "schema/websocketHelp"    // GET Returns help text & subject list for websocket usage
	settlement             = "settlement"              // GET settlement history
	stats                  = "stats"                   // GET exchange-wide and per-series turnover and volume statistics
	statsHistory           = "stats/history"           // GET historical exchange-wide and per-series turnover and volume statistics
	statsHistoryUSD        = "stats/historyUSD"        // GET a summary of exchange statistics in USD
	trade                  = "trade"                   // GET Trades
	tradeBucketed          = "trade/bucketed"          // GET previous trades in time buckets
	user                   = "user"                    // GET your user model
	userAffiliateStatus    = "user/affiliateStatus"    // GET your current affiliate/referral status
	userCancelWithdrawal   = "user/cancelWithdrawal"   // POST Cancel a withdrawal
	userCheckReferralCode  = "user/checkReferralCode"  // GET Check if a referral code is valid
	userCommission         = "user/commission"         // GET your account's commission status
	userCommunicationToken = "user/communicationToken" // POST Register your communication token for mobile clients
	userConfirmEmail       = "user/confirmEmail"       // POST Confirm your email address with a token
	userConfirmWithdrawal  = "user/confirmWithdrawal"  // POST Confirm a withdrawal
	userDepositAddress     = "user/depositAddress"     // GET a deposit address
	userExecutionHistory   = "user/executionHistory"   // GET the execution history by day
	userLogout             = "user/logout"             // POST Log out of BitMEX
	userMargin             = "user/margin"             // GET your account's margin status. Send a currency of "all" to receive an array of all supported currencies
	userMinWithdrawalFee   = "user/minWithdrawalFee"   // GET the minimum withdrawal fee for a currency
	userPreferences        = "user/preferences"        // POST Save user preferences
	userQuoteFillRatio     = "user/quoteFillRatio"     // GET 7 days worth of Quote Fill Ratio statistics
	userRequestWithdrawal  = "user/requestWithdrawal"  // POST Request a withdrawal to an external wallet
	userWallet             = "user/wallet"             // GET your current wallet information
	userWalletHistory      = "user/walletHistory"      // GET a history of all of your wallet transactions (deposits, withdrawals, PNL)
	userWalletSummary      = "user/walletSummary"      // GET a summary of all of your wallet transactions (deposits, withdrawals, PNL)
	userEvent              = "userEvent"               // GET your user events
)
