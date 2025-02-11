package config

var (
	//restApi config
	BaseUrl       = "https://api.bitget.com"
	ApiKey        = "bg_912e9f7e29f77def9f76613852e19898"
	SecretKey     = "d4df54759dfe9d6aedd9ea4774cf9b95b2b0593ab19cc82d588c64478e8c4314"
	PASSPHRASE    = "Lanan666"
	TimeoutSecond = 30

	//websocket config
	PublicWsUrl  = "wss://ws.bitget.com/v2/ws/public"
	PrivateWsUrl = "wss://ws.bitget.com/v2/ws/private"
	IsTestNet    = false
)

func InitConfig(apiKey, secretKey, passphrase string, isTestNet bool) {
	ApiKey = apiKey
	SecretKey = secretKey
	PASSPHRASE = passphrase
	IsTestNet = isTestNet
}
