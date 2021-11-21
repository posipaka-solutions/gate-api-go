package main

import (
	"github.com/posipaka-trade/gate-api-go/pkg/gate"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"os"
)

func main() {
	mgr := gate.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})
	//	defer mgr.Finish()

	//time.Sleep(5 * time.Second)
	_, _ = mgr.GetOrderBook(symbol.Assets{
		Base:  "QI",
		Quote: "USDT",
	})
	//startTime := time.Now()
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
	//
	//fmt.Println(time.Since(startTime).String())
}
