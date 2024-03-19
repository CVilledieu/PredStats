package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Data struct {
	// Define your struct based on the API response
}

func fetchData() (*Data, error) {
	resp, err := http.Get("http://example.com/api") // Replace with your API
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	data, err := fetchData()
	if err != nil {
		fmt.Println(err)
		return
	}

	a := app.New()
	w := a.NewWindow("API Data")

	label := widget.NewLabel(fmt.Sprintf("%v", data)) // Update this based on your struct
	w.SetContent(label)

	w.ShowAndRun()
}
