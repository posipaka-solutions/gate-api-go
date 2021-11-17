package trade

type OrderBook struct {
	Asks []AskAndBids
	Bids []AskAndBids
}
type AskAndBids struct {
	Price    float64
	Quantity float64
}
