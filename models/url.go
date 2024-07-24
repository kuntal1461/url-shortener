package models

import "time"





type URL struct{

	ID 		int  		`db:"id"`
	OriginalURL string 	`db:"original"`
	ShortURL string  	`db:"shortened"`
	CreatedAt time.Time  `db:"created_at"`
	ExpriesAt *time.Time  `db:"expires_at"`
	Count      int        `db:"count"`
}