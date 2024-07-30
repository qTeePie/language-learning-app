package client

// ClientInterface defines the methods for the translation client
type ClientInterface interface {
	Translate(sourceLang, targetLang, text string) (string, error)
}
