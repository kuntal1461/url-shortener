package main

import (
	"log"
	"url-shortener/controllers"
	"url-shortener/db"
	
)


func main()  {
	if err := db.SetUpDatabase(); err !=nil{
		log.Fatal("Failed to connect with the Db",err)
	}

	//urlRepo := repositories.NewUrlRepository()
	//urlService := servicesimpl.NewUrlService(urlRepo)
	
	controllers.SetupRoutes()
	

}
