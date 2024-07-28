import requests
from fastapi import HTTPException
from config import TRANSLATION_API_URL, DICTIONARY_API_URL

def fetch_translation(word: str) -> str:
    response = requests.get(TRANSLATION_API_URL.format(word=word))
    if response.status_code != 200:
        raise HTTPException(status_code=404, detail="Translation not found")
    translation_data = response.json()
    italian_word = translation_data['responseData']['translatedText']
    return italian_word

def fetch_definition(word: str) -> dict:
    response = requests.get(DICTIONARY_API_URL.format(word=word))
    if response.status_code != 200:
        raise HTTPException(status_code=404, detail="Definition not found")
    definition_data = response.json()
    return {
        "definitions": definition_data[0].get("meanings", []),
        "phonetics": definition_data[0].get("phonetics", []),
        "origin": definition_data[0].get("origin", ""),
        "license": definition_data[0].get("license", {}),
        "source_urls": definition_data[0].get("sourceUrls", [])
    }
