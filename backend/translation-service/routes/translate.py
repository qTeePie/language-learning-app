from fastapi import APIRouter
from services.translation_service import fetch_translation, fetch_definition

router = APIRouter()

@router.get("/translate/{word}")
async def translate(word: str):
    italian_word = fetch_translation(word)
    definition_info = fetch_definition(italian_word)
    result = {
        "word": word,
        "translation": italian_word,
        "definition": definition_info
    }
    print(result)  # Write out to console
    return result
