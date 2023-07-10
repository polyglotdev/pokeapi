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
