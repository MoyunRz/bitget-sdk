package constants

const (
	/*
	  http headers
	*/
	ContentType        = "Content-Type"
	BgAccessKey        = "ACCESS-KEY"
	BgAccessSign       = "ACCESS-SIGN"
	BgAccessTimestamp  = "ACCESS-TIMESTAMP"
	BgAccessPassphrase = "ACCESS-PASSPHRASE"
	ApplicationJson    = "application/json"

	EN_US  = "en_US"
	ZH_CN  = "zh_CN"
	LOCALE = "locale="

	/*
	  http methods
	*/
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"

	/*
	 others
	*/
	ResultDataJsonString = "resultDataJsonString"
	ResultPageJsonString = "resultPageJsonString"

	/**
	 * SPOT URL
	 */
	SpotPublic  = "/api/v2/spot/public"
	SpotMarket  = "/api/v2/spot/market"
	SpotAccount = "/api/v2/spot/account"
	SpotTrade   = "/api/v2/spot/trade"

	/**
	 * MIX URL
	 */
	MixPlan     = "/api/v2/mix/plan"
	MixMarket   = "/api/v2/mix/market"
	MixAccount  = "/api/v2/mix/account"
	MixOrder    = "/api/v2/mix/order"
	MixPosition = "/api/v2/mix/position"
	MixTrace    = "/api/v2/mix/trace"

	/**
	websocket
	*/
	WsAuthMethod        = "GET"
	WsAuthPath          = "/user/verify"
	WsOpLogin           = "login"
	WsOpUnsubscribe     = "unsubscribe"
	WsOpSubscribe       = "subscribe"
	TimerIntervalSecond = 5
	ReconnectWaitSecond = 60
)
