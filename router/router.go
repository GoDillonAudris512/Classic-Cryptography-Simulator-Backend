package router

import (
	"github.com/gorilla/mux"

	"classic-crypt/middleware"
)

func Router() *mux.Router {
	// General Endpoints
	router := mux.NewRouter()
	router.HandleFunc("/api", middleware.MainHandler).Methods("GET")

	// Functional Endpoints
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Standard Vigenere Cipher
	apiRouter.HandleFunc("/vigenere", middleware.HandleVigenere).Methods("POST")

	// Auto-Key Vigenere Cipher
	apiRouter.HandleFunc("/auto-vigenere", middleware.HandleAutoVigenere).Methods("POST")
	
	return router
}
