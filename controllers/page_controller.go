package controllers

import (
	"log"
	"net/http"
	"path/filepath"
)

// SetupRoutes sets up the HTTP routes for the application.
func SetupRoutes() {
	log.Printf("Setting up routes")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthCheckHandler)
}

// homeHandler serves the index.html file.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Serving the index.html for path: %s", r.URL.Path)
	// Use absolute path for the file to ensure it works in Vercel
	absPath, err := filepath.Abs("static/index.html")
	if err != nil {
		log.Printf("Error finding absolute path: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, absPath)
}

// healthCheckHandler is a simple endpoint to check if the server is running.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	log.Printf("Health check passed")
}
