package main

import (
	"log"

	"github.com/cmilhench/go-starter-transformer/internal/fetcher"
	"github.com/cmilhench/go-starter-transformer/internal/sender"
	"github.com/cmilhench/go-starter-transformer/internal/transformer"
)

func main() {
	// Example source API
	sourceURL := "https://httpbin.org/json"

	// Example destination API
	destURL := "https://httpbin.org/status/501"

	// Fetch data from source API
	rawData, err := fetcher.FetchData(sourceURL)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	log.Println("Successfully fetched data!")

	// Transform the fetched data
	transformedData, err := transformer.TransformData(rawData)
	if err != nil {
		log.Fatalf("Failed to transform data: %v", err)
	}
	log.Println("Successfully transformed data!")

	// Send transformed data to destination API
	err = sender.SendData(destURL, transformedData)
	if err != nil {
		log.Fatalf("Failed to send data: %v", err)
	}
	log.Println("Successfully sent data!")
}
