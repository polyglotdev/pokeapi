package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Client is the http client for the pokeapi -- for testing purposes only
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// Pokemon is the struct for a pokemon
type Pokemon struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Nature is the struct for a nature
type Nature struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Stat is the struct for a stat
type Stat struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// NewClient returns a new pokeapi client
func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		baseURL:    "https://pokeapi.co/api/v2/",
	}
}

// GetPokemon takes in a pokemon name or id and returns a pokemon
func (c *Client) GetPokemon(idOrName string) (*Pokemon, error) {
	resp, err := c.httpClient.Get(c.baseURL + "pokemon/" + idOrName + "/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get pokemon: %s", resp.Status)
	}

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}

// GetNature takes in a nature name or id and returns a nature
func (c *Client) GetNature(idOrName string) (*Nature, error) {
	resp, err := c.httpClient.Get(c.baseURL + "nature/" + idOrName + "/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get nature: %s", resp.Status)
	}

	var nature Nature
	if err := json.NewDecoder(resp.Body).Decode(&nature); err != nil {
		return nil, err
	}

	return &nature, nil
}

// GetStat takes in a stat name or id and returns a stat
func (c *Client) GetStat(idOrName string) (*Stat, error) {
	resp, err := c.httpClient.Get(c.baseURL + "stat/" + idOrName + "/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get stat: %s", resp.Status)
	}

	var stat Stat
	if err := json.NewDecoder(resp.Body).Decode(&stat); err != nil {
		return nil, err
	}

	return &stat, nil
}
