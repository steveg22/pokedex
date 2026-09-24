package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *config) error {
	if cfg.PreviousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(cfg.PreviousUrl)
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

	cfg.NextUrl = response.Next
	cfg.PreviousUrl = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}
