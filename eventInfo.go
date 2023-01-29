package main

import (
	"encoding/json"
)

func eventInfo() int {
	url := "https://fantasy.premierleague.com/api/event-status/"
	resBody := callAPI(url)

	// unmarshall JSON results
	var responseObject EventStatus
	json.Unmarshal(resBody, &responseObject)	
	
	currentGameweek := responseObject.Status[0].Event
	
	return currentGameweek
}