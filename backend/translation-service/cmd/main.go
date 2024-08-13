package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	app "github.com/q10357/language-app/translation-service/internal"
	out "github.com/q10357/language-app/translation-service/pkg/client"
)

func main() {

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil && !os.IsNotExist(errLoadEnv) {
		log.Fatal("Error loading .env file", errLoadEnv)
	}

	// Read the necessary environment variables
	baseURL := os.Getenv("TRANSLATION_API_BASE_URL")
	apiKeyHeaderName := os.Getenv("API_KEY_TRANSLATION_API_HEADER_KEY")
	apiKey := os.Getenv("API_KEY_TRANSLATION_API_HEADER_VALUE")

	// Initialize client (singleton), service, and handler
	client := out.NewClient(baseURL, apiKeyHeaderName, apiKey)
	service := app.NewService(client)
	handler := app.NewHandler(service)

	r := gin.Default()

	// Set up the route for translation
	r.GET("/translate", handler.HandleTranslation)

	// Run the server on port 5000
	err := r.Run(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
