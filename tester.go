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

    // prompt user for how many managers to look through
    fmt.Println("Enter the number for which option you would like:")
    fmt.Println("(1) top transfers of top 50 managers")
    fmt.Println("(2) top transfers of top 100 managers")
    fmt.Println("(3) top transfers of top 1,000 managers")
    // fmt.Println("(4) top transfers of top 10,000 managers")
    // TODO
    // fmt.Println("(1) top transfers of top 100 managers")
    var option1 int
    var numManagers int
    fmt.Scanln(&option1)
    if option1 == 1 {
        numManagers = 50
    } else if option1 == 2 {
        numManagers = 100
    } else if option1 == 3 {
        numManagers = 1000
    }

    // fetch top 10 managers and initialize map for player counts    
	topGs := topManagers(numManagers)

    // prompt for transfers in or out
    fmt.Println("Would you like to see transfers in or transfers out?")
    fmt.Println("(1) In")
    fmt.Println("(2) Out")
    var option2 int
    var transferInOrOut string
    fmt.Scanln(&option2)
    if option2 == 1 {
        transferInOrOut = "in"
    } else {
        transferInOrOut = "out"
    }
    
    // create batches for concurrency
    b1, b2, b3, b4, b5 := numManagers/5, 2*numManagers/5, 3*numManagers/5, 4*numManagers/5, numManagers    
	batch1 := topGs[:b1]
    batch2 := topGs[b1:b2]
    batch3 := topGs[b2:b3]
    batch4 := topGs[b3:b4]
    batch5 := topGs[b4:b5] 
    c1 := make(chan map[int]int)
    c2 := make(chan map[int]int)
    c3 := make(chan map[int]int)
    c4 := make(chan map[int]int)
    c5 := make(chan map[int]int)
    // var transferInCount1, transferInCount2, transferInCount3, transferInCount4 map[int]int

    go topTransfersForGameweek(currentGameweek, batch1, c1, transferInOrOut)  
    go topTransfersForGameweek(currentGameweek, batch2, c2, transferInOrOut)  
    go topTransfersForGameweek(currentGameweek, batch3, c3, transferInOrOut)  
    go topTransfersForGameweek(currentGameweek, batch4, c4, transferInOrOut) 
    go topTransfersForGameweek(currentGameweek, batch5, c5, transferInOrOut)  
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
    
    // Print out top 10
    fmt.Println("Top 10 Players transferred ", transferInOrOut, " are:")    
    for i:= 0; i < 10; i++ {
        k := keys[i]
        fmt.Println(i+1, playerName(k, playerMap), "| Number of managers transferring ", transferInOrOut, ":", transferInCounts[k])
    }    
    

    // Formatted string, such as "2h3m0.5s" or "4.503μs"
    fmt.Println("time taken = ", duration)
}




