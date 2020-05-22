package bitmex

// exchange url
const (
	BitmexHost        = "www.bitmex.com"
	BitmexHostTestnet = "testnet.bitmex.com"
	BitmexAPIV1       = "/api/v1/"
)

// ContextKey type for context key value
type contextKey string

// ContextAPIKey key for secret value context
const ContextAPIKey contextKey = "apiKey"

// limits
const (
	MAXCount float32 = 1000
)

// order side
const (
	Buy  = "Buy"
	Sell = "Sell"
)

// order execInst
const (
	ParticipateDoNotInitiate = "ParticipateDoNotInitiate"
	AllOrNone                = "AllOrNone"
	MarkPrice                = "MarkPrice"
	IndexPrice               = "IndexPrice"
	LastPrice                = "LastPrice"
	Close                    = "Close"
	ReduceOnly               = "ReduceOnly"
)

// order pegPriceType
const (
	LastPeg         = "LastPeg"
	MidPricePeg     = "MidPricePeg"
	MarketPeg       = "MarketPeg"
	PrimaryPeg      = "PrimaryPeg"
	TrailingStopPeg = "TrailingStopPeg"
)

// order ordType
const (
	Limit           = "Limit"
	Market          = "Market"
	Stop            = "Stop"
	StopLimit       = "StopLimit"
	MarketIfTouched = "MarketIfTouched"
	LimitIfTouched  = "LimitIfTouched"
	Pegged          = "Pegged"
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
