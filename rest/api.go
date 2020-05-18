package rest

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
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

// limits
const (
	MAXCount float32 = 1000
)

// endpoints
const (
	announcement       = "announcement"        // site announcement
	announcementUrgent = "announcement/urgent" // urgent (banner) announcement
	apiKey             = "apiKey"              // api key
	chat               = "chat"                /*
		GET chat messages
		POST
	*/
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
	order                      = "order"                       /*
		GET your orders
		POST Create a new order
		PUT Amend the quantity or price of an open order
		DELETE Cancel order(s). Send multiple order IDs to cancel in bulk
	*/
	orderAll  = "order/all"  // DELETE all your orders
	orderBulk = "order/bulk" /*
		POST Create multiple new orders for the same symbol
		PUT Amend multiple orders for the same symbol
	*/
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

// ContextKey type for context key value
type ContextKey string

// KeyCtxSecret key for secret value context
const KeyCtxSecret ContextKey = "secret"

// headers
const (
	signatureAPI = "api-signature"
	keyAPI       = "api-key"
	expiresAPI   = "api-expires"

	remainingRequestNum = "x-ratelimit-remaining"
	limitRequestNum     = "x-ratelimit-limit"
	limitReset          = "x-ratelimit-reset"
)

var expires = time.Now().Add(time.Hour).Unix()

func (c *Client) get(ctx context.Context, path *url.URL) ([]byte, error) {
	var out []byte

	req, err := request(ctx, http.MethodGet, c.key, path, "")
	if err != nil {
		return out, fmt.Errorf("#Client.get req: %v", err)
	}

	resp, err := do(req)
	if err != nil {
		return out, fmt.Errorf("#Client.get resp: %v", err)
	}
	defer resp.Body.Close()

	out, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("#Client.get reading response: %v", err)
	}

	return out, nil
}

func request(ctx context.Context, method, key string, url *url.URL, dataPost string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("#request: %v", err)
	}

	req.Header.Set("content-type", "application/json; charset=utf-8")

	if v, ok := req.Context().Value(KeyCtxSecret).(string); ok {
		expires := strconv.FormatInt(expires, 10)

		sign := signature(v, method+url.RequestURI()+expires)

		if method == http.MethodPost {
			sign = signature(v, method+url.RequestURI()+expires+dataPost)
		}

		req.Header.Set(signatureAPI, sign)
		req.Header.Set(keyAPI, key)
		req.Header.Set(expiresAPI, expires)
	}

	if req.Context().Err() != nil {
		return nil, fmt.Errorf("#request could not authoarize: %v", err)
	}

	return req, nil
}

func do(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("#do resp: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("#do err reading error message from response")
		}
		return nil, fmt.Errorf("#do got status: %v\n%s", resp.Status, bs)
	}

	// resp.Header.Get(limitRequestNum)
	if resp.Header.Get(remainingRequestNum) == "1" {
		num, err := strconv.ParseInt(resp.Header.Get(limitReset), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("#do parseint: %v", err)
		}

		fmt.Println("Wait Ratelimit Reset...")
		time.Sleep(time.Until(time.Unix(num, 0)))
	}

	return resp, nil
}

func signature(secret, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
