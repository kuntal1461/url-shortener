package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //PostgreSql driver
)

var Db *sql.DB // public (Exported)

// init initializes the database connection
func Init() {
	var err error


	// user name need to change accordingly 
	connStr := "user=super_shortener dbname=url-shortener sslmode=disable password=Admin"

	Db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Failed to open the db connetion", err)

	}
	err = Db.Ping()
	if err != nil {
		log.Fatal("failed to connect with the Db ", err)
	}
	log.Println("The Db connection established")

	// Run a simple query to verify the connection
	var version string
	err = Db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal("Failed to run query: ", err)
	}
	log.Println("PostgreSQL version: ", version)

}

func GetDB() *sql.DB {
	return Db
}
