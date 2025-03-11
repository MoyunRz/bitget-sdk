package ws

import (
	"crypto/tls"
	"strings"
	"sync"

	"github.com/MoyunRz/bitget-sdk/common"
	model2 "github.com/MoyunRz/bitget-sdk/model"
	"github.com/kurosann/aqt-core/ws"
)

type BitgetWsClient struct {
	bitgetBaseWsClient *common.BitgetBaseWsClient
	NeedLogin          bool
	once               sync.Once
}

func NewBitgetWsClient(needLogin bool) *BitgetWsClient {
	p := &BitgetWsClient{
		bitgetBaseWsClient: common.NewBitgetClient(ws.Dialer{
			Tls: &tls.Config{
				ServerName: "ws.bitget.com",
			},
		}, needLogin),
		NeedLogin: needLogin,
	}
	return p
}

func (p *BitgetWsClient) init() {
	p.bitgetBaseWsClient.Connect()
}

func (p *BitgetWsClient) UnSubscribe(list []model2.SubscribeReq) {
	p.once.Do(p.init)
	for _, req := range list {
		req = toUpperReq(req)
		p.bitgetBaseWsClient.UnListen(req)
	}
}

func toUpperReq(req model2.SubscribeReq) model2.SubscribeReq {
	req.InstType = strings.ToUpper(req.InstType)
	req.InstId = strings.ToUpper(req.InstId)
	req.Channel = strings.ToLower(req.Channel)
	return req
}

func (p *BitgetWsClient) Subscribe(list []model2.SubscribeReq, listener common.OnReceive) {
	p.once.Do(p.init)
	for _, req := range list {
		req = toUpperReq(req)
		p.bitgetBaseWsClient.Listen(req, listener)
	}
}
