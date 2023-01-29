package main

import (
	"encoding/json"
	"strconv"
)

func team(managerID int, gameweek int) []int{
	// set url for API call
	url := "https://fantasy.premierleague.com/api/entry/" + strconv.Itoa(managerID) + "/event/" + strconv.Itoa(gameweek) + "/picks/"

	// call API and unmarshall into teamweekly struct 
	resBody := callAPI(url)
	var responseObject TeamWeekly
	json.Unmarshal(resBody, &responseObject)

	// use the array of players to return array of player IDs
	teamIDs := responseObject.Picks	
	team := make([]int, len(teamIDs))
	for i,player := range teamIDs {
		team[i] = player.Element
	}
	return team
}

// Set Difference: A - B
func Difference(a, b []int) (diff []int) {
	m := make(map[int]bool)

	for _, item := range b {
			m[item] = true
	}

	for _, item := range a {
			if _, ok := m[item]; !ok {
					diff = append(diff, item)
			}
	}
	return
}