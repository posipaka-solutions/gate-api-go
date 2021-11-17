package gateresponse

import (
	"errors"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"github.com/posipaka-trade/gate-api-go/pkg/gate/trade"
	"net/http"
	"strconv"
)

func ParseOrderBook(response *http.Response) (trade.OrderBook, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return trade.OrderBook{}, err
	}
	body, isOk := bodyI.(map[string]interface{})
	if isOk != true {
		return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book body")
	}
	asks, isOk := body[pnames.Asks].([]interface{})
	if isOk != true {
		return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book asks")
	}
	bids, isOk := body[pnames.Bids].([]interface{})
	if isOk != true {
		return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book bids")
	}

	return asksAndBidsParse(asks, bids)
}

//TODO: Make parallel asks and bids parsing
func asksAndBidsParse(asks, bids []interface{}) (trade.OrderBook, error) {
	var orderBook trade.OrderBook
	var askAndBid trade.AskAndBids
	var err error

	for i := 0; i < len(asks); i++ {
		asksI, isOk := asks[i].([]interface{})
		if isOk != true {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book asks value")
		}
		priceStr, isOk := asksI[0].(string)
		if isOk != true {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book ask price")
		}
		quantityStr, isOk := asksI[1].(string)
		if isOk != true {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book ask quantity")
		}
		askAndBid.Price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when parsing order book ask price to float64")
		}
		askAndBid.Quantity, err = strconv.ParseFloat(quantityStr, 64)
		if err != nil {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when parsing order book ask quantity to float64")
		}
		orderBook.Asks = append(orderBook.Asks, askAndBid)
	}
	for i := 0; i < len(bids); i++ {
		bidsI, isOk := bids[i].([]interface{})
		if isOk != true {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book bids value")
		}
		priceStr, isOk := bidsI[0].(string)
		if isOk != true {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book bid price")
		}
		quantityStr, isOk := bidsI[1].(string)
		if isOk != true {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when casting order book bid quantity")
		}
		askAndBid.Price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when parsing order book bid price to float64")
		}
		askAndBid.Quantity, err = strconv.ParseFloat(quantityStr, 64)
		if err != nil {
			return trade.OrderBook{}, errors.New("[gateresponse] -> error when parsing order book bid quantity to float64")
		}
		orderBook.Bids = append(orderBook.Bids, askAndBid)
	}
	return orderBook, nil
}
