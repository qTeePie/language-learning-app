package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/q10357/AuthWGo/internal/app"
)

func main() {
	errLoadEnv := godotenv.Load()

	if errLoadEnv != nil {
		if !os.IsNotExist(errLoadEnv) {
			log.Fatal("Error loading .env file", errLoadEnv)
		}
	}

	fmt.Println("Server starting...")
	mux := http.NewServeMux()

	authMux := http.NewServeMux()
	authMux.HandleFunc("/signup", app.SignupHandler)
	authMux.HandleFunc("/signin", app.SigninHandler)
	authMux.HandleFunc("/validate", app.ValidateHandler)

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	errServer := server.ListenAndServe()
	if errServer != nil {
		fmt.Printf("Error Booting the Server\nError: %s\n", errServer)
	}
}
