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
    numManagers := 50
	topGs := topManagers(numManagers)
    
    // create batches for concurrency
	batch1 := topGs[:10]
    batch2 := topGs[10:20]
    batch3 := topGs[20:30]
    batch4 := topGs[30:40]
    batch5 := topGs[40:50] 
    c1 := make(chan map[int]int)
    c2 := make(chan map[int]int)
    c3 := make(chan map[int]int)
    c4 := make(chan map[int]int)
    c5 := make(chan map[int]int)
    // var transferInCount1, transferInCount2, transferInCount3, transferInCount4 map[int]int

    go topTransfersForGameweek(currentGameweek, batch1, c1)  
    go topTransfersForGameweek(currentGameweek, batch2, c2)  
    go topTransfersForGameweek(currentGameweek, batch3, c3)  
    go topTransfersForGameweek(currentGameweek, batch4, c4) 
    go topTransfersForGameweek(currentGameweek, batch5, c5)  
    transferInCounts := make(map[int] int)
    for {
        transferInCount1, open := <- c1
        if !open {
            break
        }      
        for k, v := range transferInCount1 {
            // fmt.Println(" TES TEST ", playerName(k, playerMap)) 
            val, ok := transferInCounts[k]            
                if ok {            
                    transferInCounts[k] = val+v
                }else {                    
                    transferInCounts[k] = v
                }
        }  
    }
    for {
        transferInCount2, open := <- c2
        if !open {
            break
        }  
        for k, v := range transferInCount2 {
            val, ok := transferInCounts[k]            
                if ok {            
                    transferInCounts[k] = val+v
                }else {
                    transferInCounts[k] = v
                }
        }      
    }
    for {
        transferInCount3, open := <- c3
        if !open {
            break
        }    
        for k, v := range transferInCount3 {
            val, ok := transferInCounts[k]            
                if ok {            
                    transferInCounts[k] = val+v
                }else {
                    transferInCounts[k] = v
                }
        }    
    }
    for {
        transferInCount4, open := <- c4
        if !open {
            break
        }   
        for k, v := range transferInCount4 {
            val, ok := transferInCounts[k]            
                if ok {            
                    transferInCounts[k] = val+v
                }else {
                    transferInCounts[k] = v
                }
        }     
    }
    for {
        transferInCount5, open := <- c5
        if !open {
            break
        }   
        for k, v := range transferInCount5 {
            val, ok := transferInCounts[k]            
                if ok {            
                    transferInCounts[k] = val+v
                }else {
                    transferInCounts[k] = v
                }
        }     
    }

    // transferInCounts := topTransfersForGameweek(currentGameweek, topGs)    
    
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
        fmt.Println(i+1, playerName(k, playerMap), "number of managers transferring in = ", transferInCounts[k])
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




