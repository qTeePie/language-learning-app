package domain

// Translation represents a translation from one language to another
type TranslationResult struct {
	SourceLang     string `json:"source_lang"`
	TargetLang     string `json:"target_lang"`
	Text           string `json:"text"`
	TranslatedText string `json:"translated_text"`
}

type APIResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	Language string  `json:"language"`
	Senses   []Sense `json:"senses"`
}

type Sense struct {
	Definition   string                       `json:"definition"`
	Translations map[string]TranslationDetail `json:"translations"`
}

type TranslationDetail struct {
	Text string `json:"text"`
}
