package rest

// func (c Client) GetUser(ctx context.Context) (*models.User, error) {}

// func (c Client) GetUserAffiliateStatus(ctx context.Context) (*models.Affiliate, error) {}

// func (c Client) DeleteUserWithdrawal(ctx context.Context, token string) (models.Transaction, error) {}

// UserCheckReferralCodeConf
type UserCheckReferralCodeConf struct {
	ReferralCode string
}

// func (c Client) UserCheckReferralCode(ctx context.Context, localVarOptionals *UserCheckReferralCodeOpts) (float64, *http.Response, error) {}

// func (c Client) UserCommunicationToken(ctx context.Context, token string, platformAgent string) ([]CommunicationToken, *http.Response, error) {}

// func (c Client) UserConfirm(ctx context.Context, token string) (AccessToken, *http.Response, error) {}

// func (c Client) UserConfirmWithdrawal(ctx context.Context, token string) (Transaction, *http.Response, error) {}

// func (c Client) UserGetCommission(ctx context.Context) (UserCommissionsBySymbol, *http.Response, error) {}

// UserGetDepositAddressConf
type UserGetDepositAddressConf struct {
	Currency string
}

// func (c Client) UserGetDepositAddress(ctx context.Context, localVarOptionals *UserGetDepositAddressOpts) (string, *http.Response, error) {}

// func (c Client) UserGetExecutionHistory(ctx context.Context, symbol string, timestamp time.Time) (interface{}, *http.Response, error) {}

// UserGetMarginConf
type UserGetMarginConf struct {
	Currency string
}

// func (c Client) UserGetMargin(ctx context.Context, localVarOptionals *UserGetMarginOpts) (Margin, *http.Response, error) {}

// func (c Client) UserGetQuoteFillRatio(ctx context.Context) (QuoteFillRatio, *http.Response, error) {}

// UserGetWalletConf
type UserGetWalletConf struct {
	Currency string
}

// func (c Client) UserGetWallet(ctx context.Context, localVarOptionals *UserGetWalletOpts) (Wallet, *http.Response, error) {}

// UserGetWalletHistoryConf
type UserGetWalletHistoryConf struct {
	Currency string
	Count    float64
	Start    float64
}

// UserGetWalletSummaryConf
type UserGetWalletSummaryConf struct {
	Currency string
}

// func (c Client) UserGetWalletSummary(ctx context.Context, localVarOptionals *UserGetWalletSummaryOpts) ([]Transaction, *http.Response, error) {}

// UserMinWithdrawalFeeConf
type UserMinWithdrawalFeeConf struct {
	Currency string
}

// func (c Client) UserMinWithdrawalFee(ctx context.Context, localVarOptionals *UserMinWithdrawalFeeOpts) (interface{}, *http.Response, error) {}

// UserRequestWithdrawalConf
type UserRequestWithdrawalConf struct {
	OtpToken string
	Fee      float64
	Text     string
}

// func (c Client) UserRequestWithdrawal(ctx context.Context, currency string, amount float32, address string, localVarOptionals *UserRequestWithdrawalOpts) (Transaction, *http.Response, error) {}

// UserSavePreferencesConf
type UserSavePreferencesConf struct {
	Overwrite bool
}

// func (c Client) UserSavePreferences(ctx context.Context, prefs string, localVarOptionals *UserSavePreferencesOpts) (User, *http.Response, error) {}
