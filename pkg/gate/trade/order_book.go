package trade

type OrderBook struct {
	Asks []Ask
	Bids []Bid
}
type Ask struct {
	Price    float64
	Quantity float64
}
type Bid struct {
	Price    float64
	Quantity float64
}
