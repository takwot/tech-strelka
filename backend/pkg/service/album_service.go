package service

import (
	"github.com/takwot/tech-strelka.git/pkg/database"
	"github.com/takwot/tech-strelka.git/pkg/models"
)


type AlbumService struct {
	repo database.Repository
}


func NewAlbumService(repo database.Repository) *AlbumService {
	return &AlbumService{repo: repo}
}



func (s *AlbumService) CreateAlbum(album models.Album) (int, error) {
	return s.repo.Album.CreateAlbum(album)
}

func (s *AlbumService) GetAllAlbum() ([]models.Album, error) {
	return s.repo.Album.GetAllAlbum()
}