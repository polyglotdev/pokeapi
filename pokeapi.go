package pokeapi

import "net/http"

// Client is the http client for the pokeapi -- for testing purposes only
type Client struct {
	httpClient *http.Client
	baseURL    string
}
