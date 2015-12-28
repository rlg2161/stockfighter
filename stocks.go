// API calls to get stocks from API

package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
)

type StockList struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Venue   string
	Symbols []Stock `json:"symbols"`
}

type Stock struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Order struct {
	Price    int  `json:"price"`
	Quantity int  `json:"qty"`
	IsBuy    bool `json:"isBuy"`
}

type OrderBook struct {
	Ok        bool    `json:"ok"`
	Error     string  `json:"error"`
	Venue     string  `json:"venue"`
	Symbol    string  `json:"symbol"`
	Bids      []Order `json:"bids"`
	Asks      []Order `json:"asks"`
	Timestamp string  `json:"ts"`
}

func GetStocks(venue string) (StockList, error) {
	var stockList StockList
	response, err := http.Get("https://api.stockfighter.io/ob/api/venues/" + venue + "/stocks")
	if err != nil {
		return stockList, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return stockList, err
	}

	stockList.Venue = venue
	err = json.Unmarshal(body, &stockList)
	if err != nil {
		return stockList, err
	}
	return stockList, nil

}

func GetOrderBook(venue string, symbol string) (OrderBook, error) {
	var orderBook OrderBook
	response, err := http.Get("https://api.stockfighter.io/ob/api/venues/" + venue + "/stocks/" + symbol)
	if err != nil {
		return orderBook, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return orderBook, err
	}

	err = json.Unmarshal(body, &orderBook)
	if err != nil {
		return orderBook, err
	}
	return orderBook, nil
}
