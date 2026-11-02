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

type locationAreaDetail struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func getJSON(url string, target any) error {
	body, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status: %s", resp.Status)
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(url, body)
	}
	return json.Unmarshal(body, target)
}

func fetchLocationAreas(url string) (locationAreasResponse, error) {
	var data locationAreasResponse
	err := getJSON(url, &data)
	return data, err
}

func fetchLocationArea(name string) (locationAreaDetail, error) {
	var data locationAreaDetail
	err := getJSON(baseURL+"/location-area/"+name, &data)
	return data, err
}