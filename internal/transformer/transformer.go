package transformer

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInvalidInput = errors.New("invalid input data")
	ErrEmptyPayload = errors.New("empty payload")
)

// Post represents the structure of the input JSON
type Post struct {
	Slideshow Slideshow `json:"slideshow"`
}

// Validate checks if the Post structure is valid
func (p *Post) Validate() error {
	if p.Slideshow.Title == "" {
		return fmt.Errorf("%w: missing slideshow title", ErrInvalidInput)
	}
	return nil
}

type Slideshow struct {
	Author string `json:"author"`
	Date   string `json:"date"`
	Slides []struct {
		Title string   `json:"title"`
		Type  string   `json:"type"`
		Items []string `json:"items"`
	} `json:"slides"`
	Title string `json:"title"`
}

// TransformedPost represents the transformed data structure
type TransformedPost struct {
	Title      string `json:"title"`
	SlideCount int    `json:"slide_count"`
}

// TransformData converts the input JSON to a new structure
func TransformData(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, ErrEmptyPayload
	}

	// Parse the input JSON
	var post Post
	if err := json.Unmarshal(data, &post); err != nil {
		return nil, fmt.Errorf("error unmarshaling input data: %v", err)
	}

	if err := post.Validate(); err != nil {
		return nil, err
	}

	// Create transformed data
	transformedPost := TransformedPost{
		Title:      post.Slideshow.Title,
		SlideCount: len(post.Slideshow.Slides),
	}

	// Convert transformed data back to JSON
	transformedData, err := json.Marshal(transformedPost)
	if err != nil {
		return nil, fmt.Errorf("error marshaling transformed data: %v", err)
	}

	return transformedData, nil
}
