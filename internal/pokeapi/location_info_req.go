package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationIDInfo(area string) (LocationIDResp, error) {
	endpoint := "/location-area/" + area
	fullURL := baseURL + endpoint
	if entry, ok := c.cache.Get(fullURL); ok {
		locationInfo := LocationIDResp{}
		err := json.Unmarshal(entry, &locationInfo)
		if err != nil {
			return LocationIDResp{}, err
		}
		return locationInfo, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationIDResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationIDResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationIDResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationIDResp{}, err
	}
	locationInfo := LocationIDResp{}
	err = json.Unmarshal(data, &locationInfo)
	if err != nil {
		return LocationIDResp{}, err
	}
	c.cache.Add(fullURL, data)
	return locationInfo, nil
}
