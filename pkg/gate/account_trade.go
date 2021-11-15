package gate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/posipaka-trade/gate-api-go/internal/gaterequest"
	"github.com/posipaka-trade/gate-api-go/internal/gateresponse"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"strconv"
	"time"

	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
)

func (manager *ExchangeManager) SetOrder(parameters order.Parameters) (order.Info, error) {
	bodyJson, err := manager.createOrderRequestBody(&parameters)
	if err != nil {
		return order.Info{}, err
	}
	req, err := http.NewRequest(http.MethodPost, baseUrl+prefix+newOrderEndpoint, bytes.NewBuffer(bodyJson))

	gaterequest.SetHeader(req)

	gaterequest.MakeSign(gaterequest.SignStruct{
		Method:   http.MethodPost,
		Prefix:   prefix,
		EndPoint: newOrderEndpoint,
		Body:     bodyJson,
		Api:      manager.apiKey,
	}, req)

	response, err := manager.client.Do(req)
	if err != nil {
		return order.Info{}, err
	}

	defer gateresponse.CloseBody(response)
	return gateresponse.ParseOrderInformation(response)

}

func (manager *ExchangeManager) createOrderRequestBody(params *order.Parameters) ([]byte, error) {
	body := make(map[string]interface{})
	body[pnames.CurrencyPair] = fmt.Sprintf("%s_%s", params.Assets.Base, params.Assets.Quote)
	body[pnames.Type] = orderTypeAlias[params.Type]
	body[pnames.Side] = orderSideAlias[params.Side]
	body[pnames.Amount] = strconv.FormatFloat(params.Quantity, 'f', -1, 64)
	body[pnames.Price] = strconv.FormatFloat(params.Price, 'f', -1, 64)
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("[gate] -> Error when marshal body to bodyJson in createOrderRequestBody")
	}

	return bodyJson, nil
}

func (manager *ExchangeManager) GetAssetBalance(asset string) (float64, error) {
	queryParam := fmt.Sprintf("%s=%s", pnames.Currency, asset)
	params := fmt.Sprintf("%s%s%s?%s", baseUrl, prefix, balancesEndpoint, queryParam)
	req, err := http.NewRequest(http.MethodGet, params, nil)
	if err != nil {
		return 0, err
	}
	gaterequest.SetHeader(req)

	gaterequest.MakeSign(gaterequest.SignStruct{
		Method:     http.MethodGet,
		Prefix:     prefix,
		EndPoint:   balancesEndpoint,
		Api:        manager.apiKey,
		QueryParam: queryParam,
	}, req)

	response, err := manager.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer gateresponse.CloseBody(response)
	return gateresponse.GetAssetBalanceParser(response)
}

func (manager *ExchangeManager) GetOrderInfo(orderId string) (order.Info, error) {
	params := fmt.Sprint(baseUrl, prefix, newOrderEndpoint, "/", orderId)
	req, err := http.NewRequest(http.MethodGet, params, nil)
	if err != nil {
		return order.OrderInfo{}, err
	}
	gaterequest.SetHeader(req)

	gaterequest.MakeSign(gaterequest.SignStruct{
		Method:   http.MethodGet,
		Prefix:   prefix,
		EndPoint: fmt.Sprint(newOrderEndpoint, "/", id),
		Api:      manager.apiKey,
	}, req)
	response, err := manager.client.Do(req)
	if err != nil {
		return order.OrderInfo{}, err
	}
	defer gateresponse.CloseBody(response)
	return gateresponse.GetOrderParser(response)
}

func (manager *ExchangeManager) StoreSymbolsLimits(limits []symbol.Limits) {
}
func (manager *ExchangeManager) GetSymbolsList() []symbol.Assets {
	return nil
}
func (manager *ExchangeManager) GetServerTime() (time.Time, error) {
	return time.Time{}, nil
}
func (manager *ExchangeManager) GetOrdersList(assets symbol.Assets) ([]order.Info, error) {
	return nil, nil
}
func (manager *ExchangeManager) GetSymbolsLimits() ([]symbol.Limits, error) {
	return nil, nil
}
