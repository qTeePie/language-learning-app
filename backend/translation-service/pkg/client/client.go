package client

import (
	"fmt"
	"net/http"

	"github.com/q10357/language-app/pkg/utils/json"
)

// Client handles communication with the external translation API.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Client.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// Translate sends a translation request to the external API and returns the translation.
func (c *Client) Translate(sourceLang, targetLang, text string) (string, error) {
	fmt.Println("Stepping into client...")
	jsonData := []byte(dummyJSONData)

	translationObj, err := json.ParseTranslationObj(jsonData, targetLang)

	if err != nil {
		return "", fmt.Errorf("failed to parse translation: %w", err)
	}

	for index, obj := range translationObj {
		fmt.Printf("Result %d: %s\n", index, obj.Sense)
	}

	return "translated", nil

}

var dummyJSONData = `{
	"n_results": 1,
	"page_number": 1,
	"results_per_page": 10,
	"n_pages": 1,
	"available_n_pages": 1,
	"results": [
		{
			"id": "PW68729e22a9d0",
			"source": "password",
			"language": "en",
			"headword": {
				"text": "like",
				"pronunciation": {
					"value": "laik"
				},
				"pos": "verb"
			},
			"senses": [
				{
					"id": "PS73027db1b7a3",
					"definition": "to be pleased with; to find pleasant or agreeable",
					"translations": {
						"af": {
							"text": "hou van"
						},
						"ar": {
							"text": "أَحَبّ، هَوَى"
						},
						"az": {
							"text": "xoş gəlmək, , sevmək, istəmək"
						},
						"bg": {
							"text": "харесвам"
						},
						"br": {
							"text": "gostar"
						},
						"ca": {
							"text": "agradar"
						},
						"cs": {
							"text": "mít rád; líbit se"
						},
						"da": {
							"text": "(kunne) lide; bryde sig om"
						},
						"de": {
							"text": "mögen"
						},
						"el": {
							"text": "μου αρέσει"
						},
						"es": {
							"text": "gustar"
						},
						"et": {
							"text": "meeldima"
						},
						"fa": {
							"text": "خوش آمدن از"
						},
						"fi": {
							"text": "pitää"
						},
						"fr": {
							"text": "aimer"
						},
						"fy": {
							"text": "aardich/moai fine"
						},
						"he": {
							"text": "לֶאֱהוֹב, לְחַבֵּב"
						},
						"hi": {
							"text": "पसंद करना"
						},
						"hr": {
							"text": "sviđati se"
						},
						"hu": {
							"text": "tetszik vkinek vmi"
						},
						"id": {
							"text": "menyukai"
						},
						"is": {
							"text": "líka"
						},
						"it": [
							{ 
								"text": "probabile"
							},
							{ 
								"text": "text2"
							}
						],
						"ja": {
							"text": "好きだ"
						},
						"ko": {
							"text": "좋아하다"
						},
						"lt": {
							"text": "mėgti, kam patikti"
						},
						"lv": {
							"text": "patikt"
						},
						"ms": {
							"text": "suka"
						},
						"nl": {
							"text": "leuk vinden"
						},
						"no": {
							"text": "like, være glad i"
						},
						"pl": {
							"text": "lubić, podobać się"
						},
						"prs": {
							"text": "خوش آمدن از"
						},
						"ps": {
							"text": "د خوښى احساس كول"
						},
						"pt": {
							"text": "gostar"
						},
						"ro": {
							"text": "a-i plăcea"
						},
						"ru": {
							"text": "нравиться"
						},
						"sk": {
							"text": "mať rád; páčiť sa"
						},
						"sl": {
							"text": "imeti rad, biti všeč"
						},
						"sr": {
							"text": "sviđati se"
						},
						"sv": {
							"text": "tycka om, gilla"
						},
						"th": {
							"text": "ชอบ"
						},
						"tr": {
							"text": "sevmek, beğenmek"
						},
						"tw": {
							"text": "喜歡"
						},
						"uk": {
							"text": "подобатися, любити"
						},
						"ur": {
							"text": "کسی کو پسند کرنا"
						},
						"vi": {
							"text": "thích"
						},
						"zh": {
							"text": "喜欢"
						}
					},
					"examples": [
						{
							"text": "I like him very much."
						},
						{
							"text": "I like the way you've decorated this room."
						}
					]
				},
				{
					"id": "PSb158f035ffd6",
					"definition": "to enjoy",
					"translations": {
						"af": {
							"text": "hou van"
						},
						"ar": {
							"text": "اِسْتَمْتَع بِـ"
						},
						"az": {
							"text": "xoşuna gəlmək, zövq almaq"
						},
						"bg": {
							"text": "изпитвам удоволстие от"
						},
						"br": {
							"text": "gostar de"
						},
						"ca": {
							"text": "gaudir de"
						},
						"it": {
							"text": "piacere"
						}
					},
					"examples": [
						{
							"text": "I enjoy reading books."
						}
					]
				}
			]
		}
	]
}`
