package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client handles communication with the external translation API.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Client.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// TranslationRequest represents the request payload for the translation API.
type TranslationRequest struct {
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
	Text       string `json:"text"`
}

// TranslationResponse represents the response from the translation API.
type TranslationResponse struct {
	TranslatedText string `json:"translated_text"`
}

// Translate sends a translation request to the external API and returns the translation.
func (c *Client) Translate(sourceLang, targetLang, text string) (string, error) {
	requestPayload := TranslationRequest{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		Text:       text,
	}

	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request payload: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/translate", c.baseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 response: %s", resp.Status)
	}

	var translationResponse TranslationResponse
	err = json.NewDecoder(resp.Body).Decode(&translationResponse)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return translationResponse.TranslatedText, nil
}
