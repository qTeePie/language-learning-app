package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	output "github.com/q10357/language-app/domain"
)

func ParseTranslationObj(jsonData []byte, targetLang string) ([]output.TranslationResult, error) {
	var apiResponse output.APIResponse
	if err := json.Unmarshal(jsonData, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON %w", err)
	}

	var translations []output.TranslationResult

	for _, result := range apiResponse.Results {
		for _, sense := range result.Senses {
			translationDetail, exists := sense.Translations[targetLang]
			if !exists {
				return nil, fmt.Errorf("translation not found: %s(%s) to %s", strings.ToUpper(result.Headword.Text), result.Language, targetLang)
			}
			translations = append(translations, output.TranslationResult{
				SourceLang: result.Language,
				TargetLang: targetLang,
				Text:       result.Headword.Text,
				Sense: output.Sense{
					ID:           sense.ID,
					Definition:   sense.Definition,
					Translations: map[string]output.TranslationDetail{targetLang: translationDetail},
				},
			})
		}
	}

	return translations, nil
}
