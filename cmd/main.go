package main

import (
	"log"
	"net/http"
	"os"
	"url-shortener/configs"
	"url-shortener/controllers"
	"url-shortener/db"
)

func main() {
	// Load configuration
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up the database connection
	if err := db.SetUpDatabase(); err != nil {
		log.Fatalf("Failed to connect with the database: %v", err)
	}

	// Set up routes
	controllers.SetupRoutes()

	// Use the PORT environment variable if available
	port := cfg.ServerPort
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	// Log the starting of the server
	log.Printf("Server starting on %s\n", port)

	// Start the server with logging middleware
	if err := http.ListenAndServe(port, loggingMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

// loggingMiddleware logs each request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
