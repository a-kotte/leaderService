package main

import (
	"encoding/json"
	"strconv"
)

func topManagers(numManager int) []int {

	url := "https://fantasy.premierleague.com/api/leagues-classic/314/standings?page_standings=1"
	resBody := callAPI(url)

	// unmarshall JSON results
	var responseObject LeagueInfo
	json.Unmarshal(resBody, &responseObject)	
	topTeamIDs := make([]int, 50) 	
		
	// handle numManagers more than 50 (API limits 50 per page)
	if numManager > 50 {
		// initialize topTeams to hold first page's results
		topTeams := responseObject.Standings.Results		
		for i,team := range topTeams {
			topTeamIDs[i] = team.Entry
		}
		pages := numManager/50
		for i := 0; i < pages; i++ {
			// can add concurrency to optimize this
			page := 2 + i
			newURL := "https://fantasy.premierleague.com/api/leagues-classic/314/standings?page_standings=" + strconv.Itoa(page)
			resBody = callAPI((newURL))
			var newResponseObject LeagueInfo
			json.Unmarshal(resBody, &newResponseObject)
			nextTopTeams := newResponseObject.Standings.Results			
			for _,team := range nextTopTeams {
				topTeamIDs = append(topTeamIDs, team.Entry)				
			}			
			
			// topTeams = append(topTeams, newResponseObject.Standings.Results[:50]) 
		}
	} else {		
		topTeams := responseObject.Standings.Results[:numManager]
		topTeamIDs := make([]int, numManager) 		
		for i,team := range topTeams {
			topTeamIDs[i] = team.Entry
		}
		return topTeamIDs
	}

	return topTeamIDs
	// topTeams := responseObject.Standings.Results[:numManager]

	// topTeamIDs := make([]int, numManager) 	
	
	// for i,team := range topTeams {
	// 	topTeamIDs[i] = team.Entry
	// }
	
	// return topTeamIDs
	// top10results := results[0:2]
	// fmt.Println("value is", gameweek)
	// fmt.Println(currentGameweek)
	
}