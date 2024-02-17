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

	// Extended Vigenere Cipher
	apiRouter.HandleFunc("/extended-vigenere", middleware.HandleExtendedVigenere).Methods("POST")

	// Playfair Cipher
	apiRouter.HandleFunc("/playfair", middleware.HandlePlayfair).Methods("POST")

	// Affine Cipher
	apiRouter.HandleFunc("/affine", middleware.HandleAffine).Methods("POST")

	return router
}
