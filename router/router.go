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

	// Vigenere Encryption
	apiRouter.HandleFunc("/vigenere/encrypt", middleware.EncryptVigenere).Methods("POST")
	apiRouter.HandleFunc("/vigenere/decrypt", middleware.DecryptVigenere).Methods("POST")

	return router
}
