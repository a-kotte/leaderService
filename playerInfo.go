package main

// return the name of a player using the player's ID and the map
func playerName(playerID int, playerMap map[int]string) string {				
	return playerMap[playerID] 		
}

// create a map of player IDs -> player names using bootstrap
func playerNameMap() map[int]string {	
	
	players := bootstrap().Elements
	playerNameMap := make(map[int]string)

	for _, player := range players {				
		playerNameMap[player.ID] = player.FirstName + " " + player.SecondName										
	}	
	return playerNameMap	
}

