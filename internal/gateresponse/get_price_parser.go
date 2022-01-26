package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"net/http"
	"strconv"
)

func GetCurrentPriceParser(response *http.Response) (float64, error) {
	body, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}
	bodyIArr, isOk := body.([]map[string]interface{})
	if !isOk {
		return 0, errors.New("[gateresponse] -> Error when casting body to bodyIArr in GetCurrentPriceParser")
	}
	bodyI := bodyIArr[0]

	priceStr, isOk := bodyI[pnames.Last].(string)
	if isOk != true {
		return 0, errors.New("[gateresponse] -> Error when casting priceI to priceStr in GetCurrentPriceParser")
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, errors.New("[gateresponse] -> Error when parsing priceArr to price GetCurrentPriceParser")
	}

	return price, nil
}
