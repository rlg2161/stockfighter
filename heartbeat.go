// API calls to heartbeat API

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Heartbeat struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type HeartbeatVenue struct {
	Ok    bool   `json:"ok"`
	Venue string `json:"venue"`
}

func HeartbeatAPI() Heartbeat {

	response, err := http.Get("https://api.stockfighter.io/ob/api/heartbeat")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	var heartbeat Heartbeat
	json.Unmarshal(body, &heartbeat)
	return heartbeat

}

func HeartbeatVenueAPI(venue string) HeartbeatVenue {
	resp, err := http.Get("https://api.stockfighter.io/ob/api/venues/" + venue + "/heartbeat")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	var heartbeatVenue HeartbeatVenue
	json.Unmarshal(body, &heartbeatVenue)
	return heartbeatVenue
}
