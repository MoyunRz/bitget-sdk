package mix

import (
	"encoding/json"
	"fmt"
	"github.com/MoyunRz/bitget-sdk/common"
	"github.com/MoyunRz/bitget-sdk/constants"
	"github.com/MoyunRz/bitget-sdk/pkg/model/mix/order"
	"github.com/MoyunRz/bitget-sdk/utils"
)

type MixOrderClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MixOrderClient) Init() *MixOrderClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

/*
*
下单
*/
func (p *MixOrderClient) PlaceOrder(params order.PlaceOrderReq) (order.Response, error) {

	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return order.Response{}, jsonErr
	}

	uri := constants.MixOrder + "/placeOrder"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	// 解析 JSON
	var response order.Response
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		fmt.Println("JSON 解析错误:", err)
		return response, err
	}
	return response, err

}

/*
*
批量下单
*/
func (p *MixOrderClient) BatchOrders(params order.BatchOrdersReq) (string, error) {

	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixOrder + "/batch-orders"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err

}

/*
*
撤单
*/
func (p *MixOrderClient) CancelOrder(params order.CancelOrderReq) (order.CancelResponse, error) {

	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return order.CancelResponse{}, jsonErr
	}

	uri := constants.MixOrder + "/cancel-order"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)
	var cancelResponse order.CancelResponse
	err = json.Unmarshal([]byte(resp), &cancelResponse)
	if err != nil {
		fmt.Println("JSON 解析错误:", err)
		return cancelResponse, err
	}
	return cancelResponse, err

}

/*
*
批量撤单
*/
func (p *MixOrderClient) CancelBatchOrders(params order.CancelBatchOrdersReq) (order.CancelBatchResponse, error) {

	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return order.CancelBatchResponse{}, jsonErr
	}

	uri := constants.MixOrder + "/cancel-batch-orders"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)
	// 解析 JSON
	var cancelResponse order.CancelBatchResponse
	err = json.Unmarshal([]byte(resp), &cancelResponse)
	if err != nil {
		fmt.Println("JSON 解析错误:", err)
		return cancelResponse, err
	}
	return cancelResponse, err

}

/*
*
获取订单列表的 历史委托---带分页的
*/
func (p *MixOrderClient) History(symbol string, startTime string, endTime string, pageSize string, lastEndId string, isPre string) (string, error) {
	params := utils.NewParams()
	params["symbol"] = symbol
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["pageSize"] = pageSize

	if len(lastEndId) > 0 {
		params["lastEndId"] = lastEndId
	}
	if len(isPre) > 0 {
		params["isPre"] = isPre
	}

	uri := constants.MixOrder + "/history"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/*
*
获取当前委托----不带分页的
*/
func (p *MixOrderClient) Current(symbol string) (string, error) {
	params := utils.NewParams()
	params["symbol"] = symbol

	uri := constants.MixOrder + "/current"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/*
*
获取订单信息
*/
func (p *MixOrderClient) Detail(symbol string, orderId string) (order.DetailResponse, error) {
	params := utils.NewParams()
	params["symbol"] = symbol
	params["orderId"] = orderId

	uri := constants.MixOrder + "/detail"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	// 解析 JSON
	var response order.DetailResponse
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		fmt.Println("JSON 解析错误:", err)
		return response, err
	}
	return response, err

}

/*
*
查询成交明细
*/
func (p *MixOrderClient) Fills(symbol string, orderId string) (string, error) {
	params := utils.NewParams()
	params["symbol"] = symbol
	params["orderId"] = orderId

	uri := constants.MixOrder + "/fills"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/*
*
查询成交明细
*/
func (p *MixOrderClient) ClosePositions(params order.ClosepositionsReq) (order.ClosePositionResponse, error) {
	postBody, jsonErr := utils.ToJson(params)

	if jsonErr != nil {
		return order.ClosePositionResponse{}, jsonErr
	}

	uri := constants.MixOrder + "/close-positions"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)
	// 解析 JSON
	var closePositionResponse order.ClosePositionResponse
	err = json.Unmarshal([]byte(resp), &closePositionResponse)
	if err != nil {
		fmt.Println("JSON 解析错误:", err)
		return closePositionResponse, err
	}

	return closePositionResponse, err

}
