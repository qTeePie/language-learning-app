package json

import (
	"encoding/json"
	"fmt"
	"log"
)

// UnmarshalJSON custom unmarshaler for TranslationDetail to handle different JSON structures
func (t *TranslationDetail) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	// Attempt to unmarshal the raw text field without returning on error
	json.Unmarshal(data, &raw)

	log.Printf("raw text data: %s\n", raw.Text)

	t.Text = make(map[string][]TextItem)

	// Try to unmarshal as map[string][]TextItem (array of TextItems case)
	var arrayTexts map[string][]TextItem
	if err := json.Unmarshal(raw.Text, &arrayTexts); err == nil {
		log.Println("Unmarshalled as array of TextItems")
		t.Text = arrayTexts
		return nil
	} else {
		log.Printf("Failed to unmarshal as array of TextItems: %v\n", err)
	}

	// Try to unmarshal as map[string]string (single TextItem case)
	var singleTexts map[string]string
	if err := json.Unmarshal(raw.Text, &singleTexts); err == nil {
		log.Println("Unmarshalled as single TextItem")
		for lang, text := range singleTexts {
			t.Text[lang] = []TextItem{{Text: text}}
		}
		return nil
	} else {
		log.Printf("Failed to unmarshal as single TextItem: %v\n", err)
	}

	// Try to unmarshal as a single text value directly
	var singleTextItem map[string]string
	if err := json.Unmarshal(raw.Text, &singleTextItem); err == nil {
		log.Println("Unmarshalled as single text value")
		for lang, text := range singleTextItem {
			t.Text[lang] = []TextItem{{Text: text}}
		}
		return nil
	} else {
		log.Printf("Failed to unmarshal as single text value: %v\n", err)
	}

	return fmt.Errorf("invalid format for TranslationDetail.Text")
}
