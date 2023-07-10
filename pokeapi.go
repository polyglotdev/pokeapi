package pokeapi

import "net/http"

// Client is the http client for the pokeapi -- for testing purposes only
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient returns a new pokeapi client
func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		baseURL:    "https://pokeapi.co/api/v2/",
	}
}
