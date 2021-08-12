package gate

import (
	"errors"
	"fmt"
	"gate-api-go/internal/gaterequest"
	"gate-api-go/internal/gateresponse"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
)

func (manager *ExchangeManager) GetCurrentPrice(symbol symbol.Assets) (float64, error) {
	params := fmt.Sprintf("currency_pair=%s_%s", symbol.Base, symbol.Quote)

	req, err := http.NewRequest(http.MethodGet,fmt.Sprint(baseUrl,prefix, getPriceEndpoint,"?", params),nil)
	if err != nil {
		return 0, errors.New("[gate] -> Error in GetRequest when getting current price")
	}
	gaterequest.SetHeader(req)

	response, err := manager.client.Do(req)
	if err != nil {
		return 0, err
	}

	defer gateresponse.CloseBody(response)

	return gateresponse.GetCurrentPriceParser(response)
}
