// This package handles the internals communications with the pokeapi server. It creates clients,
// makes requests and checks for cache.
package pokeapi

import (
	"net/http"
	"time"

	"github.com/kaiserkimguin/pokedex/internal/pokecache"
)

// struct contains an http.Client to make the request and a cache to avoid additional requests.
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// creates a Client to be called by commandMap

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		// Since this Client is to be used directly it is called by value.
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
