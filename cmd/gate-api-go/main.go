package main

import (
	"fmt"
	"github.com/posipaka-trade/gate-api-go/pkg/gate"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"os"
	"time"
)

func main() {
	mgr := gate.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})
	defer mgr.Finish()

	time.Sleep(5 * time.Second)

	startTime := time.Now()
	//_, _ = mgr.GetCurrentPrice(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT",
	//})
	//fmt.Println(time.Since(startTime).String())
	//
	//time.Sleep(time.Second)
	//startTime = time.Now()
	//_, _ = mgr.GetCurrentPrice(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT",
	//})
	//fmt.Println(time.Since(startTime).String())
	//
	//time.Sleep(time.Second)
	//startTime = time.Now()
	//_, _ = mgr.GetCurrentPrice(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT",
	//})
	//fmt.Println(time.Since(startTime).String())
	//
	//time.Sleep(time.Second)
	//startTime = time.Now()
	//_, _ = mgr.GetCurrentPrice(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT",
	//})
	//fmt.Println(time.Since(startTime).String())
	//
	//time.Sleep(time.Second)
	//startTime = time.Now()
	//_, _ = mgr.GetCurrentPrice(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT",
	//})
	//
	//time.Sleep(110 * time.Second)
	//startTime = time.Now()
	//_, _ = mgr.GetCurrentPrice(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT",
	//})
	info, err := mgr.SetOrder(order.Parameters{
		Assets: symbol.Assets{
			Base:  "USDG",
			Quote: "USDT",
		},
		Side:     order.Buy,
		Type:     order.Limit,
		Quantity: 10,
		Price:    1.1,
	})
	if err != nil {
		fmt.Println(err)
	}

	sth, err := mgr.GetOrderInfo(info.Id, info.Assets)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sth)

	fmt.Println(time.Since(startTime).String())
}
