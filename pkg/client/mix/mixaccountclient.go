package mix

import (
	"github.com/MoyunRz/bitget-sdk/common"
	"github.com/MoyunRz/bitget-sdk/constants"
	"github.com/MoyunRz/bitget-sdk/pkg/model/mix/account"
	"github.com/MoyunRz/bitget-sdk/utils"
)

type MixAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MixAccountClient) Init() *MixAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

/*
*
单个币种账户信息
symbol:
marginCoin:
*/
func (p *MixAccountClient) Account(symbol string, marginCoin string) (string, error) {

	params := utils.NewParams()
	params["symbol"] = symbol
	params["marginCoin"] = marginCoin

	uri := constants.MixAccount + "/account"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/*
*
根据业务线获取账户列表
productType:
*/
func (p *MixAccountClient) Accounts(productType string) (string, error) {

	params := utils.NewParams()
	params["productType"] = productType

	uri := constants.MixAccount + "/accounts"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

/*
*
调整杠杆
*/
func (p *MixAccountClient) SetLeverage(params account.SetLeveragerReq) (string, error) {
	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/set-leverage"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}

/*
*
调整保证金
*/
func (p *MixAccountClient) SetMargin(params account.SetMarginReq) (string, error) {

	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/setMargin"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}

/*
*
调节保证金模式
*/
func (p *MixAccountClient) SetMarginMode(params account.SetMarginModeReq) (string, error) {

	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/setMarginMode"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}

/*
*
计算可开张数
*/
func (p *MixAccountClient) OpenCount(params account.OpenCountReq) (string, error) {

	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/open-count"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}
