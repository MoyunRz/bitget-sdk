package order

type CancelOrderReq struct {
	Symbol      string `json:"symbol"`
	MarginCoin  string `json:"marginCoin"`
	OrderId     string `json:"orderId"`
	ClientOid   string `json:"clientOid"`
	ProductType string `json:"productType"`
}

// CancelData 结构体表示嵌套的 "data" 字段
type CancelData struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
}

// CancelResponse 结构体表示整个 JSON 响应
type CancelResponse struct {
	Code        string     `json:"code"`
	Data        CancelData `json:"data"`
	Msg         string     `json:"msg"`
	RequestTime int64      `json:"requestTime"`
}

// Failure 结构体表示失败的订单
type Failure struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
	ErrorMsg  string `json:"errorMsg"`
}

// CancelBatchData 结构体表示嵌套的 "data" 字段
type CancelBatchData struct {
	SuccessList []CancelData `json:"successList"`
	FailureList []Failure    `json:"failureList"`
}

// CancelBatchResponse 结构体表示整个 JSON 响应
type CancelBatchResponse struct {
	Code        string          `json:"code"`
	Data        CancelBatchData `json:"data"`
	Msg         string          `json:"msg"`
	RequestTime int64           `json:"requestTime"`
}
