package mix

import (
	"github.com/MoyunRz/bitget-sdk/common"
	"github.com/MoyunRz/bitget-sdk/constants"
	"github.com/MoyunRz/bitget-sdk/utils"
)

type MixPositionClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MixPositionClient) Init() *MixPositionClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

/*
*
获取单个合约仓位信息
*/
func (p *MixPositionClient) SinglePosition(symbol string, marginCoin string) (string, error) {

	params := utils.NewParams()
	params["symbol"] = symbol
	params["marginCoin"] = marginCoin

	uri := constants.MixPosition + "/singlePosition"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/*
*
获取全部合约仓位信息
*/
func (p *MixPositionClient) AllPosition(productType string, marginCoin string) (string, error) {

	params := utils.NewParams()
	params["productType"] = productType
	params["marginCoin"] = marginCoin

	uri := constants.MixPosition + "/allPosition"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}
