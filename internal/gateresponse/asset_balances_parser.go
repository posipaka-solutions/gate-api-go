package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"net/http"
	"strconv"
)

func GetAssetBalanceParser(response *http.Response) (float64, error) {
	body, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}
	bodyIArr, isOk := body.([]map[string]interface{})
	if !isOk {
		return 0, errors.New("[gateresponse] -> Error when casting body to bodyIArr in GetAssetBalanceParser")
	}
	available, isOk := bodyIArr[0][pnames.Available].(string)
	if !isOk {
		return 0, errors.New("[gateresponse] -> Error when casting bodyI to available in GGetAssetBalanceParser")
	}
	balance, err := strconv.ParseFloat(available, 64)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
