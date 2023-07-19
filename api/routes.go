package api

import (
	"github.com/gorilla/mux"
)

func AllRouters() *mux.Router {
	// Initialize the Gorilla Mux router
	router := mux.NewRouter()

	// Define the API endpoints
	router.HandleFunc("/notes", GetNotes).Methods("GET")
	router.HandleFunc("/signup", CreateSignup).Methods("POST")
	router.HandleFunc("/login", CreateLogin).Methods("POST")
	router.HandleFunc("/notes", CreateNotes).Methods("POST")
	router.HandleFunc("/notes", DeleteResource).Methods("DELETE")

	return router
}
