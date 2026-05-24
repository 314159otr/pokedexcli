package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)
const BASE_URL string = "https://pokeapi.co/api/v2"

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func (client *Client) GetLocationAreas(pageUrl *string) (LocationAreaResponse, error) {
	url := BASE_URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	if err := json.Unmarshal(data, &locationAreaResponse); err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}












