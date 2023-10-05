package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/noteTakingApplication/api"
)



func main() {
	
	router := api.AllRouters()

	
	// Start the server on localhost:8000
	fmt.Println("Server listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
