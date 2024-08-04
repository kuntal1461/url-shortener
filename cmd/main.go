package main

import (
	"log"
	"net/http"
	"url-shortener/configs"
	"url-shortener/controllers"
	"url-shortener/db"
)


func main()  {

	cfg , err := configs.LoadConfig()
	if err !=nil {
		log.Fatalf("Failed to load configuration: %v", err)
		
	}
	if err := db.SetUpDatabase(); err !=nil{
		log.Fatal("Failed to connect with the Db",err)
	}

	//urlRepo := repositories.NewUrlRepository()
	//urlService := servicesimpl.NewUrlService(urlRepo)
	
	controllers.SetupRoutes()
	log.Printf("Server starting on %s\n", cfg.ServerPort)
    if err := http.ListenAndServe(cfg.ServerPort, nil); err != nil {
        log.Fatalf("Failed to start the server: %v", err)
    }

}
