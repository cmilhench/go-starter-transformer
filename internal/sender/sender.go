package sender

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

// SendData sends the transformed data to a specified URL
func SendData(url string, data []byte) error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new HTTP request with the transformed data
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client with custom settings
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
