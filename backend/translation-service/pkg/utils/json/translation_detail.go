package json

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON custom unmarshaler for TranslationDetail to handle different JSON structures
func (t *TranslationDetail) UnmarshalJSON(data []byte) error {
	// Try unmarshaling into a single TextItem first
	var single TextItem
	if err := json.Unmarshal(data, &single); err == nil {
		t.TextItems = []TextItem{single}
		return nil
	}

	// Try unmarshaling into a slice of TextItems
	var slice []TextItem
	if err := json.Unmarshal(data, &slice); err == nil {
		t.TextItems = slice
		return nil
	}

	fmt.Printf("TextItem in custom unmashal function: %s\n", single)
	// Return an error if neither worked
	return fmt.Errorf("data does not match any format")
}
