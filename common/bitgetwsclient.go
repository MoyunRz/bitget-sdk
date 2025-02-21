package common

import (
	"fmt"
	"github.com/MoyunRz/bitget-sdk/config"
	"github.com/MoyunRz/bitget-sdk/constants"
	"github.com/MoyunRz/bitget-sdk/logging/applogger"
	model2 "github.com/MoyunRz/bitget-sdk/model"
	"github.com/MoyunRz/bitget-sdk/pkg/safe"
	"github.com/MoyunRz/bitget-sdk/utils"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron"
	"runtime/debug"
	"sync"
	"time"
)

type BitgetBaseWsClient struct {
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

	p.tickerLoop()
	p.ExecuterPing()
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
	p.ReadLoop()
}

func (p *BitgetBaseWsClient) ExecuterPing() {
	c := cron.New()
	_ = c.AddFunc("*/15 * * * * *", p.ping)
	c.Start()
}
func (p *BitgetBaseWsClient) ping() {
	p.Send("ping")
}

func (p *BitgetBaseWsClient) SendByType(req model2.WsBaseReq) {
	json, _ := utils.ToJson(req)
	p.Send(json)
}

func (p *BitgetBaseWsClient) Send(data string) {
	if p.WebSocketClient == nil {
		applogger.Error("WebSocket sent error: no connection available")
		return
	}
	applogger.Info("sendMessage:%s", data)
	p.SendMutex.Lock()
	err := p.WebSocketClient.WriteMessage(websocket.TextMessage, []byte(data))
	p.SendMutex.Unlock()
	if err != nil {
		applogger.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
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
			}
		}
	}
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
			time.Sleep(6 * time.Second)
			continue
		}
		wg := sync.WaitGroup{}
		var message string
		safe.Go(func() {
			wg.Add(1)
			defer func() {
				// Panic 恢复处理
				if r := recover(); r != nil {
					// 记录 panic 详细信息
					applogger.Error("Panic in WebSocket read goroutine: %v", r)
					// 打印完整的堆栈信息
					debug.PrintStack()
					// 可选：发送错误通知
					if p.ErrorListener != nil {
						p.ErrorListener(fmt.Sprintf("WebSocket read panic: %v", r))
					}
				}
				// 确保 WaitGroup 计数器被正确减少
				wg.Done()
			}()
			_, buf, err := p.WebSocketClient.ReadMessage()
			if err != nil {
				applogger.Info("Read error: %s", err)
				return
			}
			message = string(buf)
		})
		wg.Wait()
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
		p.handleMessage(message)
	}

}

func (p *BitgetBaseWsClient) GetListener(argJson interface{}) OnReceive {

	mapData := argJson.(map[string]interface{})

	subscribeReq := model2.SubscribeReq{
		InstType: fmt.Sprintf("%v", mapData["instType"]),
		Channel:  fmt.Sprintf("%v", mapData["channel"]),
		InstId:   fmt.Sprintf("%v", mapData["instId"]),
	}

	v, e := p.ScribeMap[subscribeReq]

	if !e {
		return p.Listener
	}
	return v
}

type OnReceive func(message string)

func (p *BitgetBaseWsClient) handleMessage(msg string) {
	fmt.Println("default:" + msg)
}
