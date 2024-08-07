package repositories

import (
	"url-shortener/db"
	"url-shortener/models"
)


type URLRepository interface {
	Create(url *models.URL) error
}

type urlRepositoryImpl struct{}

// if the function is upercase then it will be public .
func NewUrlRepository()  URLRepository{
	return &urlRepositoryImpl{}
}

func (repo *urlRepositoryImpl) Create( url *models.URL)  error{
	const query = `INSERT INTO urls (original, shortened, created_at, expires_at, count) 
                   VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err :=db.DB.QueryRow(query,url.OriginalURL,url.ShortURL,url.CreatedAt,url.ExpriesAt,url.Count).Scan(url.ID)
	if err !=nil {
		return err
		
	}
	return nil
}