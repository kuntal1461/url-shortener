package servicesimpl

import (
	"url-shortener/models"
	"url-shortener/repositories"
	"url-shortener/services"
)




type urlServiceImpl struct{
	repo repositories.URLRepository
}

func NewUrlService(repo repositories.URLRepository) services.URLService{
	return &urlServiceImpl{repo: repo}
}
func(s *urlServiceImpl) CreateURL(url *models.URL) error{
	return s.repo.Create(url)
}