package main

import (
	"fmt"
	"github.com/posipaka-trade/gate-api-go/pkg/gate"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"os"
)

func main() {
	mgr := gate.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})
	result, err := mgr.SetOrder(order.Parameters{
		Assets: symbol.Assets{
			Base:  "XRP",
			Quote: "USDT"},
		Side:     order.Buy,
		Type:     order.Limit,
		Quantity: 4,
		Price:    2,
	})
	fmt.Println(result)
	balance, err := mgr.GetAssetBalance("USDT")
	fmt.Println(balance)

	if err != nil {
		panic(err.Error())
	}
}
