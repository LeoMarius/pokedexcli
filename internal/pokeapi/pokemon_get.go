package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
)

var My_pokedex = make(map[string]Pokemon)

// ListLocations récupère la liste des zones de localisation
func (c *Client) GetPokemon(pokemon_name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon_name

	// Vérifier si les données sont dans le cache
	if data, found := c.cache.Get(url); found {
		// fmt.Println("Données du cache")
		var pokemonResp Pokemon
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	// Si les données ne sont pas en cache, on effectue la requête HTTP
	// fmt.Println("Données de l'API")
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// Lire la réponse HTTP
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Stocker la réponse dans le cache
	c.cache.Add(url, dat)

	// Désérialiser la réponse
	var pokemonResp Pokemon
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}

func (c *Client) AddPokemon(pokemon Pokemon) error {

	My_pokedex[pokemon.Name] = pokemon
	return nil

}

func (c *Client) InspectPokemon(pokemon string) (Pokemon, error) {

	p, found := My_pokedex[pokemon]
	if found {
		return p, nil
	}
	return Pokemon{}, errors.New("Pokemon not found")

}

func (c *Client) GetPokedex() (map[string]Pokemon, error) {

	return My_pokedex, nil

}
