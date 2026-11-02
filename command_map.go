package main

import "fmt"

func commandMap(cfg *config) error {
	url := baseURL + "/location-area"
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}
	return showLocationPage(cfg, url)
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return showLocationPage(cfg, *cfg.prevLocationsURL)
}

func showLocationPage(cfg *config, url string) error {
	data, err := fetchLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}
	return nil
}