package controllers

import "net/http"



func SetupRoutes()  {
	http.HandleFunc("/",homeHandler)
	
}

func homeHandler(w http.ResponseWriter , r *http.Request)  {

	if (r.Method != "GET"){
		http.Error(w,"Method not aloowed",http.StatusMethodNotAllowed)
	}
	http.ServeFile(w,r,"static.index.html")
	
}