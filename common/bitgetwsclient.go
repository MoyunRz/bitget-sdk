package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/MoyunRz/bitget-sdk/config"
	"github.com/MoyunRz/bitget-sdk/constants"
	"github.com/MoyunRz/bitget-sdk/logging/applogger"
	model2 "github.com/MoyunRz/bitget-sdk/model"
	"github.com/MoyunRz/bitget-sdk/pkg/safe"
	"github.com/MoyunRz/bitget-sdk/utils"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type BitgetBaseWsClient struct {
	ctx              context.Context
	NeedLogin        bool
	Connection       bool
	LoginStatus      bool
	Listener         OnReceive
	ErrorListener    OnReceive
	Ticker           *time.Ticker
	SendMutex        *sync.Mutex
	WebSocketClient  *websocket.Conn
	LastReceivedTime time.Time
	AllSuribe        *model2.Set
	Signer           *Signer
	ScribeMap        map[model2.SubscribeReq]OnReceive
}

func (p *BitgetBaseWsClient) Init() *BitgetBaseWsClient {
	p.ctx = context.Background()
	p.Connection = false
	p.AllSuribe = model2.NewSet()
	p.Signer = new(Signer).Init(config.SecretKey)
	p.ScribeMap = make(map[model2.SubscribeReq]OnReceive)
	p.SendMutex = &sync.Mutex{}
	p.Ticker = time.NewTicker(constants.TimerIntervalSecond * time.Second)
	p.LastReceivedTime = time.Now()

	return p
}

func (p *BitgetBaseWsClient) SetListener(msgListener OnReceive, errorListener OnReceive) {
	p.Listener = msgListener
	p.ErrorListener = errorListener
}

func (p *BitgetBaseWsClient) Connect() {
	safe.Go(func() { p.tickerLoop() })
}

func (p *BitgetBaseWsClient) ConnectWebSocket() {
	var err error
	applogger.Info("WebSocket connecting...")
	if p.NeedLogin {
		p.WebSocketClient, _, err = websocket.DefaultDialer.Dial(config.PrivateWsUrl, nil)
		if err != nil {
			fmt.Printf("WebSocket connected error: %s\n", err)
			return
		}
		applogger.Info("WebSocket connected")
	} else {
		p.WebSocketClient, _, err = websocket.DefaultDialer.Dial(config.PublicWsUrl, nil)
		if err != nil {
			fmt.Printf("WebSocket connected error: %s\n", err)
			return
		}
		applogger.Info("WebSocket connected")
	}
	p.Connection = true
	p.LastReceivedTime = time.Now()

	ctx, cancel := context.WithCancel(p.ctx)
	defer cancel()
	safe.Go(func() {
		p.ExecuterPing(ctx)
		p.ReadLoop()
	})

}

func (p *BitgetBaseWsClient) Login() {
	timesStamp := utils.TimesStampSec()
	sign := p.Signer.Sign(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)

	loginReq := model2.WsLoginReq{
		ApiKey:     config.ApiKey,
		Passphrase: config.PASSPHRASE,
		Timestamp:  timesStamp,
		Sign:       sign,
	}
	var args []interface{}
	args = append(args, loginReq)

	baseReq := model2.WsBaseReq{
		Op:   constants.WsOpLogin,
		Args: args,
	}
	p.SendByType(baseReq)
}

func (p *BitgetBaseWsClient) StartReadLoop() {
	safe.Go(func() {
		p.ReadLoop()
	})
}

func (p *BitgetBaseWsClient) ExecuterPing(ctx context.Context) error {

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := p.ping()
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}

}
func (p *BitgetBaseWsClient) ping() error {
	return p.Send("ping")
}

func (p *BitgetBaseWsClient) SendByType(req model2.WsBaseReq) {
	json, _ := utils.ToJson(req)
	p.Send(json)
}

func (p *BitgetBaseWsClient) Send(data string) error {
	if p.WebSocketClient == nil {
		applogger.Error("WebSocket sent error: no connection available")
		return errors.New("WebSocket sent error: no connection available")
	}
	applogger.Info("sendMessage:%s", data)
	p.SendMutex.Lock()
	err := p.WebSocketClient.WriteMessage(websocket.TextMessage, []byte(data))
	p.SendMutex.Unlock()
	if err != nil {
		applogger.Error("WebSocket sent error: data=%s, error=%s", data, err)
		return err
	}
	return nil
}

func (p *BitgetBaseWsClient) tickerLoop() {
	applogger.Info("tickerLoop started")
	for {
		select {
		case <-p.Ticker.C:
			elapsedSecond := time.Now().Sub(p.LastReceivedTime).Seconds()
			if elapsedSecond > constants.ReconnectWaitSecond {
				applogger.Info("WebSocket reconnect...")
				p.disconnectWebSocket()
				p.ConnectWebSocket()
				p.ReConnectSend()
			}
		}
	}
}

// ReConnectSend 重连后重新订阅
func (p *BitgetBaseWsClient) ReConnectSend() {
	var args []interface{}
	for req, _ := range p.ScribeMap {
		args = append(args, req)
	}
	wsBaseReq := model2.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}
	p.SendByType(wsBaseReq)
}

func (p *BitgetBaseWsClient) disconnectWebSocket() {
	if p.WebSocketClient == nil {
		return
	}

	fmt.Println("WebSocket disconnecting...")
	err := p.WebSocketClient.Close()
	if err != nil {
		applogger.Error("WebSocket disconnect error: %s\n", err)
		return
	}

	applogger.Info("WebSocket disconnected")
}

func (p *BitgetBaseWsClient) ReadLoop() {
	for {
		if p.WebSocketClient == nil {
			applogger.Info("Read error: no connection available")
			time.Sleep(3 * time.Second)
			continue
		}
		var message string
		_, buf, err := p.WebSocketClient.ReadMessage()
		if err != nil {
			applogger.Info("Read error: %s", err)
			return
		}
		message = string(buf)
		if message == "" {
			continue
		}
		p.LastReceivedTime = time.Now()
		if message == "pong" {
			applogger.Info("Keep connected:" + message)
			continue
		}
		jsonMap := utils.JSONToMap(message)

		v, e := jsonMap["code"]

		if e && int(v.(float64)) != 0 {
			p.ErrorListener(message)
			continue
		}

		v, e = jsonMap["event"]
		if e && v == "login" {
			applogger.Info("login msg:" + message)
			p.LoginStatus = true
			continue
		}

		v, e = jsonMap["data"]
		if e {
			listener := p.GetListener(jsonMap["arg"])
			listener(message)
			continue
		}
		applogger.Info("Received msg:" + message)
		//p.handleMessage(message)
	}

}

func (p *BitgetBaseWsClient) GetListener(argJson interface{}) OnReceive {
	mapData := argJson.(map[string]interface{})
	instType := fmt.Sprintf("%s", mapData["instType"])
	channel := fmt.Sprintf("%s", mapData["channel"])
	instId := fmt.Sprintf("%s", mapData["instId"])
	for req, receive := range p.ScribeMap {
		if req.InstType == instType && req.Channel == channel && req.InstId == instId {
			return receive
		}
	}
	return p.Listener
}

type OnReceive func(message string)

func (p *BitgetBaseWsClient) handleMessage(msg string) {
	fmt.Println("default:" + msg)
}
