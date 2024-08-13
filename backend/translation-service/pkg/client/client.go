package client

import (
	"fmt"
	"io"
	"net/http"
	"sync" ////thread-safe initialization

	"github.com/q10357/language-app/translation-service/pkg/utils"
	"github.com/q10357/language-app/translation-service/pkg/utils/json"
)

// Client handles communication with the external translation API.
type Client struct {
	baseURL          string
	apiKey           string
	apiKeyHeaderName string
	httpClient       *http.Client
}

// Singleton instance
var (
	clientInstance *Client
	once           sync.Once
)

// NewClient creates a new Client.
func NewClient(baseURL, apiKeyHeaderName, apiKey string) *Client {
	once.Do(func() {

		clientInstance = &Client{
			baseURL:          baseURL,
			apiKey:           apiKey,
			apiKeyHeaderName: apiKeyHeaderName,
			httpClient:       &http.Client{},
		}
	})
	return clientInstance
}

// FetchTranslations sends a translation request to the external API and returns the translation.
func (c *Client) FetchTranslations(sourceLang, targetLang, text string, pos string) ([]json.TranslationResult, error) {
	// Create instance of URLBuilder with extern translation API url as base
	urlBuilder := utils.NewURLBuilder(c.baseURL)
	// Define params
	params := map[string]string{
		"text":     text,
		"language": sourceLang,
		"pos":      pos,
		"source":   "password",
	}

	// Build final URL, encoding the url+params to valid format
	translationURL, err := urlBuilder.SetQueryParam(params).Build()
	if err != nil {
		return nil, err
	}

	// Create HTTP request with API key as header
	req, err := http.NewRequest("GET", translationURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(c.apiKeyHeaderName, c.apiKey)
	// req.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // Ensure response body is closed

	// Read and print the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	translationObj, err := json.ParseTranslationObj(body, targetLang)

	if err != nil {
		return nil, err
	}

	for index, obj := range translationObj {
		fmt.Printf("Result %d: %s\n", index, obj.Sense)
	}

	return translationObj, nil

}
