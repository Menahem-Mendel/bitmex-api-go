package rest

// type UserApiService service

// /*
// UserApiService Cancel a withdrawal.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param token

// @return Transaction
// */
// func (a *UserApiService) UserCancelWithdrawal(ctx context.Context, token string) (Transaction, *http.Response, error) {

// /*
// UserApiService Check if a referral code is valid.
// If the code is valid, responds with the referral code&#39;s discount (e.g. &#x60;0.1&#x60; for 10%). Otherwise, will return a 404 or 451 if invalid.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param optional nil or *UserCheckReferralCodeOpts - Optional Parameters:
//      * @param "ReferralCode" (optional.String) -

// @return float64
// */

// type UserCheckReferralCodeOpts struct {
// 	ReferralCode optional.String
// }

// func (a *UserApiService) UserCheckReferralCode(ctx context.Context, localVarOptionals *UserCheckReferralCodeOpts) (float64, *http.Response, error) {

// /*
// UserApiService Register your communication token for mobile clients
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param token
//  * @param platformAgent

// @return []CommunicationToken
// */
// func (a *UserApiService) UserCommunicationToken(ctx context.Context, token string, platformAgent string) ([]CommunicationToken, *http.Response, error) {

// /*
// UserApiService Confirm your email address with a token.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param token

// @return AccessToken
// */
// func (a *UserApiService) UserConfirm(ctx context.Context, token string) (AccessToken, *http.Response, error) {

// /*
// UserApiService Confirm a withdrawal.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param token

// @return Transaction
// */
// func (a *UserApiService) UserConfirmWithdrawal(ctx context.Context, token string) (Transaction, *http.Response, error) {

// /*
// UserApiService Get your user model.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

// @return User
// */
// func (a *UserApiService) UserGet(ctx context.Context) (User, *http.Response, error) {

// /*
// UserApiService Get your current affiliate/referral status.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

// @return Affiliate
// */
// func (a *UserApiService) UserGetAffiliateStatus(ctx context.Context) (Affiliate, *http.Response, error) {

// /*
// UserApiService Get your account&#39;s commission status.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

// @return UserCommissionsBySymbol
// */
// func (a *UserApiService) UserGetCommission(ctx context.Context) (UserCommissionsBySymbol, *http.Response, error) {

// /*
// UserApiService Get a deposit address.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param optional nil or *UserGetDepositAddressOpts - Optional Parameters:
//      * @param "Currency" (optional.String) -

// @return string
// */

// type UserGetDepositAddressOpts struct {
// 	Currency optional.String
// }

// func (a *UserApiService) UserGetDepositAddress(ctx context.Context, localVarOptionals *UserGetDepositAddressOpts) (string, *http.Response, error) {

// /*
// UserApiService Get the execution history by day.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param symbol
//  * @param timestamp

// @return interface{}
// */
// func (a *UserApiService) UserGetExecutionHistory(ctx context.Context, symbol string, timestamp time.Time) (interface{}, *http.Response, error) {

// /*
// UserApiService Get your account&#39;s margin status. Send a currency of \&quot;all\&quot; to receive an array of all supported currencies.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param optional nil or *UserGetMarginOpts - Optional Parameters:
//      * @param "Currency" (optional.String) -

// @return Margin
// */

// type UserGetMarginOpts struct {
// 	Currency optional.String
// }

// func (a *UserApiService) UserGetMargin(ctx context.Context, localVarOptionals *UserGetMarginOpts) (Margin, *http.Response, error) {

// /*
// UserApiService Get 7 days worth of Quote Fill Ratio statistics.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

// @return QuoteFillRatio
// */
// func (a *UserApiService) UserGetQuoteFillRatio(ctx context.Context) (QuoteFillRatio, *http.Response, error) {

// /*
// UserApiService Get your current wallet information.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param optional nil or *UserGetWalletOpts - Optional Parameters:
//      * @param "Currency" (optional.String) -

// @return Wallet
// */

// type UserGetWalletOpts struct {
// 	Currency optional.String
// }

// func (a *UserApiService) UserGetWallet(ctx context.Context, localVarOptionals *UserGetWalletOpts) (Wallet, *http.Response, error) {

// /*
// UserApiService Get a history of all of your wallet transactions (deposits, withdrawals, PNL).
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param optional nil or *UserGetWalletHistoryOpts - Optional Parameters:
//      * @param "Currency" (optional.String) -
//      * @param "Count" (optional.Float64) -  Number of results to fetch.
//      * @param "Start" (optional.Float64) -  Starting point for results.

// @return []Transaction
// */

// type UserGetWalletHistoryOpts struct {
// 	Currency optional.String
// 	Count optional.Float64
// 	Start optional.Float64
// }

// /*
// UserApiService Get a summary of all of your wallet transactions (deposits, withdrawals, PNL).
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param optional nil or *UserGetWalletSummaryOpts - Optional Parameters:
//      * @param "Currency" (optional.String) -

// @return []Transaction
// */

// type UserGetWalletSummaryOpts struct {
// 	Currency optional.String
// }

// func (a *UserApiService) UserGetWalletSummary(ctx context.Context, localVarOptionals *UserGetWalletSummaryOpts) ([]Transaction, *http.Response, error) {

// /*
// UserApiService Get the minimum withdrawal fee for a currency.
// This is changed based on network conditions to ensure timely withdrawals. During network congestion, this may be high. The fee is returned in the same currency.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param optional nil or *UserMinWithdrawalFeeOpts - Optional Parameters:
//      * @param "Currency" (optional.String) -

// @return interface{}
// */

// type UserMinWithdrawalFeeOpts struct {
// 	Currency optional.String
// }

// func (a *UserApiService) UserMinWithdrawalFee(ctx context.Context, localVarOptionals *UserMinWithdrawalFeeOpts) (interface{}, *http.Response, error) {

// /*
// UserApiService Request a withdrawal to an external wallet.
// This will send a confirmation email to the email address on record.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param currency Currency you&#39;re withdrawing. Options: &#x60;XBt&#x60;
//  * @param amount Amount of withdrawal currency.
//  * @param address Destination Address.
//  * @param optional nil or *UserRequestWithdrawalOpts - Optional Parameters:
//      * @param "OtpToken" (optional.String) -  2FA token. Required if 2FA is enabled on your account.
//      * @param "Fee" (optional.Float64) -  Network fee for Bitcoin withdrawals. If not specified, a default value will be calculated based on Bitcoin network conditions. You will have a chance to confirm this via email.
//      * @param "Text" (optional.String) -  Optional annotation, e.g. &#39;Transfer to home wallet&#39;.

// @return Transaction
// */

// type UserRequestWithdrawalOpts struct {
// 	OtpToken optional.String
// 	Fee optional.Float64
// 	Text optional.String
// }

// func (a *UserApiService) UserRequestWithdrawal(ctx context.Context, currency string, amount float32, address string, localVarOptionals *UserRequestWithdrawalOpts) (Transaction, *http.Response, error) {

// /*
// UserApiService Save user preferences.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param prefs
//  * @param optional nil or *UserSavePreferencesOpts - Optional Parameters:
//      * @param "Overwrite" (optional.Bool) -  If true, will overwrite all existing preferences.

// @return User
// */

// type UserSavePreferencesOpts struct {
// 	Overwrite optional.Bool
// }

// func (a *UserApiService) UserSavePreferences(ctx context.Context, prefs string, localVarOptionals *UserSavePreferencesOpts) (User, *http.Response, error)
