package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/noteTakingApplication/api"
)

// var users map[string]User

// var accountsLogin []Login
// var (
// 	accounts    map[string]Signup
// 	userSID     map[string]string
// 	notes       []Note
// 	noteCounter uint32
// )

func main() {
	// accounts = make(map[string]Signup) // Initialize the map
	// userSID = make(map[string]string)  // Initialize the map
	// notes = make([]Note, 0)
	// noteCounter = 0
	// // Initialize the Gorilla Mux router
	// router := mux.NewRouter()
	router := api.AllRouters()

	// // Define the API endpoints
	// router.HandleFunc("/notes", getNotes).Methods("GET")
	// router.HandleFunc("/signup", createSignup).Methods("POST")
	// router.HandleFunc("/login", createLogin).Methods("POST")
	// router.HandleFunc("/notes", createNotes).Methods("POST")
	// router.HandleFunc("/notes", deleteResource).Methods("DELETE")

	// Start the server on localhost:8000
	fmt.Println("Server listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
