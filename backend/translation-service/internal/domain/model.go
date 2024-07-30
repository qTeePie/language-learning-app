package domain

// Translation represents a translation from one language to another
type Translation struct {
	SourceLang     string `json:"source_lang"`
	TargetLang     string `json:"target_lang"`
	Text           string `json:"text"`
	TranslatedText string `json:"translated_text"`
}
