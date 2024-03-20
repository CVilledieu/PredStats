package APIData

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetData(url string) (*struct{}, error) {
	resp, err := http.Get(url) // Replace with your API
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data struct{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
