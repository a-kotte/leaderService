package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {    
    
    // fetch current game week and map of player IDs to Names
    currentGameweek := eventInfo()    
    playerMap := playerNameMap()        

    // start timer
    start := time.Now()
    
    // fetch top 10 managers and initialize map for player counts
    topGs := topManagers(50)
    transferInCounts := make(map[int]int)
    
    for _,G := range topGs {
        // fetch current and prior team and use difference to make list of transferred players
        currTeam := team(G, currentGameweek)
        oldTeam := team(G, currentGameweek-1)                
        newPlayers := Difference(currTeam, oldTeam)

        // for each transferred in player, add to the map
        for _,player:= range newPlayers {            
            val, ok := transferInCounts[player]            
            if ok {            
                transferInCounts[player] = val+1
            }else {
                transferInCounts[player] = 1
            }
        }
    }
    duration := time.Since(start)

    // Sort the list of keys by value
    keys := make([]int, 0, len(transferInCounts))
    for key := range transferInCounts {
        keys = append(keys, key)
    }
    sort.SliceStable(keys, func(i, j int) bool{
        return transferInCounts[keys[i]] > transferInCounts[keys[j]]
    })    
    
    fmt.Println("Top Players are: ")
    for i, k := range keys {
        fmt.Println(i+1, playerName(k, playerMap))
    }

    
    // fmt.Println("len is ", len(topGs))
    // for _,G := range topGs {
    //     fmt.Println("Top G is ", G)
    // }
    
    // playerMap := playerNameMap()

    // playerName := playerName(357, playerMap)
    
    // fmt.Println("player name is", playerName)

    // Code to measure
    

    // Formatted string, such as "2h3m0.5s" or "4.503μs"
    fmt.Println("time taken = ", duration)
}




