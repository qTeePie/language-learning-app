package utils

import (
	"fmt"
	"net/url"
)

// URLBuilder constructs URLs with query parameters
type URLBuilder struct {
	baseURL string
	params  url.Values
}

// NewURLBuilder initializes a new URLBuilder with a base URL
func NewURLBuilder(baseURL string) *URLBuilder {
	return &URLBuilder{
		baseURL: baseURL,
		params:  url.Values{},
	}
}

// SetQueryParam sets a query parameter
func (b *URLBuilder) SetQueryParam(params map[string]string) *URLBuilder {
	for key, value := range params {
		b.params.Set(key, value)
	}
	return b
}

// Build constructs the final URL
func (b *URLBuilder) Build() (string, error) {
	u, err := url.Parse(b.baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}
	u.RawQuery = b.params.Encode()
	return u.String(), nil
}
