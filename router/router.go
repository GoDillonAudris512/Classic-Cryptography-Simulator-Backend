package router

import (
	"github.com/gorilla/mux"

	"classic-crypt/middleware"
)

func Router() *mux.Router {
	// General Endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api", middleware.MainHandler).Methods("GET")

	// Functional Endpoint
	// apiRouter := router.PathPrefix("/api").Subrouter()

	return router
}
