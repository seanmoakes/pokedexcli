package pokeapi

import (
	"net/http"
	"time"

	"github.com/seanmoakes/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	pokedex    map[string]Pokemon
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache:   pokecache.NewCache(cacheInterval),
		pokedex: make(map[string]Pokemon),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
