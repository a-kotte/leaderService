package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func callAPI(url string) []byte{
	// call GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
        os.Exit(1)
	}	
	userAgent := "fplLeaderboardService/1.0.0"		
	req.Header.Set("User-Agent", userAgent )
	res, err := http.DefaultClient.Do(req)	
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	// Read resukt
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	return resBody
}