package gate

import (
	"errors"
	"fmt"
	"github.com/posipaka-trade/gate-api-go/internal/gaterequest"
	"github.com/posipaka-trade/gate-api-go/internal/gateresponse"
	"github.com/posipaka-trade/gate-api-go/pkg/gate/trade"

	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
)

func (manager *ExchangeManager) GetCurrentPrice(symbol symbol.Assets) (float64, error) {
	params := fmt.Sprintf("currency_pair=%s_%s", symbol.Base, symbol.Quote)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprint(baseUrl, prefix, getPriceEndpoint, "?", params), nil)
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

func (manager *ExchangeManager) GetMarketTrades(symbol symbol.Assets) ([]trade.MarketTrades, error) {
	params := fmt.Sprintf("currency_pair=%s_%s", symbol.Base, symbol.Quote)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprint(baseUrl, prefix, getMarketTrades, "?", params), nil)
	if err != nil {
		return nil, err
	}
	gaterequest.SetHeader(req)

	response, err := manager.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer gateresponse.CloseBody(response)

	return gateresponse.ParseMarketTrades(response)
}

func (manager *ExchangeManager) GetOrderBook(symbol symbol.Assets) (trade.OrderBook, error) {
	params := fmt.Sprintf("currency_pair=%s_%s&limit=1000", symbol.Base, symbol.Quote)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprint(baseUrl, prefix, getOrderBook, "?", params), nil)
	if err != nil {
		return trade.OrderBook{}, err
	}
	gaterequest.SetHeader(req)

	response, err := manager.client.Do(req)
	if err != nil {
		return trade.OrderBook{}, err
	}

	defer gateresponse.CloseBody(response)

	return gateresponse.ParseOrderBook(response)
}
