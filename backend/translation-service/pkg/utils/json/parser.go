package json

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ParseTranslationObj(jsonData []byte, targetLang string) ([]TranslationResult, error) {
	var apiResponse APIResponse
	if err := json.Unmarshal(jsonData, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON : %w", err)
	}

	var translations []TranslationResult

	for _, result := range apiResponse.Results {
		for _, sense := range result.Senses {
			translationDetail, exists := sense.Translations[targetLang]
			if !exists {
				return nil, fmt.Errorf("translation not found: %s(%s) to %s", strings.ToUpper(result.Headword.Text), result.Language, targetLang)
			}

			translations = append(translations, TranslationResult{
				SourceLang: result.Language,
				TargetLang: targetLang,
				Headword:   result.Headword,
				Sense: Sense{
					ID:           sense.ID,
					Definition:   sense.Definition,
					Translations: map[string]TranslationDetail{targetLang: translationDetail},
				},
			})
		}
	}

	return translations, nil
}
