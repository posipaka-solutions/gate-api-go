package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"net/http"
)

func ParseSetOrder(response *http.Response) (float64, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}
	body, isOk := bodyI.(map[string]interface{})
	if !isOk {
		return 0, errors.New("[gateresponse] -> Error when casting bodyI to body in ParseSetOrder")
	}
	status, isOk := body[pnames.Status].(string)
	if !isOk {
		return 0, errors.New("[gateresponse] -> Error when casting body to status in ParseSetOrder")
	}
	if status == pnames.Open {
		return 1, nil
	}
	return 0, errors.New("[gateresponse] -> Error when placing order in ParseSetOrder")
}
