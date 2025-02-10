package spot

import (
	"github.com/MoyunRz/bitget-sdk/common"
	"github.com/MoyunRz/bitget-sdk/utils"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init() *SpotAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *SpotAccountClient) Info() (string, error) {
	params := utils.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/info", params)
	return resp, err
}

func (p *SpotAccountClient) Assets(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/assets", params)
	return resp, err
}

func (p *SpotAccountClient) Bills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/bills", params)
	return resp, err
}

func (p *SpotAccountClient) TransferRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/transferRecords", params)
	return resp, err
}
