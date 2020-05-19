package rest

// func (a *UserApiService) UserCancelWithdrawal(ctx context.Context, token string) (Transaction, *http.Response, error) {

type UserCheckReferralCodeConf struct {
	ReferralCode string
}

// func (a *UserApiService) UserCheckReferralCode(ctx context.Context, localVarOptionals *UserCheckReferralCodeOpts) (float64, *http.Response, error) {

// func (a *UserApiService) UserCommunicationToken(ctx context.Context, token string, platformAgent string) ([]CommunicationToken, *http.Response, error) {

// func (a *UserApiService) UserConfirm(ctx context.Context, token string) (AccessToken, *http.Response, error) {

// func (a *UserApiService) UserConfirmWithdrawal(ctx context.Context, token string) (Transaction, *http.Response, error) {

// func (a *UserApiService) UserGet(ctx context.Context) (User, *http.Response, error) {

// func (a *UserApiService) UserGetAffiliateStatus(ctx context.Context) (Affiliate, *http.Response, error) {

// func (a *UserApiService) UserGetCommission(ctx context.Context) (UserCommissionsBySymbol, *http.Response, error) {

type UserGetDepositAddressConf struct {
	Currency string
}

// func (a *UserApiService) UserGetDepositAddress(ctx context.Context, localVarOptionals *UserGetDepositAddressOpts) (string, *http.Response, error) {

// func (a *UserApiService) UserGetExecutionHistory(ctx context.Context, symbol string, timestamp time.Time) (interface{}, *http.Response, error) {

type UserGetMarginConf struct {
	Currency string
}

// func (a *UserApiService) UserGetMargin(ctx context.Context, localVarOptionals *UserGetMarginOpts) (Margin, *http.Response, error) {

// func (a *UserApiService) UserGetQuoteFillRatio(ctx context.Context) (QuoteFillRatio, *http.Response, error) {

type UserGetWalletConf struct {
	Currency string
}

// func (a *UserApiService) UserGetWallet(ctx context.Context, localVarOptionals *UserGetWalletOpts) (Wallet, *http.Response, error) {

type UserGetWalletHistoryConf struct {
	Currency string
	Count    float64
	Start    float64
}

type UserGetWalletSummaryConf struct {
	Currency string
}

// func (a *UserApiService) UserGetWalletSummary(ctx context.Context, localVarOptionals *UserGetWalletSummaryOpts) ([]Transaction, *http.Response, error) {

type UserMinWithdrawalFeeConf struct {
	Currency string
}

// func (a *UserApiService) UserMinWithdrawalFee(ctx context.Context, localVarOptionals *UserMinWithdrawalFeeOpts) (interface{}, *http.Response, error) {

type UserRequestWithdrawalConf struct {
	OtpToken string
	Fee      float64
	Text     string
}

// func (a *UserApiService) UserRequestWithdrawal(ctx context.Context, currency string, amount float32, address string, localVarOptionals *UserRequestWithdrawalOpts) (Transaction, *http.Response, error) {

type UserSavePreferencesConf struct {
	Overwrite bool
}

// func (a *UserApiService) UserSavePreferences(ctx context.Context, prefs string, localVarOptionals *UserSavePreferencesOpts) (User, *http.Response, error)
