package gate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gate-api-go/internal/gaterequest"
	"gate-api-go/internal/gateresponse"
	"gate-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
)

func (manager *ExchangeManager) SetOrder(parameters order.Parameters) (float64, error) {
	bodyJson,err := manager.createOrderRequestBody(&parameters)
	if err != nil{
		return 0, err
	}
	req,err := http.NewRequest(http.MethodPost, baseUrl+prefix+newOrderEndpoint,bytes.NewBuffer(bodyJson))

	gaterequest.SetHeader(req)

	gaterequest.MakeSign(gaterequest.SignStruct{
		Method:   http.MethodPost,
		Prefix:   prefix,
		EndPoint: newOrderEndpoint,
		Body:     bodyJson,
		Api:      manager.apiKey,
	},req)

	response, err := manager.client.Do(req)
	if err != nil {
		return 0, err
	}

	defer gateresponse.CloseBody(response)
	return gateresponse.ParseSetOrder(response)

}

func (manager *ExchangeManager) createOrderRequestBody(params *order.Parameters) ([]byte, error) {
	body := make(map[string]interface{})
	body[pnames.Symbol] = fmt.Sprintf("%s_%s", params.Assets.Base,params.Assets.Quote)
	body[pnames.Type] 	= orderTypeAlias[params.Type]
	body[pnames.Side]   = orderSideAlias[params.Side]
	body[pnames.Amount] = fmt.Sprint(params.Quantity)
	body[pnames.Price]  = fmt.Sprint(params.Price)

	bodyJson,err := json.Marshal(body)
	if err != nil{
		return nil, errors.New("[gate] -> Error when marshal body to bodyJson in createOrderRequestBody")
	}

	return bodyJson,nil
}
