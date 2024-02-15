package main

import (
	"fmt"
	"github.com/joho/godotenv"
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

	var port = EnvPortOr("8080")
	fmt.Println("Starting server...")
	fmt.Println("Listening from http://localhost" + port + "/api")
	log.Fatal(http.ListenAndServe(port, router))
}

func EnvPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}

	return ":" + port
}
