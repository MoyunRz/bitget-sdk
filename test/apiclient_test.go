package test

import (
	"fmt"
	"github.com/MoyunRz/bitget-sdk/pkg/client"
	"github.com/MoyunRz/bitget-sdk/pkg/client/mix"
	"github.com/MoyunRz/bitget-sdk/utils"
	"testing"
)

func Test_PlaceOrder(t *testing.T) {
	client := new(mix.MixOrderClient).Init()

	params := utils.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	params["marginCoin"] = "USDT"
	params["side"] = "open_long"
	params["orderType"] = "limit"
	params["price"] = "27012"
	params["size"] = "0.01"
	params["timInForceValue"] = "normal"

	resp, err := client.PlaceOrder(params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_post(t *testing.T) {
	client := new(client.BitgetApiClient).Init()

	params := utils.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	params["marginCoin"] = "USDT"
	params["side"] = "open_long"
	params["orderType"] = "limit"
	params["price"] = "27012"
	params["size"] = "0.01"
	params["timInForceValue"] = "normal"

	resp, err := client.Post("/api/mix/v1/order/placeOrder", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get(t *testing.T) {
	client := new(client.BitgetApiClient).Init()

	params := utils.NewParams()
	params["productType"] = "umcbl"

	resp, err := client.Get("/api/mix/v1/account/accounts", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get_with_params(t *testing.T) {
	client := new(client.BitgetApiClient).Init()

	params := utils.NewParams()

	resp, err := client.Get("/api/spot/v1/account/getInfo", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get_with_encode_params(t *testing.T) {
	client := new(client.BitgetApiClient).Init()

	params := utils.NewParams()
	params["symbol"] = "$AIUSDT"
	params["businessType"] = "spot"

	resp, err := client.Get("/api/v2/common/trade-rate", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}
