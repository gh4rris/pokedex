package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonTypes(area string) (PokemonTypesResp, error) {
	endpoint := "/location-area/" + area
	fullURL := baseURL + endpoint
	if entry, ok := c.cache.Get(fullURL); ok {
		pokemonTypesResp := PokemonTypesResp{}
		err := json.Unmarshal(entry, &pokemonTypesResp)
		if err != nil {
			return PokemonTypesResp{}, err
		}
		return pokemonTypesResp, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonTypesResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonTypesResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return PokemonTypesResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonTypesResp{}, err
	}
	pokemonTypesResp := PokemonTypesResp{}
	err = json.Unmarshal(data, &pokemonTypesResp)
	if err != nil {
		return PokemonTypesResp{}, err
	}
	return pokemonTypesResp, nil
}
