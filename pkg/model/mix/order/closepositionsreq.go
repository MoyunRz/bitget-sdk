package order

type ClosepositionsReq struct {
	Symbol      string `json:"symbol"`
	HoldSide    string `json:"holdSide"`
	ProductType string `json:"productType"`
}

// ClosePositionSuccess 结构体表示成功的订单
type ClosePositionSuccess struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
	Symbol    string `json:"symbol"`
}

// ClosePositionFailure 结构体表示失败的订单
type ClosePositionFailure struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
	Symbol    string `json:"symbol"`
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}

// ClosePositionData 结构体表示嵌套的 "data" 字段
type ClosePositionData struct {
	SuccessList []ClosePositionSuccess `json:"successList"`
	FailureList []ClosePositionFailure `json:"failureList"`
}

// ClosePositionResponse 结构体表示整个 JSON 响应
type ClosePositionResponse struct {
	Code        string            `json:"code"`
	Data        ClosePositionData `json:"data"`
	Msg         string            `json:"msg"`
	RequestTime int64             `json:"requestTime"`
}
