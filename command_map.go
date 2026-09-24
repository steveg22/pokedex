package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiRes struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []location
}

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func commandMap(cfg *config) error {
	res, err := http.Get(cfg.NextUrl)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	var response ApiRes
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	cfg.PreviousUrl = response.Previous
	cfg.NextUrl = response.Next

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}
