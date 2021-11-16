package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"net/http"
	"strconv"
)

type MarketTrades struct {
	Id         int
	CreateTime float64
	Side       string
	Amount     float64
	Price      float64
}

func ParseMarketTrades(response *http.Response) ([]MarketTrades, error) {
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

func structureMarketTrades(body []map[string]interface{}) ([]MarketTrades, error) {
	marketTrades := make([]MarketTrades, len(body))
	var err error

	for i := 0; i < len(body); i++ {
		id, isOk := body[i][pnames.Id].(string)
		if isOk != true {
			return nil, errors.New("[gateresponse] -> error when casting market trades id")
		}
		marketTrades[i].Id, err = strconv.Atoi(id)

		createTimeStr, isOk := body[i][pnames.CreateTimeMs].(string)
		if isOk != true {
			return nil, errors.New("[gateresponse] -> error when casting market trades create_time_ms")
		}
		marketTrades[i].CreateTime, err = strconv.ParseFloat(createTimeStr, 64)
		if err != nil {
			return nil, errors.New("[gateresponse] -> error when parsing create_time_ms to float64")
		}

		marketTrades[i].Side, isOk = body[i][pnames.Side].(string)
		if isOk != true {
			return nil, errors.New("[gateresponse] -> error when casting market trades side")
		}

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
