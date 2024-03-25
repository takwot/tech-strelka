package service

import (
	"github.com/takwot/tech-strelka.git/pkg/database"
	"github.com/takwot/tech-strelka.git/pkg/models"
)

type PhotoService struct {
	repo database.Repository
}

func NewPhotoService(repo database.Repository) *PhotoService {
	return &PhotoService{repo: repo}
}

func (s *PhotoService) CreatePhoto(photo models.Photo) (int, error) {
	return s.repo.Photo.CreatePhoto(photo)
}
func (s *PhotoService) GetPhoto(id int) (models.Photo, error) {
	return s.repo.Photo.GetPhoto(id)
}

func (s *PhotoService) DeletePhoto(id int) error {
	return s.repo.Photo.DeletePhoto(id)
}
