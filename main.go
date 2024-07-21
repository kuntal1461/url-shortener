package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/db"
	"url-shortener/handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This is the new url shortener project")

	// Initialize the database connection
	db.Init()

	// Set up the router
	r := mux.NewRouter()
	// Log the router state
	log.Printf("Router state before adding routes: %+v\n", r)

	r.HandleFunc("/shorten", handlers.CreateShortURL).Methods("POST")
	r.HandleFunc("/{id}", handlers.RedirectURL).Methods("GET")

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static/")) // Serve files from the 'static' directory
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	// Start the server
	log.Fatal(http.ListenAndServe(":8082", r))

}
