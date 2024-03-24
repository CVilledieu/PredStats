package main

import (
	"fmt"

	APIData "PredStats/FetchData"
)

func main() {
	// Get the current time in seconds since epoch
	//timeSinceEpoch := strconv.FormatInt(time.Now().Unix(), 10)

	// Replace with your API
	// Call the fetchData function from the fetchdata package
	url := "https://pred.saibotu.de/api/public/get-matches-since/1711235812/"
	data, err := APIData.GetData(url)
	if err != nil {
		fmt.Println("there was an error: ", err)
		return
	}
	matchId := *data
	firstMatch := matchId[0]
	fmt.Println("Match ID: ", firstMatch.PlayerData[0].MinionData.MinionsKilled)
}
