package config

var (
	//restApi config
	BaseUrl       = "https://api.bitget.com"
	ApiKey        = "bg_b2f572a5c924ad2ad59e62710914d9a8"
	SecretKey     = "aa3d8bc37f3cee4f30f798813da2e5fa07c2b3db24d5691653add36513aeaa2a"
	PASSPHRASE    = "Lanan666"
	TimeoutSecond = 30

	//websocket config
	PublicWsUrl  = "wss://ws.bitget.com/v2/ws/public"
	PrivateWsUrl = "wss://ws.bitget.com/v2/ws/private"
)

func InitConfig(apiKey, secretKey, passphrase string) {
	ApiKey = apiKey
	SecretKey = secretKey
	PASSPHRASE = passphrase
}
