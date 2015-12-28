package main

import (
  "fmt"
)

func main() {
  hb := HeartbeatAPI()
  hbv := HeartbeatVenueAPI("TESTEX")
  sl, err := GetStocks("TESTEX")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(hb)
  fmt.Println(hbv)
  fmt.Println(sl)
}