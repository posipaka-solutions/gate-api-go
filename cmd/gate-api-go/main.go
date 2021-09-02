package main

import (
	"fmt"
	"github.com/posipaka-trade/gate-api-go/pkg/gate"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
)

func main() {
	mgr := gate.New(exchangeapi.ApiKey{
		Key:    "",
		Secret: "",
	})
	result, err := mgr.SetOrder(order.Parameters{
		Assets: symbol.Assets{
			Base:  "ETH",
			Quote: "USDT"},
		Side:     order.Buy,
		Type:     order.Limit,
		Quantity: 0.001,
		Price:    3300,
	})
	fmt.Println(result)
	price, err := mgr.GetCurrentPrice(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})
	fmt.Println(price)

	if err != nil {
		panic(err.Error())
	}
}
