package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"url-shortener/db"
	"url-shortener/models"

	"github.com/segmentio/ksuid"
)

// CreateShortURL handles the creation of short URLs
func CreateShortURL(w http.ResponseWriter, r *http.Request) {

	var request struct {
		OriginalUrl string `json:"original_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error decoding request:", err)
		return
	}

	id := ksuid.New().String()

	shortenedURL := fmt.Sprintf("http://localhost:8082/%s", id)

	url := models.URL{
		ID:         id,
		Original:   request.OriginalUrl,
		Shortened:  shortenedURL,
		Created_at: time.Now(),
	}

	_, err := db.GetDB().Exec("INSERT INTO urls (id, original, shortened, created_at) VALUES ($1, $2, $3, $4)", url.ID, url.Original, url.Shortened, url.Created_at)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error inserting URL into the database:", err)
		return

	}
	log.Println("Successfully created short URL:", shortenedURL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)

}
