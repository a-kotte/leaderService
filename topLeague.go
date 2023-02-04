package main

import (
	"encoding/json"
)

func topManagers(numManager int) []int {

	url := "https://fantasy.premierleague.com/api/leagues-classic/314/standings?page_standings=1"

	resBody := callAPI(url)

	// unmarshall JSON results
	var responseObject LeagueInfo
	json.Unmarshal(resBody, &responseObject)	
	topTeams := responseObject.Standings.Results[:numManager]

	topTeamIDs := make([]int, numManager) 	
	
	for i,team := range topTeams {
		topTeamIDs[i] = team.Entry
	}
	
	return topTeamIDs
	// top10results := results[0:2]
	// fmt.Println("value is", gameweek)
	// fmt.Println(currentGameweek)
	
}