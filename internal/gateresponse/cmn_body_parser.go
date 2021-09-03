package gateresponse

import (
	"encoding/json"
	"errors"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"io/ioutil"
	"net/http"
)

const (
	label   = "label"
	message = "message"
)

func getResponseBody(response *http.Response) (interface{}, error) {
	if response.StatusCode/100 != 2 && response.Body == nil {
		return nil, &exchangeapi.ExchangeError{
			Type:    exchangeapi.HttpErr,
			Code:    response.StatusCode,
			Message: response.Status,
		}
	}

	respondBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if respondBody[0] == '[' && respondBody[len(respondBody)-1] == ']' && respondBody[1] != '[' {
		var body []map[string]interface{}
		err = json.Unmarshal(respondBody, &body)
		if err != nil {
			return nil, err
		}

		return body, err
	}

	var body map[string]interface{}
	err = json.Unmarshal(respondBody, &body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode/100 != 2 {
		return nil, parseGateError(body, response)
	}

	return body, nil

}

func parseGateError(body map[string]interface{}, response *http.Response) error {

	message, isOkay := body[message].(string)
	if !isOkay {
		return errors.New("[gateresponse] -> failed to parse binance error message")
	}

	return &exchangeapi.ExchangeError{
		Type:    exchangeapi.GateErr,
		Code:    response.StatusCode,
		Message: message,
	}
}
