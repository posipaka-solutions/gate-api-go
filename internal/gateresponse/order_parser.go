package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ParseOrderInformation(response *http.Response) (order.Info, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return order.Info{}, err
	}
	body, isOkay := bodyI.(map[string]interface{})
	if !isOkay {
		return order.Info{}, errors.New("[gateresponse] -> Error when casting interface{} of response to key-value pair")
	}

	orderInfo, err := retrieveOrderInfoValues(body)
	if err != nil {
		return order.Info{}, err
	}

	if orderInfo.Status == order.Filled {
		orderInfo.FilledPrice = orderInfo.QuoteQuantity / orderInfo.BaseQuantity
	}

	orderInfo.BaseQuantity = orderInfo.BaseQuantity - orderInfo.Commission

	return orderInfo, nil
}

func retrieveOrderInfoValues(body map[string]interface{}) (order.Info, error) {
	orderInfo, err := getMainOrderParameters(body)
	if err != nil {
		return order.Info{}, err
	}

	orderInfo.Assets, err = getOrderAssets(body)
	if err != nil {
		return order.Info{}, err
	}

	orderInfo.Status, err = getOrderStatus(body)
	if err != nil {
		return order.Info{}, err
	}

	orderInfo.Side, err = getOrderSide(body)
	if err != nil {
		return order.Info{}, err
	}

	transactTime, isOkay := body[pnames.UpdateTimeMs].(float64)
	if !isOkay {
		return order.Info{}, errors.New("[gateresponse] -> Field `update_time_ms` does not exist")
	}
	orderInfo.TransactionTime = time.UnixMilli(int64(transactTime))

	orderInfo.Id, isOkay = body[pnames.Id].(string)
	if !isOkay {
		return order.Info{}, errors.New("[gateresponse] -> Field `id` does not exist")
	}
	orderInfo.Type = order.Limit

	return orderInfo, nil
}

func getOrderSide(body map[string]interface{}) (order.Side, error) {
	sideStr, isOkay := body[pnames.Side].(string)
	if !isOkay {
		return order.UnknownSide, errors.New("[gateresponse] -> Field `side` does not exist")
	}

	switch sideStr {
	case "buy":
		return order.Buy, nil
	case "sell":
		return order.Sell, nil
	default:
		return order.UnknownSide, nil
	}
}

func getOrderAssets(body map[string]interface{}) (symbol.Assets, error) {
	pair, isOkay := body[pnames.CurrencyPair].(string)
	if !isOkay {
		return symbol.Assets{}, errors.New("[gateresponse] -> Field `currency_pair` does not exist")
	}

	currencies := strings.Split(pair, "_")
	if len(currencies) != 2 {
		return symbol.Assets{}, errors.New("[gateresponse] -> Failed currency pair split: " + pair)
	}

	return symbol.Assets{
		Base:  currencies[0],
		Quote: currencies[1],
	}, nil
}

func getOrderStatus(body map[string]interface{}) (order.Status, error) {
	statusStr, isOkay := body[pnames.Status].(string)
	if !isOkay {
		return order.UnknownStatus, errors.New("[gateresponse] -> Field `status does not exist`")
	}

	switch statusStr {
	case "open":
		return order.Open, nil
	case "closed":
		return order.Filled, nil
	case "cancelled":
		return order.Canceled, nil
	default:
		return order.UnknownStatus, nil
	}
}

func getMainOrderParameters(body map[string]interface{}) (order.Info, error) {
	baseAmountStr, isOkay := body[pnames.Amount].(string)
	if !isOkay {
		return order.Info{}, errors.New("[gateresponse] -> Field `amount` does not exist")
	}

	quoteAmountStr, isOkay := body[pnames.FilledTotal].(string)
	if !isOkay {
		return order.Info{}, errors.New("[gateresponse] -> Field `filled_total` does not exist")
	}

	feeStr, isOkay := body[pnames.Fee].(string) // fee in base
	if !isOkay {
		return order.Info{}, errors.New("[gateresponse] -> Field `fee` does not exist")
	}

	orderPriceStr, isOkay := body[pnames.Price].(string)
	if !isOkay {
		return order.Info{}, errors.New("[gateresponse] -> Field `price` does not exist")
	}

	var orderInfo order.Info
	var err error
	orderInfo.Price, err = strconv.ParseFloat(orderPriceStr, 64)
	if err != nil {
		return order.Info{}, errors.New("[gateresponse] -> " + err.Error())
	}

	orderInfo.Commission, err = strconv.ParseFloat(feeStr, 64)
	if err != nil {
		return order.Info{}, errors.New("[gateresponse] -> " + err.Error())
	}

	orderInfo.BaseQuantity, err = strconv.ParseFloat(baseAmountStr, 64)
	if err != nil {
		return order.Info{}, errors.New("[gateresponse] -> " + err.Error())
	}

	orderInfo.QuoteQuantity, err = strconv.ParseFloat(quoteAmountStr, 64)
	if err != nil {
		return order.Info{}, errors.New("[gateresponse] -> " + err.Error())
	}

	return orderInfo, nil
}
