package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"BugsDoingCoding/Bootdev.git/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

var cache = pokecache.NewCache(5 * time.Minute)

type locationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchLocationAreas(url string) (locationAreasResponse, error) {
	body, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return locationAreasResponse{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return locationAreasResponse{}, fmt.Errorf("bad status: %s", resp.Status)
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return locationAreasResponse{}, err
		}
		cache.Add(url, body)
	}

	var data locationAreasResponse
	err := json.Unmarshal(body, &data)
	if err != nil {
		return locationAreasResponse{}, err
	}
	return data, nil
}