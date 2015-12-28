package main

import (
	"fmt"
)

func main() {
	// hb := HeartbeatAPI()
	// hbv := HeartbeatVenueAPI("TESTEX")
	sl, err := GetStocks("TESTEX")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sl)
	ob, err := GetOrderBook("TESTEX", "FOOBAR")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ob)
}
