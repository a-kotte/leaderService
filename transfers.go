package main

import (
	"encoding/json"
	"strconv"
)

func transfers(managerID int) TransferHistory {
	// set url for API call
	url := "https://fantasy.premierleague.com/api/entry/" + strconv.Itoa(managerID) + "/transfers"

	// call API and unmarshall into teamweekly struct 
	resBody := callAPI(url)
	var transfers TransferHistory
	json.Unmarshal(resBody, &transfers)
	return transfers
}

func transfersInForGameweek(managerID, gameweek int) []int {
	// fetch all transfers for manager	
	transfers := transfers(managerID)

	// only store transfer In's for the gameweek
	var gameweekTransfersIn []int
	for _, transfer := range transfers {
		if transfer.Event < gameweek {
			return gameweekTransfersIn
		}
		
		if transfer.Event == gameweek {
			gameweekTransfersIn = append(gameweekTransfersIn, transfer.ElementIn)
		}			
	}
	return gameweekTransfersIn
}

func transfersOutForGameweek(managerID, gameweek int) []int {
	// fetch all transfers for manager	
	transfers := transfers(managerID)

	// only store transfer In's for the gameweek
	var gameweekTransfersOut []int
	for _, transfer := range transfers {
		if transfer.Event < gameweek {
			return gameweekTransfersOut
		}
		
		if transfer.Event == gameweek {
			gameweekTransfersOut = append(gameweekTransfersOut, transfer.ElementOut)
		}			
	}
	return gameweekTransfersOut
}

// // return top transfers IN for top 50 managers
// func topTransfersForGameweek(gameweek int, managers []int) map[int]int {
// 	//  initialize map for player counts	
//     transferInCounts := make(map[int]int)	
    
//     for _,G := range managers {
//         // fetch current and prior team and use difference to make list of transferred players
//         currTeam := team(G, gameweek)
//         oldTeam := team(G, gameweek-1)                
//         newPlayers := Difference(currTeam, oldTeam)
//         // newPlayers := transfersInForGameweek(G, gameweek)

//         // for each transferred in player, add to the map
//         for _,player:= range newPlayers {            
//             val, ok := transferInCounts[player]            
//             if ok {            
//                 transferInCounts[player] = val+1
//             }else {
//                 transferInCounts[player] = 1
//             }
//         }
//     }	
// 	return transferInCounts
// }

// return top transfers IN for top 50 managers
func topTransfersForGameweek(gameweek int, managers []int, c chan map[int]int, transferInOrOut string ) {
	//  initialize map for player counts	
    transferCounts := make(map[int]int)	
    
    for _,G := range managers {
        // fetch current and prior team and use difference to make list of transferred players
        currTeam := team(G, gameweek)
        oldTeam := team(G, gameweek-1)  
		var newPlayers []int      
		if transferInOrOut == "in"{
			newPlayers = Difference(currTeam, oldTeam)
		} else {
			newPlayers = Difference(oldTeam, currTeam)
		}
        		
        // newPlayers := transfersInForGameweek(G, gameweek)

        // for each transferred in player, add to the map
        for _,player:= range newPlayers {            
            val, ok := transferCounts[player]            
            if ok {            
                transferCounts[player] = val+1
            }else {
                transferCounts[player] = 1
            }
        }
    }
	c <- transferCounts	
	close(c)
}

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