package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define routes
	// router.HandleFunc("/users", getUsers).Methods("GET")

	log.Println("Starting API on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}