package gate

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
	"sync"
	"time"
)

const baseUrl = "https://api.gateio.ws"

type ExchangeManager struct {
	apiKey exchangeapi.ApiKey

	client *http.Client

	wg        sync.WaitGroup
	isWorking bool
}

func New(key exchangeapi.ApiKey) *ExchangeManager {
	mgr := &ExchangeManager{
		apiKey: key,
		client: &http.Client{
			Transport: http.DefaultTransport,
		},
		isWorking: true,
	}

	mgr.wg.Add(1)
	go func() {
		defer mgr.wg.Done()
		for mgr.isWorking {
			_, _ = mgr.client.Get(baseUrl + prefix + "/")
			time.Sleep(75 * time.Second)
		}
	}()

	return mgr
}

// Finish completes inner goroutines
func (manager *ExchangeManager) Finish() {
	manager.isWorking = false
	manager.wg.Wait()
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
