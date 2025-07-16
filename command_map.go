package main

import (
	"fmt"
	"net/http"
	"io"
)

type Config struct {
	Next             *string
	Previous         *string
}

func commandMap() error {
	locationAreaEndpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=%d", 0, 20)
	res, err := http.Get(locationAreaEndpoint)
	if err != nil {
		return err
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Response: %s\n", body)

	res.Body.Close()
	return nil
}