package client

// ClientInterface defines the methods for the translation client
type ClientInterface interface {
	FetchTranslation(sourceLang, targetLang, text string) (string, error)
}
