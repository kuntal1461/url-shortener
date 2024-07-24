package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)


var DB *sqlx.DB

func setUpDataBase()  error {
	var err error
	// Update the connection string below with your actual database credentials
    dsn := "user=super_shortener dbname=url-shortener sslmode=disable password=Admin"
	DB,err :=sqlx.Connect("postgres",dsn)

	if err != nil {
		return err
		
	}
	return nil
	
}