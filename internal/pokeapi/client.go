package pokeapi

import (
	"net/http"
	"time"

	"github.com/LeoMarius/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Minute), // Exemple d'intervalle de cache

	}
}
