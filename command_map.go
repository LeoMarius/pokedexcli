package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var ofset int = 0

func commandMap() error {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", ofset)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	data := MapResults{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	ofset += 20

	// fmt.Println(data)

	// fmt.Printf("%s", body)

	return nil
}

func commandMapBack() error {
	if ofset > 40 {
		// fmt.Println(ofset)
		ofset -= 40
	} else {
		fmt.Println()
		fmt.Println("you're on the first page")
		fmt.Println()
		ofset = 0
	}
	commandMap()

	// ofset -= 20

	return nil
}

type LocationArea struct {
	Name string `json:"name"` // key will be "name"
	Url  string `json:"url"`  // key will be "url"
}

type MapResults struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous interface{}    `json:"previous"`
	Results  []LocationArea `json:"results"`
}
