package APIData

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetData(url string) (*[]PredecessorMatch, error) {
	resp, err := http.Get(url) // Replace with your API
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tenMatches := make([]PredecessorMatch, 10)

	err = json.Unmarshal(body, &tenMatches)
	if err != nil {
		return nil, err
	}

	return &tenMatches, nil
}
