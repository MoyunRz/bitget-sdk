package common

import (
	"github.com/MoyunRz/bitget-sdk/utils"
	"testing"
)

func TestBitgetRestClient_HttpExecuter(t *testing.T) {
	restClient := new(BitgetRestClient).Init()
	params := utils.NewParams()
	params["productType"] = "umcbl"
	resp, _ := restClient.DoGet("/api/mix/v1/account/accounts", params)

	println(resp)
}
