package trade

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"time"
)

type MarketTrades struct {
	Id         int
	CreateTime time.Time
	Side       order.Side
	Amount     float64
	Price      float64
}
