package main

import "encoding/json"


func bootstrap() BootstrapInfo {
	url := "https://fantasy.premierleague.com/api/bootstrap-static/"

	resBody := callAPI(url)

	// unmarshall JSON results
	var responseObject BootstrapInfo
	json.Unmarshal(resBody, &responseObject)	
	
	return responseObject
}