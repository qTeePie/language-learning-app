// internal/app/service.go
package translate

import (
	"errors"
	"fmt"

	out "github.com/q10357/language-app/pkg/client"
	"github.com/q10357/language-app/pkg/utils/json"
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
func (s *Service) GetTranslation(sourceLang, targetLang, word, pos string) ([]json.TranslationResult, error) {
	translation, err := s.client.Translate(sourceLang, targetLang, word)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, errors.New("translation not found")
	}
	return translation, nil
}
