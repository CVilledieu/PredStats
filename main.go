package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func main() {
	//Time since Epoch
	now := time.Now()
	Epoch := strconv.FormatInt(now.Unix(), 10)

	//End point url should end with time since Epoch
	MatchesUrl := "https://backend.production.omeda-aws.com/api/public/get-matches-since/" + Epoch

	resp, err := http.Get(MatchesUrl)
	if err != nil {
		err.Error()
	}

	resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err.Error()
	}
	fmt.Println(body)

}
