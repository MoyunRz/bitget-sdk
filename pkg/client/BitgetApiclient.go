package client

import (
	"github.com/MoyunRz/bitget-sdk/common"
	"github.com/MoyunRz/bitget-sdk/utils"
)

type BitgetApiClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *BitgetApiClient) Init() *BitgetApiClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *BitgetApiClient) Post(url string, params map[string]string) (string, error) {
	postBody, jsonErr := utils.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost(url, postBody)
	return resp, err
}

func (p *BitgetApiClient) Get(url string, params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet(url, params)
	return resp, err
}
