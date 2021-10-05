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
		Key:    "d51bd83810fadd460c3699f31b945a0c",
		Secret: "d4d27639f9de6ed3f37ab1cbd2e98d26274eed1abb2a1100a1758c207fb2dbc0",
	})
	result, err := mgr.SetOrder(order.Parameters{
		Assets: symbol.Assets{
			Base:  "ETH",
			Quote: "USDT"},
		Side:     order.Buy,
		Type:     order.Limit,
		Quantity: 0.001,
		Price:    2000,
	})
	fmt.Println(result)
	price, err := mgr.GetCurrentPrice(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})
	fmt.Println(price)
	balance, err := mgr.GetAssetBalance("USDT")

	fmt.Println(balance)

	if err != nil {
		panic(err.Error())
	}
}
