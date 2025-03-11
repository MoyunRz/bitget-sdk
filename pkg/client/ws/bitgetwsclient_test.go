package ws

import (
	"fmt"
	"testing"
	"time"

	"github.com/MoyunRz/bitget-sdk/model"
	"github.com/kurosann/aqt-core/logger"
	"github.com/stretchr/testify/assert"
)

func TestBitgetWsClient_New(t *testing.T) {
	t.Setenv("http_proxy", "http://127.0.0.1:7890")
	t.Setenv("https_proxy", "http://127.0.0.1:7890")
	t.Setenv("all_proxy", "socks5://127.0.0.1:7890")
	client := NewBitgetWsClient(false)

	var channelsDef []model.SubscribeReq
	subReqDef1 := model.SubscribeReq{
		InstType: "USDT-FUTURES",
		Channel:  "books1",
		InstId:   "ETHUSDT",
	}
	channelsDef = append(channelsDef, subReqDef1)

	testChannel := make(chan string)
	client.Subscribe(channelsDef, func(message string) {
		testChannel <- message
	})

	var channels []model.SubscribeReq
	subReq1 := model.SubscribeReq{
		InstType: "USDT-FUTURES",
		Channel:  "books1",
		InstId:   "ETHUSDT",
	}
	channels = append(channels, subReq1)
	client.Subscribe(channels, func(message string) {
		testChannel <- message
	})
	logger.Debug("subscribe done")
	go func() {
		time.Sleep(3 * time.Second)
		close(testChannel)
	}()

	mapTest := make(map[string]bool)

	for message := range testChannel {
		if _, ok := mapTest[message]; !ok {
			mapTest[message] = true
		}
		fmt.Println("appoint:" + message)
	}
	assert.GreaterOrEqual(t, len(mapTest), 2)
}
