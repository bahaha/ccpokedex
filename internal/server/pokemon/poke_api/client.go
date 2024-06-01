package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var (
	api = "https://pokeapi.co/api/v2"
)

func send(endpoint string, resp interface{}) (int, error) {
	// TODO: cache

	req, err := http.NewRequest(http.MethodGet, api+endpoint, nil)
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
