package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	if data, ok := c.cache.Get(url); ok {
		var locationResp Location
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, nil
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, nil
	}

	var locationResp Location
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, nil
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
