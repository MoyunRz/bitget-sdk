package order

type PlaceOrderReq struct {
	Symbol                string `json:"symbol"`
	MarginCoin            string `json:"marginCoin"`
	MarginMode            string `json:"marginMode"`
	TradeSide             string `json:"tradeSide"`
	ProductType           string `json:"productType"`
	Side                  string `json:"side"`
	ClientOid             string `json:"clientOid"`
	Size                  string `json:"size"`
	OrderType             string `json:"orderType"`
	Price                 string `json:"price"`
	TimeInForceValue      string `json:"timeInForceValue"`
	PresetTakeProfitPrice string `json:"presetTakeProfitPrice"`
	PresetStopLossPrice   string `json:"presetStopLossPrice"`
}

// Data 结构体表示嵌套的 "data" 字段
type Data struct {
	ClientOid string `json:"clientOid"`
	OrderId   string `json:"orderId"`
}

// Response 结构体表示整个 JSON 响应
type Response struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
	Data        Data   `json:"data"`
}

// DetailData 结构体表示嵌套的 "data" 字段
type DetailData struct {
	Symbol                 string `json:"symbol"`
	Size                   string `json:"size"`
	OrderId                string `json:"orderId"`
	ClientOid              string `json:"clientOid"`
	BaseVolume             string `json:"baseVolume"`
	PriceAvg               string `json:"priceAvg"`
	Fee                    string `json:"fee"`
	Price                  string `json:"price"`
	State                  string `json:"state"`
	Side                   string `json:"side"`
	Force                  string `json:"force"`
	TotalProfits           string `json:"totalProfits"`
	PosSide                string `json:"posSide"`
	MarginCoin             string `json:"marginCoin"`
	PresetStopSurplusPrice string `json:"presetStopSurplusPrice"`
	PresetStopLossPrice    string `json:"presetStopLossPrice"`
	QuoteVolume            string `json:"quoteVolume"`
	OrderType              string `json:"orderType"`
	Leverage               string `json:"leverage"`
	MarginMode             string `json:"marginMode"`
	ReduceOnly             string `json:"reduceOnly"`
	EnterPointSource       string `json:"enterPointSource"`
	TradeSide              string `json:"tradeSide"`
	PosMode                string `json:"posMode"`
	OrderSource            string `json:"orderSource"`
	CancelReason           string `json:"cancelReason"`
	CTime                  string `json:"cTime"`
	UTime                  string `json:"uTime"`
}

// DetailResponse 结构体表示整个 JSON 响应
type DetailResponse struct {
	Code        string     `json:"code"`
	Msg         string     `json:"msg"`
	RequestTime int64      `json:"requestTime"`
	Data        DetailData `json:"data"`
}
