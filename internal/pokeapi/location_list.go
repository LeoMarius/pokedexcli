package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

// ListLocations récupère la liste des zones de localisation
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Vérifier si les données sont dans le cache
	if data, found := c.cache.Get(url); found {
		fmt.Println("Données du cache")
		var locationsResp RespShallowLocations
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	// Si les données ne sont pas en cache, on effectue la requête HTTP
	fmt.Println("Données de l'API")
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	// Lire la réponse HTTP
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Stocker la réponse dans le cache
	c.cache.Add(url, dat)

	// Désérialiser la réponse
	var locationsResp RespShallowLocations
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
