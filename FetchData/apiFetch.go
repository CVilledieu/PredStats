package APIData

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetData(url string) (*PredecessorMatch, error) {
	resp, err := http.Get(url) // Replace with your API
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := [10]PredecessorMatch{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	fmt.Println(&data)
	firstMatch, err := returnOneMatch(data)
	if err != nil {
		return nil, err
	}

	return firstMatch, nil
}

func returnOneMatch(match [10]PredecessorMatch) (*PredecessorMatch, error) {

	singleMatch := match[0]
	return &singleMatch, nil
}
