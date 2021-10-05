package gate

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
)

const baseUrl = "https://api.gateio.ws"

type ExchangeManager struct {
	apiKey exchangeapi.ApiKey

	client *http.Client
}

func New(key exchangeapi.ApiKey) *ExchangeManager {
	return &ExchangeManager{
		apiKey: key,
		client: &http.Client{},
	}
}

var orderSideAlias = map[order.Side]string{
	order.Buy:  "buy",
	order.Sell: "sell",
}

var orderTypeAlias = map[order.Type]string{
	order.Limit:  "limit",
	order.Market: "market",
}

const (
	prefix           = "/api/v4"
	newOrderEndpoint = "/spot/orders"
	getPriceEndpoint = "/spot/tickers"
	balancesEndpoint = "/spot/accounts"
)
