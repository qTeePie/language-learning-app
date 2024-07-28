from fastapi import FastAPI
from routes import translate

app = FastAPI()

app.include_router(translate.router)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="127.0.0.1", port=8000, reload=True)
