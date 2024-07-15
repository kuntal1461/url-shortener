package models

import "time"



type URL struct{
	ID string `json:"id"`
	Original string `json:"original"`
	Shortened string `json:"shortened"`
	Created_at time.Time `json:"created_at"`
}