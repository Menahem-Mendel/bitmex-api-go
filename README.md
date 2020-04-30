# Go API client for swagger

## REST API for the BitMEX Trading Platform  [View Changelog](/app/apiChangelog)  -  #### Getting Started  Base URI: [https://www.bitmex.com/api/v1](/api/v1)  ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](/app/restAPI).  _All_ table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  _This is only a small subset of what is available, to get you started._  Fill in the parameters and click the `Try it out!` button to try any of these queries.  - [Pricing Data](#!/Quote/Quote_get)  - [Trade Data](#!/Trade/Trade_get)  - [OrderBook Data](#!/OrderBook/OrderBook_getL2)  - [Settlement Data](#!/Settlement/Settlement_get)  - [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  ##### Swagger Specification  [⇩ Download Swagger JSON](swagger.json)  -  ## All API Endpoints  Click to expand a section. 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.2.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://www.bitmex.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APIKeyApi* | [**APIKeyGet**](docs/APIKeyApi.md#apikeyget) | **Get** /apiKey | Get your API Keys.
*AnnouncementApi* | [**AnnouncementGet**](docs/AnnouncementApi.md#announcementget) | **Get** /announcement | Get site announcements.
*AnnouncementApi* | [**AnnouncementGetUrgent**](docs/AnnouncementApi.md#announcementgeturgent) | **Get** /announcement/urgent | Get urgent (banner) announcements.
*ChatApi* | [**ChatGet**](docs/ChatApi.md#chatget) | **Get** /chat | Get chat messages.
*ChatApi* | [**ChatGetChannels**](docs/ChatApi.md#chatgetchannels) | **Get** /chat/channels | Get available channels.
*ChatApi* | [**ChatGetConnected**](docs/ChatApi.md#chatgetconnected) | **Get** /chat/connected | Get connected users.
*ChatApi* | [**ChatNew**](docs/ChatApi.md#chatnew) | **Post** /chat | Send a chat message.
*ExecutionApi* | [**ExecutionGet**](docs/ExecutionApi.md#executionget) | **Get** /execution | Get all raw executions for your account.
*ExecutionApi* | [**ExecutionGetTradeHistory**](docs/ExecutionApi.md#executiongettradehistory) | **Get** /execution/tradeHistory | Get all balance-affecting executions. This includes each trade, insurance charge, and settlement.
*FundingApi* | [**FundingGet**](docs/FundingApi.md#fundingget) | **Get** /funding | Get funding history.
*GlobalNotificationApi* | [**GlobalNotificationGet**](docs/GlobalNotificationApi.md#globalnotificationget) | **Get** /globalNotification | Get your current GlobalNotifications.
*InstrumentApi* | [**InstrumentGet**](docs/InstrumentApi.md#instrumentget) | **Get** /instrument | Get instruments.
*InstrumentApi* | [**InstrumentGetActive**](docs/InstrumentApi.md#instrumentgetactive) | **Get** /instrument/active | Get all active instruments and instruments that have expired in &lt;24hrs.
*InstrumentApi* | [**InstrumentGetActiveAndIndices**](docs/InstrumentApi.md#instrumentgetactiveandindices) | **Get** /instrument/activeAndIndices | Helper method. Gets all active instruments and all indices. This is a join of the result of /indices and /active.
*InstrumentApi* | [**InstrumentGetActiveIntervals**](docs/InstrumentApi.md#instrumentgetactiveintervals) | **Get** /instrument/activeIntervals | Return all active contract series and interval pairs.
*InstrumentApi* | [**InstrumentGetCompositeIndex**](docs/InstrumentApi.md#instrumentgetcompositeindex) | **Get** /instrument/compositeIndex | Show constituent parts of an index.
*InstrumentApi* | [**InstrumentGetIndices**](docs/InstrumentApi.md#instrumentgetindices) | **Get** /instrument/indices | Get all price indices.
*InsuranceApi* | [**InsuranceGet**](docs/InsuranceApi.md#insuranceget) | **Get** /insurance | Get insurance fund history.
*LeaderboardApi* | [**LeaderboardGet**](docs/LeaderboardApi.md#leaderboardget) | **Get** /leaderboard | Get current leaderboard.
*LeaderboardApi* | [**LeaderboardGetName**](docs/LeaderboardApi.md#leaderboardgetname) | **Get** /leaderboard/name | Get your alias on the leaderboard.
*LiquidationApi* | [**LiquidationGet**](docs/LiquidationApi.md#liquidationget) | **Get** /liquidation | Get liquidation orders.
*OrderApi* | [**OrderAmend**](docs/OrderApi.md#orderamend) | **Put** /order | Amend the quantity or price of an open order.
*OrderApi* | [**OrderAmendBulk**](docs/OrderApi.md#orderamendbulk) | **Put** /order/bulk | Amend multiple orders for the same symbol.
*OrderApi* | [**OrderCancel**](docs/OrderApi.md#ordercancel) | **Delete** /order | Cancel order(s). Send multiple order IDs to cancel in bulk.
*OrderApi* | [**OrderCancelAll**](docs/OrderApi.md#ordercancelall) | **Delete** /order/all | Cancels all of your orders.
*OrderApi* | [**OrderCancelAllAfter**](docs/OrderApi.md#ordercancelallafter) | **Post** /order/cancelAllAfter | Automatically cancel all your orders after a specified timeout.
*OrderApi* | [**OrderClosePosition**](docs/OrderApi.md#ordercloseposition) | **Post** /order/closePosition | Close a position. [Deprecated, use POST /order with execInst: &#39;Close&#39;]
*OrderApi* | [**OrderGetOrders**](docs/OrderApi.md#ordergetorders) | **Get** /order | Get your orders.
*OrderApi* | [**OrderNew**](docs/OrderApi.md#ordernew) | **Post** /order | Create a new order.
*OrderApi* | [**OrderNewBulk**](docs/OrderApi.md#ordernewbulk) | **Post** /order/bulk | Create multiple new orders for the same symbol.
*OrderBookApi* | [**OrderBookGetL2**](docs/OrderBookApi.md#orderbookgetl2) | **Get** /orderBook/L2 | Get current orderbook in vertical format.
*PositionApi* | [**PositionGet**](docs/PositionApi.md#positionget) | **Get** /position | Get your positions.
*PositionApi* | [**PositionIsolateMargin**](docs/PositionApi.md#positionisolatemargin) | **Post** /position/isolate | Enable isolated margin or cross margin per-position.
*PositionApi* | [**PositionTransferIsolatedMargin**](docs/PositionApi.md#positiontransferisolatedmargin) | **Post** /position/transferMargin | Transfer equity in or out of a position.
*PositionApi* | [**PositionUpdateLeverage**](docs/PositionApi.md#positionupdateleverage) | **Post** /position/leverage | Choose leverage for a position.
*PositionApi* | [**PositionUpdateRiskLimit**](docs/PositionApi.md#positionupdaterisklimit) | **Post** /position/riskLimit | Update your risk limit.
*QuoteApi* | [**QuoteGet**](docs/QuoteApi.md#quoteget) | **Get** /quote | Get Quotes.
*QuoteApi* | [**QuoteGetBucketed**](docs/QuoteApi.md#quotegetbucketed) | **Get** /quote/bucketed | Get previous quotes in time buckets.
*SchemaApi* | [**SchemaGet**](docs/SchemaApi.md#schemaget) | **Get** /schema | Get model schemata for data objects returned by this API.
*SchemaApi* | [**SchemaWebsocketHelp**](docs/SchemaApi.md#schemawebsockethelp) | **Get** /schema/websocketHelp | Returns help text &amp; subject list for websocket usage.
*SettlementApi* | [**SettlementGet**](docs/SettlementApi.md#settlementget) | **Get** /settlement | Get settlement history.
*StatsApi* | [**StatsGet**](docs/StatsApi.md#statsget) | **Get** /stats | Get exchange-wide and per-series turnover and volume statistics.
*StatsApi* | [**StatsHistory**](docs/StatsApi.md#statshistory) | **Get** /stats/history | Get historical exchange-wide and per-series turnover and volume statistics.
*StatsApi* | [**StatsHistoryUSD**](docs/StatsApi.md#statshistoryusd) | **Get** /stats/historyUSD | Get a summary of exchange statistics in USD.
*TradeApi* | [**TradeGet**](docs/TradeApi.md#tradeget) | **Get** /trade | Get Trades.
*TradeApi* | [**TradeGetBucketed**](docs/TradeApi.md#tradegetbucketed) | **Get** /trade/bucketed | Get previous trades in time buckets.
*UserApi* | [**UserCancelWithdrawal**](docs/UserApi.md#usercancelwithdrawal) | **Post** /user/cancelWithdrawal | Cancel a withdrawal.
*UserApi* | [**UserCheckReferralCode**](docs/UserApi.md#usercheckreferralcode) | **Get** /user/checkReferralCode | Check if a referral code is valid.
*UserApi* | [**UserCommunicationToken**](docs/UserApi.md#usercommunicationtoken) | **Post** /user/communicationToken | Register your communication token for mobile clients
*UserApi* | [**UserConfirm**](docs/UserApi.md#userconfirm) | **Post** /user/confirmEmail | Confirm your email address with a token.
*UserApi* | [**UserConfirmWithdrawal**](docs/UserApi.md#userconfirmwithdrawal) | **Post** /user/confirmWithdrawal | Confirm a withdrawal.
*UserApi* | [**UserGet**](docs/UserApi.md#userget) | **Get** /user | Get your user model.
*UserApi* | [**UserGetAffiliateStatus**](docs/UserApi.md#usergetaffiliatestatus) | **Get** /user/affiliateStatus | Get your current affiliate/referral status.
*UserApi* | [**UserGetCommission**](docs/UserApi.md#usergetcommission) | **Get** /user/commission | Get your account&#39;s commission status.
*UserApi* | [**UserGetDepositAddress**](docs/UserApi.md#usergetdepositaddress) | **Get** /user/depositAddress | Get a deposit address.
*UserApi* | [**UserGetExecutionHistory**](docs/UserApi.md#usergetexecutionhistory) | **Get** /user/executionHistory | Get the execution history by day.
*UserApi* | [**UserGetMargin**](docs/UserApi.md#usergetmargin) | **Get** /user/margin | Get your account&#39;s margin status. Send a currency of \&quot;all\&quot; to receive an array of all supported currencies.
*UserApi* | [**UserGetQuoteFillRatio**](docs/UserApi.md#usergetquotefillratio) | **Get** /user/quoteFillRatio | Get 7 days worth of Quote Fill Ratio statistics.
*UserApi* | [**UserGetWallet**](docs/UserApi.md#usergetwallet) | **Get** /user/wallet | Get your current wallet information.
*UserApi* | [**UserGetWalletHistory**](docs/UserApi.md#usergetwallethistory) | **Get** /user/walletHistory | Get a history of all of your wallet transactions (deposits, withdrawals, PNL).
*UserApi* | [**UserGetWalletSummary**](docs/UserApi.md#usergetwalletsummary) | **Get** /user/walletSummary | Get a summary of all of your wallet transactions (deposits, withdrawals, PNL).
*UserApi* | [**UserLogout**](docs/UserApi.md#userlogout) | **Post** /user/logout | Log out of BitMEX.
*UserApi* | [**UserMinWithdrawalFee**](docs/UserApi.md#userminwithdrawalfee) | **Get** /user/minWithdrawalFee | Get the minimum withdrawal fee for a currency.
*UserApi* | [**UserRequestWithdrawal**](docs/UserApi.md#userrequestwithdrawal) | **Post** /user/requestWithdrawal | Request a withdrawal to an external wallet.
*UserApi* | [**UserSavePreferences**](docs/UserApi.md#usersavepreferences) | **Post** /user/preferences | Save user preferences.
*UserEventApi* | [**UserEventGet**](docs/UserEventApi.md#usereventget) | **Get** /userEvent | Get your user events


## Documentation For Models

 - [AccessToken](docs/AccessToken.md)
 - [Affiliate](docs/Affiliate.md)
 - [Announcement](docs/Announcement.md)
 - [ApiKey](docs/ApiKey.md)
 - [Chat](docs/Chat.md)
 - [ChatChannel](docs/ChatChannel.md)
 - [CommunicationToken](docs/CommunicationToken.md)
 - [ConnectedUsers](docs/ConnectedUsers.md)
 - [ErrorError](docs/ErrorError.md)
 - [Execution](docs/Execution.md)
 - [Funding](docs/Funding.md)
 - [GlobalNotification](docs/GlobalNotification.md)
 - [IndexComposite](docs/IndexComposite.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [Instrument](docs/Instrument.md)
 - [InstrumentInterval](docs/InstrumentInterval.md)
 - [Insurance](docs/Insurance.md)
 - [Leaderboard](docs/Leaderboard.md)
 - [Liquidation](docs/Liquidation.md)
 - [Margin](docs/Margin.md)
 - [ModelError](docs/ModelError.md)
 - [Order](docs/Order.md)
 - [OrderBookL2](docs/OrderBookL2.md)
 - [Position](docs/Position.md)
 - [Quote](docs/Quote.md)
 - [QuoteFillRatio](docs/QuoteFillRatio.md)
 - [Settlement](docs/Settlement.md)
 - [Stats](docs/Stats.md)
 - [StatsHistory](docs/StatsHistory.md)
 - [StatsUsd](docs/StatsUsd.md)
 - [Trade](docs/Trade.md)
 - [TradeBin](docs/TradeBin.md)
 - [Transaction](docs/Transaction.md)
 - [User](docs/User.md)
 - [UserCommissionsBySymbol](docs/UserCommissionsBySymbol.md)
 - [UserEvent](docs/UserEvent.md)
 - [UserPreferences](docs/UserPreferences.md)
 - [Wallet](docs/Wallet.md)
 - [XAny](docs/XAny.md)


## Documentation For Authorization

## apiExpires
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
## apiKey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
## apiSignature
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

support@bitmex.com

