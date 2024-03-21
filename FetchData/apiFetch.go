package APIData

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetData(url string) (*data, error) {
	resp, err := http.Get(url) // Replace with your API
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type Data interface {
		Matches() []string
	}
	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
