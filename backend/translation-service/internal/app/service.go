// internal/app/service.go
package app

import (
	"errors"

	out "github.com/q10357/language-app/pkg/client"
)

// Service provides translation services
type Service struct {
	client *out.Client
}

// NewService creates a new Service
func NewService(client *out.Client) *Service {
	return &Service{client: client}
}

// GetTranslation returns the translation for the given parameters
func (s *Service) GetTranslation(sourceLang, targetLang, word, pos string) (string, error) {
	translation, err := s.client.Translate(sourceLang, targetLang, word)
	if err != nil {
		return "", errors.New("translation not found")
	}
	return translation, nil
}
