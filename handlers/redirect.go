package handlers

import (
	"database/sql"
	"net/http"
	"url-shortener/db"
	"url-shortener/models"
)

// RedirectURL handles the redirection to original URLs
func RedirectURL(w http.ResponseWriter, r *http.Request)  {
	id:=r.URL.Path[len("/"):]

	var url models.URL

	err :=db.GetDB().QueryRow("SELECT original FROM urls WHERE id = $1", id).Scan(&url.Original)

	if err== sql.ErrNoRows{
		http.NotFound(w,r)
		return
	}else if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url.Original, http.StatusFound)
	
}
