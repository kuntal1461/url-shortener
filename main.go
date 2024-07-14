package main

import (
	"fmt"
	"url-shortener/db"
)


func main()  {
	fmt.Println("This is the new url shortener project")
	
	// Initialize the database connection
	db.Init()

	database := db.GetDB();

	if database !=nil {
		fmt.Println("The connectin is working and verifed ")
	}else{
		fmt.Println("The connection is not working ")
	}
	


}
 