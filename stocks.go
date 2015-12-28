// API calls to get stocks from API

package main

import (
  "io/ioutil"
  "encoding/json"
  "fmt"
  "net/http"
)

type StockList struct {
  Ok bool `json:"ok"`
  Error string `json:"error"`
  Venue string
  Symbols []Stock `json:"symbols"`
}

type Stock struct {
  Name string `json:"name"`
  Symbol string `json:"symbol"`
}

func GetStocks(venue string) (StockList, error) {

  response, err := http.Get("https://api.stockfighter.io/ob/api/venues/" + venue + "/stocks")
  if err != nil {
    fmt.Println(err)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  }

  var i interface{}
  err = json.Unmarshal(body, &i)
  // if err != nil {
  //   fmt.Println(err)
  // }
  fmt.Println(i)

  var stockList StockList
  stockList.Venue = venue
  err = json.Unmarshal(body, &stockList)
  if err != nil {
    return stockList, err
  }
  return stockList, nil

}