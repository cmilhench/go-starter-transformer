package fetcher

import (
	"context"
	"io"
	"net/http"
	"time"
)

// FetchData retrieves data from a given URL with a timeout
func FetchData(url string) ([]byte, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Create an HTTP client with custom settings
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
