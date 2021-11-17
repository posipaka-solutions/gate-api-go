package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"github.com/posipaka-trade/gate-api-go/pkg/gate/trade"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
	"strconv"
	"time"
)

func ParseMarketTrades(response *http.Response) ([]trade.MarketTrades, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return nil, err
	}
	body, isOk := bodyI.([]map[string]interface{})
	if isOk != true {
		return nil, errors.New("[gateresponse] -> error when casting market trades body")
	}

	return structureMarketTrades(body)
}

func structureMarketTrades(body []map[string]interface{}) ([]trade.MarketTrades, error) {
	marketTrades := make([]trade.MarketTrades, len(body))
	var err error

	for i := 0; i < len(body); i++ {
		id, isOk := body[i][pnames.Id].(string)
		if isOk != true {
			return nil, errors.New("[gateresponse] -> error when casting market trades id")
		}
		marketTrades[i].Id, err = strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("[gateresponse] -> error when parsing id to int")
		}

		createTimeStr, isOk := body[i][pnames.CreateTime].(string)
		if isOk != true {
			return nil, errors.New("[gateresponse] -> error when casting market trades create_time_ms")
		}
		createTime, err := strconv.ParseInt(createTimeStr, 10, 64)
		if err != nil {
			return nil, errors.New("[gateresponse] -> error when parsing create_time to int64")
		}
		marketTrades[i].CreateTime = time.Unix(createTime, 0)

		marketTrades[i].Side, err = getOrderSide(body[i])

		amountStr, isOk := body[i][pnames.Amount].(string)
		if isOk != true {
			return nil, errors.New("[gateresponse] -> error when casting market trades amount")
		}
		marketTrades[i].Amount, err = strconv.ParseFloat(amountStr, 64)
		if err != nil {
			return nil, errors.New("[gateresponse] -> error when parsing amount to float64")
		}

		priceStr, isOk := body[i][pnames.Price].(string)
		if isOk != true {
			return nil, errors.New("[gateresponse] -> error when casting market trades price")
		}
		marketTrades[i].Price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, errors.New("[gateresponse] -> error when parsing price to float64")
		}
	}
	return marketTrades, nil
}

func getOrderSide(body map[string]interface{}) (order.Side, error) {
	sideStr, isOkay := body[pnames.Side].(string)
	if !isOkay {
		return order.OtherSide, errors.New("[gateresponse] -> error when casting market side to string")
	}

	switch sideStr {
	case "buy":
		return order.Buy, nil
	case "sell":
		return order.Sell, nil
	default:
		return order.OtherSide, nil
	}
}
