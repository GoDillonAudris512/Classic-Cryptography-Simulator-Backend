package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"classic-crypt/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading .env file")
	}

	router := router.Router()

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	corsHandler := corsOptions.Handler(router)

	var port = EnvPortOr("8080")
	fmt.Println("Starting server...")
	fmt.Println("Listening from http://localhost" + port + "/api")
	log.Fatal(http.ListenAndServe(port, corsHandler))
}

func EnvPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}

	return ":" + port
}
