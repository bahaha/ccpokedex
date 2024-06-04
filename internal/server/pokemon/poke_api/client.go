package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var (
	api = "https://pokeapi.co/api/v2"
	// TODO: extract as an isolated module
	trends_api = "https://zukan.pokemon.co.jp/zukan-api/api"
)

func send(endpoint string, resp interface{}) (int, error) {
	// TODO: cache

	url := api + endpoint
	if endpoint == "/pickup" {
		url = trends_api + endpoint
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if err != nil {
		return 0, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(req)
	status := response.StatusCode
	if err != nil {
		return status, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return status, err
	}

	return status, json.Unmarshal(body, &resp)
}
