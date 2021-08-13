package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"net/http"
	"strconv"
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
	idStr, isOk := body[pnames.Id].(string)
	if !isOk {
		return 0, errors.New("[gateresponse] -> Error when casting body to id in ParseSetOrder")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("[gateresponse] -> Error when parsing idStr to id in ParseSetOrder")
	}
	if status == pnames.Open {
		return 1, nil
	} else if id != 0 {
		return 1, nil
	}

	return 0, errors.New("[gateresponse] -> Error when placing order in ParseSetOrder")
}
