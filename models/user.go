package models

import (
	"time"
)

// User Account Operations
// GET /user your user model
type User struct {
	ID           float32          `json:"id,omitempty"`
	OwnerID      float32          `json:"ownerId,omitempty"`
	Firstname    string           `json:"firstname,omitempty"`
	Lastname     string           `json:"lastname,omitempty"`
	Username     string           `json:"username"`
	Email        string           `json:"email"`
	Phone        string           `json:"phone,omitempty"`
	Created      time.Time        `json:"created,omitempty"`
	LastUpdated  time.Time        `json:"lastUpdated,omitempty"`
	Preferences  *UserPreferences `json:"preferences,omitempty"`
	TFAEnabled   string           `json:"TFAEnabled,omitempty"`
	AffiliateID  string           `json:"affiliateID,omitempty"`
	PgpPubKey    string           `json:"pgpPubKey,omitempty"`
	Country      string           `json:"country,omitempty"`
	GeoipCountry string           `json:"geoipCountry,omitempty"`
	GeoipRegion  string           `json:"geoipRegion,omitempty"`
	Typ          string           `json:"typ,omitempty"`
}

// Transaction
// GET /user/walletHistory a history of all of your wallet transactions (deposits, withdrawals, PNL)
// GET /user/walletSummary a summary of all of your wallet transactions (deposits, withdrawals, PNL)
// POST /user/requestWithdrawal Request a withdrawal to an external wallet
// POST /user/cancelWithdrawal Cancel a withdrawal
type Transaction struct {
	TransactID     string    `json:"transactID"`
	Account        float32   `json:"account,omitempty"`
	Currency       string    `json:"currency,omitempty"`
	TransactType   string    `json:"transactType,omitempty"`
	Amount         float32   `json:"amount,omitempty"`
	Fee            float32   `json:"fee,omitempty"`
	TransactStatus string    `json:"transactStatus,omitempty"`
	Address        string    `json:"address,omitempty"`
	Tx             string    `json:"tx,omitempty"`
	Text           string    `json:"text,omitempty"`
	TransactTime   time.Time `json:"transactTime,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
}

// Affiliate
// GET /user/affiliateStatus your current affiliate/referral status
type Affiliate struct {
	Account          float32   `json:"account"`
	Currency         string    `json:"currency"`
	PrevPayout       float32   `json:"prevPayout,omitempty"`
	PrevTurnover     float32   `json:"prevTurnover,omitempty"`
	PrevComm         float32   `json:"prevComm,omitempty"`
	PrevTimestamp    time.Time `json:"prevTimestamp,omitempty"`
	ExecTurnover     float32   `json:"execTurnover,omitempty"`
	ExecComm         float32   `json:"execComm,omitempty"`
	TotalReferrals   float32   `json:"totalReferrals,omitempty"`
	TotalTurnover    float32   `json:"totalTurnover,omitempty"`
	TotalComm        float32   `json:"totalComm,omitempty"`
	PayoutPcnt       float64   `json:"payoutPcnt,omitempty"`
	PendingPayout    float32   `json:"pendingPayout,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	ReferrerAccount  float64   `json:"referrerAccount,omitempty"`
	ReferralDiscount float64   `json:"referralDiscount,omitempty"`
	AffiliatePayout  float64   `json:"affiliatePayout,omitempty"`
}

// CommunicationToken
// POST /user/communicationToken Register your communication token for mobile clients
type CommunicationToken struct {
	ID          string  `json:"id"`
	UserID      float32 `json:"userId"`
	DeviceToken string  `json:"deviceToken"`
	Channel     string  `json:"channel"`
}

// Margin
// GET /user/margin your account's margin status. Send a currency of "all" to receive an array of all supported currencies
type Margin struct {
	Account            float32   `json:"account"`
	Currency           string    `json:"currency"`
	RiskLimit          float32   `json:"riskLimit,omitempty"`
	PrevState          string    `json:"prevState,omitempty"`
	State              string    `json:"state,omitempty"`
	Action             string    `json:"action,omitempty"`
	Amount             float32   `json:"amount,omitempty"`
	PendingCredit      float32   `json:"pendingCredit,omitempty"`
	PendingDebit       float32   `json:"pendingDebit,omitempty"`
	ConfirmedDebit     float32   `json:"confirmedDebit,omitempty"`
	PrevRealisedPnl    float32   `json:"prevRealisedPnl,omitempty"`
	PrevUnrealisedPnl  float32   `json:"prevUnrealisedPnl,omitempty"`
	GrossComm          float32   `json:"grossComm,omitempty"`
	GrossOpenCost      float32   `json:"grossOpenCost,omitempty"`
	GrossOpenPremium   float32   `json:"grossOpenPremium,omitempty"`
	GrossExecCost      float32   `json:"grossExecCost,omitempty"`
	GrossMarkValue     float32   `json:"grossMarkValue,omitempty"`
	RiskValue          float32   `json:"riskValue,omitempty"`
	TaxableMargin      float32   `json:"taxableMargin,omitempty"`
	InitMargin         float32   `json:"initMargin,omitempty"`
	MaintMargin        float32   `json:"maintMargin,omitempty"`
	SessionMargin      float32   `json:"sessionMargin,omitempty"`
	TargetExcessMargin float32   `json:"targetExcessMargin,omitempty"`
	VarMargin          float32   `json:"varMargin,omitempty"`
	RealisedPnl        float32   `json:"realisedPnl,omitempty"`
	UnrealisedPnl      float32   `json:"unrealisedPnl,omitempty"`
	IndicativeTax      float32   `json:"indicativeTax,omitempty"`
	UnrealisedProfit   float32   `json:"unrealisedProfit,omitempty"`
	SyntheticMargin    float32   `json:"syntheticMargin,omitempty"`
	WalletBalance      float32   `json:"walletBalance,omitempty"`
	MarginBalance      float32   `json:"marginBalance,omitempty"`
	MarginBalancePcnt  float64   `json:"marginBalancePcnt,omitempty"`
	MarginLeverage     float64   `json:"marginLeverage,omitempty"`
	MarginUsedPcnt     float64   `json:"marginUsedPcnt,omitempty"`
	ExcessMargin       float32   `json:"excessMargin,omitempty"`
	ExcessMarginPcnt   float64   `json:"excessMarginPcnt,omitempty"`
	AvailableMargin    float32   `json:"availableMargin,omitempty"`
	WithdrawableMargin float32   `json:"withdrawableMargin,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
	GrossLastValue     float32   `json:"grossLastValue,omitempty"`
	Commission         float64   `json:"commission,omitempty"`
}

// UserPreferences
// POST /user/preferences Save user preferences
type UserPreferences struct {
	AlertOnLiquidations     bool         `json:"alertOnLiquidations,omitempty"`
	AnimationsEnabled       bool         `json:"animationsEnabled,omitempty"`
	AnnouncementsLastSeen   time.Time    `json:"announcementsLastSeen,omitempty"`
	ChatChannelID           float64      `json:"chatChannelID,omitempty"`
	ColorTheme              string       `json:"colorTheme,omitempty"`
	Currency                string       `json:"currency,omitempty"`
	Debug                   bool         `json:"debug,omitempty"`
	DisableEmails           []string     `json:"disableEmails,omitempty"`
	DisablePush             []string     `json:"disablePush,omitempty"`
	HideConfirmDialogs      []string     `json:"hideConfirmDialogs,omitempty"`
	HideConnectionModal     bool         `json:"hideConnectionModal,omitempty"`
	HideFromLeaderboard     bool         `json:"hideFromLeaderboard,omitempty"`
	HideNameFromLeaderboard bool         `json:"hideNameFromLeaderboard,omitempty"`
	HideNotifications       []string     `json:"hideNotifications,omitempty"`
	Locale                  string       `json:"locale,omitempty"`
	MsgsSeen                []string     `json:"msgsSeen,omitempty"`
	OrderBookBinning        *interface{} `json:"orderBookBinning,omitempty"`
	OrderBookType           string       `json:"orderBookType,omitempty"`
	OrderClearImmediate     bool         `json:"orderClearImmediate,omitempty"`
	OrderControlsPlusMinus  bool         `json:"orderControlsPlusMinus,omitempty"`
	ShowLocaleNumbers       bool         `json:"showLocaleNumbers,omitempty"`
	Sounds                  []string     `json:"sounds,omitempty"`
	StrictIPCheck           bool         `json:"strictIPCheck,omitempty"`
	StrictTimeout           bool         `json:"strictTimeout,omitempty"`
	TickerGroup             string       `json:"tickerGroup,omitempty"`
	TickerPinned            bool         `json:"tickerPinned,omitempty"`
	TradeLayout             string       `json:"tradeLayout,omitempty"`
}

// AccessToken
// POST user/confirmEmail Confirm your email address with a token
// time to live in seconds (2 weeks by default)
type AccessToken struct {
	ID      string    `json:"id"`
	TTL     float64   `json:"ttl,omitempty"`
	Created time.Time `json:"created,omitempty"`
	UserID  float64   `json:"userId,omitempty"`
}

// QuoteFillRatio
// GET /user/quoteFillRatio 7 days worth of Quote Fill Ratio statistics
type QuoteFillRatio struct {
	Date                time.Time `json:"date"`
	Account             float64   `json:"account,omitempty"`
	QuoteCount          float64   `json:"quoteCount,omitempty"`
	DealtCount          float64   `json:"dealtCount,omitempty"`
	QuotesMavg7         float64   `json:"quotesMavg7,omitempty"`
	DealtMavg7          float64   `json:"dealtMavg7,omitempty"`
	QuoteFillRatioMavg7 float64   `json:"quoteFillRatioMavg7,omitempty"`
}

// Wallet
// GET /user/wallet your current wallet information
type Wallet struct {
	Account          float32   `json:"account"`
	Currency         string    `json:"currency"`
	PrevDeposited    float32   `json:"prevDeposited,omitempty"`
	PrevWithdrawn    float32   `json:"prevWithdrawn,omitempty"`
	PrevTransferIn   float32   `json:"prevTransferIn,omitempty"`
	PrevTransferOut  float32   `json:"prevTransferOut,omitempty"`
	PrevAmount       float32   `json:"prevAmount,omitempty"`
	PrevTimestamp    time.Time `json:"prevTimestamp,omitempty"`
	DeltaDeposited   float32   `json:"deltaDeposited,omitempty"`
	DeltaWithdrawn   float32   `json:"deltaWithdrawn,omitempty"`
	DeltaTransferIn  float32   `json:"deltaTransferIn,omitempty"`
	DeltaTransferOut float32   `json:"deltaTransferOut,omitempty"`
	DeltaAmount      float32   `json:"deltaAmount,omitempty"`
	Deposited        float32   `json:"deposited,omitempty"`
	Withdrawn        float32   `json:"withdrawn,omitempty"`
	TransferIn       float32   `json:"transferIn,omitempty"`
	TransferOut      float32   `json:"transferOut,omitempty"`
	Amount           float32   `json:"amount,omitempty"`
	PendingCredit    float32   `json:"pendingCredit,omitempty"`
	PendingDebit     float32   `json:"pendingDebit,omitempty"`
	ConfirmedDebit   float32   `json:"confirmedDebit,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	Addr             string    `json:"addr,omitempty"`
	Script           string    `json:"script,omitempty"`
	WithdrawalLock   []string  `json:"withdrawalLock,omitempty"`
}
