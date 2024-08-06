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
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := db.SetUpDatabase(); err != nil {
		log.Fatal("Failed to connect with the Db", err)
	}

	controllers.SetupRoutes()

	// Use the PORT environment variable if available
	port := cfg.ServerPort
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	log.Printf("Server starting on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
