package services

import "url-shortener/models"


type URLService interface{
	CreateURL(url *models.URL) error
}