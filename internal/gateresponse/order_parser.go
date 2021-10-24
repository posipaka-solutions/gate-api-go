package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
	"strconv"
)

func ParseSetOrder(response *http.Response) (order.OrderInfo, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return order.OrderInfo{}, err
	}
	body, isOkay := bodyI.(map[string]interface{})
	if !isOkay {
		return order.OrderInfo{}, errors.New("[gateresponse] -> Error when casting bodyI to body in ParseSetOrder")
	}
	var orderInfo order.OrderInfo
	cryptoAmountStr, isOkay := body[pnames.Amount].(string)
	if isOkay != true {
		return order.OrderInfo{}, errors.New("[gateresponse] -> Error when parsing amount to string")
	}
	orderInfo.Quantity, err = strconv.ParseFloat(cryptoAmountStr, 64)
	if err != nil {
		return order.OrderInfo{}, err
	}
	usdtAmountStr, isOkay := body[pnames.FilledTotal].(string)
	if isOkay != true {
		return order.OrderInfo{}, errors.New("[gateresponse] -> Error when parsing filled_total to string")
	}
	usdtAmount, err := strconv.ParseFloat(usdtAmountStr, 64)
	orderInfo.Price = usdtAmount / orderInfo.Quantity

	return orderInfo, nil
}
