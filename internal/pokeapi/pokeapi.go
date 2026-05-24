package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

const BASE_URL string = "https://pokeapi.co/api/v2"

type LocationAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct{
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (client *Client) GetLocationAreaByName(name string) (LocationAreaResponse, error) {
	url := BASE_URL + "/location-area/" + name
	if data, found := client.cache.Get(url); found {
		locationAreaResponse := LocationAreaResponse{}
		if err := json.Unmarshal(data, &locationAreaResponse); err != nil {
			return LocationAreaResponse{}, err
		}

		return locationAreaResponse, nil
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

	client.cache.Add(url, data)

	locationAreaResponse := LocationAreaResponse{}
	if err := json.Unmarshal(data, &locationAreaResponse); err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}

func (client *Client) GetLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := BASE_URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	if data, found := client.cache.Get(url); found {
		locationAreasResponse := LocationAreasResponse{}
		if err := json.Unmarshal(data, &locationAreasResponse); err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	client.cache.Add(url, data)

	locationAreasResponse := LocationAreasResponse{}
	if err := json.Unmarshal(data, &locationAreasResponse); err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasResponse, nil
}












