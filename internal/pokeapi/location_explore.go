package pokeapi

import (
	"encoding/json"
	"io"
)

// ListLocations récupère la liste des zones de localisation
func (c *Client) ExploreLocations(pageURL *string, explore_area string) (RespExploreLocations, error) {
	url := baseURL + "/location-area/" + explore_area
	if pageURL != nil {
		url = *pageURL
	}

	// Vérifier si les données sont dans le cache
	if data, found := c.cache.Get(url); found {
		// fmt.Println("Données du cache")
		var locationsResp RespExploreLocations
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespExploreLocations{}, err
		}
		return locationsResp, nil
	}

	// Si les données ne sont pas en cache, on effectue la requête HTTP
	// fmt.Println("Données de l'API")
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return RespExploreLocations{}, err
	}
	defer resp.Body.Close()

	// Lire la réponse HTTP
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocations{}, err
	}

	// Stocker la réponse dans le cache
	c.cache.Add(url, dat)

	// Désérialiser la réponse
	var locationsResp RespExploreLocations
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespExploreLocations{}, err
	}

	return locationsResp, nil
}
