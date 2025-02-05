package test

import (
	"fmt"
	"github.com/MoyunRz/bitget-sdk/model"
	"github.com/MoyunRz/bitget-sdk/pkg/client/ws"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	client := new(ws.BitgetWsClient).Init(false, func(message string) {
		fmt.Println("default error:" + message)
	}, func(message string) {
		fmt.Println("default error:" + message)
	})

	//"instType": "USDT-FUTURES",
	//	"channel": "ticker",
	//	"instId": "BTCUSDT"
	//var channelsDef []model.SubscribeReq
	//subReqDef1 := model.SubscribeReq{
	//	InstType: "USDT-FUTURES",
	//	Channel:  "ticker",
	//	InstId:   "ETHUSDT",
	//}
	//channelsDef = append(channelsDef, subReqDef1)
	//client.SubscribeDef(channelsDef)

	var channels []model.SubscribeReq
	subReq1 := model.SubscribeReq{
		InstType: "USDT-FUTURES",
		Channel:  "ticker",
		InstId:   "ETHUSDT",
	}
	channels = append(channels, subReq1)
	client.Subscribe(channels, func(message string) {
		fmt.Println("appoint:" + message)
	})
	fmt.Println("Press ENTER to unsubscribe and stop...")
	time.Sleep(time.Minute * 10)
}
