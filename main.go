package main

import (
	"fmt"

	APIData "PredStats/FetchData"
)

func main() {

	url := "https://api.example.com"  // Replace with your API
	data, err := APIData.GetData(url) // Call the fetchData function from the fetchdata package
	if err != nil {
		fmt.Println(err)
		return
	}

}
