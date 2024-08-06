package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"os"
)

var DB *sqlx.DB

// SetUpDatabase connects to the PostgreSQL database using environment variables.
func SetUpDatabase() error {
	// Use the environment variable for the connection string
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		log.Fatal("Environment variable POSTGRES_URL is not set")
	}

	log.Printf("Connecting to the database with DSN: %s", dsn)

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Printf("Error connecting to the database: %v", err)
		return err
	}

	log.Println("Successfully connected to the database")
	return nil
}

// HealthCheck performs a simple ping to the database to ensure it's accessible.
func HealthCheck() error {
	return DB.Ping()
}
