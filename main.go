package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/db"
	"url-shortener/handlers"

	"github.com/gorilla/mux"
)


func main()  {
	fmt.Println("This is the new url shortener project")
	
	// Initialize the database connection
	db.Init()

	// Set up the router
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.CreateShortURL).Methods("POST")
	r.HandleFunc("/{id}", handlers.RedirectURL).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Start the server
	log.Fatal(http.ListenAndServe(":8082", r))
	


}
 