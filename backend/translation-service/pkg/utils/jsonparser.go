package jsonparser

/*
import (
	"encoding/json"
	"fmt"
	"strings"

	output "github.com/q10357/language-app/domain"
)

func ParseTranslation(jsonData []byte, targetLang string) ([]translate.TranslationResult, error) {
	var apiResponse output.APIResponse
	if err := json.Unmarshal(jsonData, &apiResponse); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON %w", err)
	}

	var translations []output.TranslationResult

	for _, result := range apiResponse.Results {
		for _, sense := range result.Senses {
			translations, exists := sense.Translations[targetLang]
			if !exists {
				return nil, fmt.Errorf("Translation not found: %s(%s) to %s", strings.ToUpper())
			}
		}
	}
}
*/
