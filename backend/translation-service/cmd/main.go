package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	app "github.com/q10357/language-app/internal"
	out "github.com/q10357/language-app/pkg/client"
)

func main() {

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		if !os.IsNotExist(errLoadEnv) {
			log.Fatal("Error loading .env file", errLoadEnv)
		}
	}

	// Read the necessary environment variables
	baseURL := os.Getenv("TRANSLATION_API_BASE_URL")
	apiKey := os.Getenv("TRANSLATION_API_KEY")

	r := gin.Default()

	// Initialize client, service, and handler
	client := out.NewClient(baseURL, apiKey)
	service := app.NewService(client)
	handler := app.NewHandler(service)

	// Set up the route for translation
	r.GET("/translate", handler.HandleTranslation)

	// Run the server on port 5000
	err := r.Run(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
