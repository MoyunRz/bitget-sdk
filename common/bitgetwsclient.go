package common

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/MoyunRz/bitget-sdk/config"
	"github.com/MoyunRz/bitget-sdk/constants"
	model2 "github.com/MoyunRz/bitget-sdk/model"
	"github.com/MoyunRz/bitget-sdk/pkg/safe"
	"github.com/MoyunRz/bitget-sdk/utils"
	"github.com/kurosann/aqt-core/logger"
	"github.com/kurosann/aqt-core/ws"
	"go.uber.org/zap"
)

var _ ws.MsgHandler = (*BitgetBaseWsClient)(nil)

// BitgetBaseWsClient WebSocket 客户端结构体
type BitgetBaseWsClient struct {
	keepAlive   *ws.KeepAliver
	needLogin   bool
	loginStatus bool
	rw          sync.RWMutex
	signer      *Signer
	scribeMap   map[model2.SubscribeReq]OnReceive
}

type OnReceive func(message string)

func NewBitgetClient(dialer ws.Dialer, needLogin bool) *BitgetBaseWsClient {
	p := &BitgetBaseWsClient{
		scribeMap: make(map[model2.SubscribeReq]OnReceive),
		needLogin: needLogin,
		signer:    new(Signer).Init(config.SecretKey),
	}

	wsURL := config.PublicWsUrl
	if needLogin {
		wsURL = config.PrivateWsUrl
	}

	keepAlive := &ws.KeepAliver{
		Address:      wsURL,
		Dialer:       dialer,
		Delay:        15 * time.Second,
		MsgHandler:   p,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	p.keepAlive = keepAlive
	return p
}

// Connect 建立连接
func (p *BitgetBaseWsClient) Connect() {
	logger.Debug("connecting...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ctx, err := p.keepAlive.KeepAlive(ctx)
	if err != nil {
		logger.Error("keepalive error", zap.Error(err))
		return
	}
	if p.needLogin {
		logger.Debug("loging...")
		p.login()
		logger.Debug("login success")
	}
	logger.Debug("keep alive...")
	safe.Go(func() {
		if err := p.executePing(ctx); err != nil {
			logger.Error("ping error", zap.Error(err))
		}
	})
}

// executePing 执行 ping
func (p *BitgetBaseWsClient) executePing(ctx context.Context) error {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			ping := &model2.WsBaseReq{
				Op: "ping",
			}
			return p.SendReq(*ping)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// SendByType 发送特定类型的消息
func (p *BitgetBaseWsClient) SendReq(req model2.WsBaseReq) error {
	msgByte, err := utils.ToJson(req)
	if err != nil {
		return err
	}
	return p.keepAlive.Send([]byte(msgByte))
}

// getListener 获取消息监听器
func (p *BitgetBaseWsClient) getListener(argJson interface{}) OnReceive {
	if argJson == nil {
		return nil
	}

	mapData, ok := argJson.(map[string]interface{})
	if !ok {
		return nil
	}

	instType := fmt.Sprintf("%s", mapData["instType"])
	channel := fmt.Sprintf("%s", mapData["channel"])
	instId := fmt.Sprintf("%s", mapData["instId"])

	p.rw.RLock()
	defer p.rw.RUnlock()

	for req, receive := range p.scribeMap {
		if req.InstType == instType && req.Channel == channel && req.InstId == instId {
			return receive
		}
	}

	return nil
}

// Close 关闭连接
func (p *BitgetBaseWsClient) Close() {
	if p.keepAlive != nil {
		p.keepAlive.Close()
		p.keepAlive = nil
	}
}

func (p *BitgetBaseWsClient) login() {
	timesStamp := utils.TimesStampSec()
	sign := p.signer.Sign(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)

	loginReq := model2.WsLoginReq{
		ApiKey:     config.ApiKey,
		Passphrase: config.PASSPHRASE,
		Timestamp:  timesStamp,
		Sign:       sign,
	}

	wsBaseReq := model2.WsBaseReq{
		Op:   constants.WsOpLogin,
		Args: []interface{}{loginReq},
	}
	if err := p.SendReq(wsBaseReq); err != nil {
		logger.Error("login error", zap.Error(err))
	}
}

func (p *BitgetBaseWsClient) OnError(err error) {
	logger.Error("error", zap.Error(err))
}

func (p *BitgetBaseWsClient) OnReceive(msg []byte) error {
	message := string(msg)
	if message == "" || message == "pong" {
		return nil
	}

	jsonMap := utils.JSONToMap(message)

	if code, exists := jsonMap["code"]; exists && int(code.(float64)) != 0 {
		logger.Error("error", zap.String("message", message))
		return nil
	}

	if event, exists := jsonMap["event"]; exists && event == "login" {
		p.loginStatus = true
		return nil
	}

	if _, exists := jsonMap["data"]; exists {
		if listener := p.getListener(jsonMap["arg"]); listener != nil {
			listener(message)
			return nil
		}
	}

	return nil
}

func (p *BitgetBaseWsClient) OnReconnect() (msg [][]byte) {
	msg = make([][]byte, 0)
	for req := range p.scribeMap {
		wsBaseReq := model2.WsBaseReq{
			Op:   constants.WsOpSubscribe,
			Args: []interface{}{req},
		}
		if json, err := utils.ToJson(wsBaseReq); err == nil {
			msg = append(msg, []byte(json))
		}
	}
	return msg
}

func (p *BitgetBaseWsClient) Listen(req model2.SubscribeReq, listener OnReceive) {
	p.rw.Lock()
	defer p.rw.Unlock()

	if _, ok := p.scribeMap[req]; !ok {
		p.scribeMap[req] = listener
		wsBaseReq := model2.WsBaseReq{
			Op:   constants.WsOpSubscribe,
			Args: []interface{}{req},
		}
		if err := p.SendReq(wsBaseReq); err != nil {
			p.OnError(err)
		}
	}
}

func (p *BitgetBaseWsClient) UnListen(req model2.SubscribeReq) {
	p.rw.Lock()
	defer p.rw.Unlock()

	if _, ok := p.scribeMap[req]; ok {
		delete(p.scribeMap, req)
		wsBaseReq := model2.WsBaseReq{
			Op:   constants.WsOpUnsubscribe,
			Args: []interface{}{req},
		}
		if err := p.SendReq(wsBaseReq); err != nil {
			p.OnError(err)
		}
	}
}
