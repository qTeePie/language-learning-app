package json

// TranslationResult represents a translation from one language to another
type TranslationResult struct {
	SourceLang string   `json:"source_lang"`
	TargetLang string   `json:"target_lang"`
	Headword   Headword `json:"headword"`
	Sense      Sense    `json:"sense"`
}

type APIResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	ID       string   `json:"id"`
	Source   string   `json:"source"`
	Language string   `json:"language"`
	Headword Headword `json:"headword"`
	Senses   []Sense  `json:"senses"`
}

type Headword struct {
	Text          string        `json:"text"`
	Pronunciation Pronunciation `json:"pronunciation"`
	Pos           string        `json:"pos"`
}

type Pronunciation struct {
	Value string `json:"value"`
}

type Sense struct {
	ID           string                       `json:"id"`
	Definition   string                       `json:"definition"`
	Translations map[string]TranslationDetail `json:"translations"`
}

type TranslationDetail struct {
	TextItems []TextItem `json:"options"`
}

type TextItem struct {
	Text   string `json:"text"`
	Gender string `json:"gender,omitempty"`
}
