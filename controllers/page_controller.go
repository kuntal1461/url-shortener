package controllers

import (
	"log"
	"net/http"
)



func SetupRoutes()  {
	log.Printf("calling the controller ")
	http.HandleFunc("/",homeHandler)
	
}

func homeHandler(w http.ResponseWriter , r *http.Request)  {

	if (r.Method != "GET"){
		http.Error(w,"Method not aloowed",http.StatusMethodNotAllowed)
	}
	log.Printf("calling the indexhtml")
	http.ServeFile(w,r,"../static/index.html")
	log.Printf("Requested path: %s", r.URL.Path)

	
}