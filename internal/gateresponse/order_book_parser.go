package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"net/http"
	"strconv"
)

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

func ParseOrderBook(response *http.Response) (OrderBook, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return OrderBook{}, err
	}
	body, isOk := bodyI.(map[string]interface{})
	if isOk != true {
		return OrderBook{}, errors.New("[gateresponse] -> error when casting order book body")
	}
	asks, isOk := body[pnames.Asks].([]interface{})
	if isOk != true {
		return OrderBook{}, errors.New("[gateresponse] -> error when casting order book asks")
	}
	bids, isOk := body[pnames.Bids].([]interface{})
	if isOk != true {
		return OrderBook{}, errors.New("[gateresponse] -> error when casting order book bids")
	}

	return asksAndBidsParse(asks, bids)
}

//TODO: Make parallel asks and bids parsing
func asksAndBidsParse(asks, bids []interface{}) (OrderBook, error) {
	var orderBook OrderBook
	var ask Ask
	var bid Bid
	var err error

	for i := 0; i < len(asks); i++ {
		asksI, isOk := asks[i].([]interface{})
		if isOk != true {
			return OrderBook{}, errors.New("[gateresponse] -> error when casting order book asks value")
		}
		priceStr, isOk := asksI[0].(string)
		if isOk != true {
			return OrderBook{}, errors.New("[gateresponse] -> error when casting order book ask price")
		}
		quantityStr, isOk := asksI[1].(string)
		if isOk != true {
			return OrderBook{}, errors.New("[gateresponse] -> error when casting order book ask quantity")
		}
		ask.Price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return OrderBook{}, errors.New("[gateresponse] -> error when parsing order book ask price to float64")
		}
		ask.Quantity, err = strconv.ParseFloat(quantityStr, 64)
		if err != nil {
			return OrderBook{}, errors.New("[gateresponse] -> error when parsing order book ask quantity to float64")
		}
		orderBook.Asks = append(orderBook.Asks, ask)
	}
	for i := 0; i < len(bids); i++ {
		bidsI, isOk := bids[i].([]interface{})
		if isOk != true {
			return OrderBook{}, errors.New("[gateresponse] -> error when casting order book bids value")
		}
		priceStr, isOk := bidsI[0].(string)
		if isOk != true {
			return OrderBook{}, errors.New("[gateresponse] -> error when casting order book bid price")
		}
		quantityStr, isOk := bidsI[1].(string)
		if isOk != true {
			return OrderBook{}, errors.New("[gateresponse] -> error when casting order book bid quantity")
		}
		bid.Price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return OrderBook{}, errors.New("[gateresponse] -> error when parsing order book bid price to float64")
		}
		bid.Quantity, err = strconv.ParseFloat(quantityStr, 64)
		if err != nil {
			return OrderBook{}, errors.New("[gateresponse] -> error when parsing order book bid quantity to float64")
		}
		orderBook.Bids = append(orderBook.Bids, bid)
	}
	return orderBook, nil
}
