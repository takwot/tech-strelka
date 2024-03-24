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

func (s *AlbumService) GetAlbum(id int) (models.Album, error) {
	return s.repo.Album.GetAlbum(id)
}

func (s *AlbumService) UpdateAlbum(id int, newPhotoIDs []int) error {
	return s.repo.Album.UpdateAlbum(id, newPhotoIDs)
}

func (s *AlbumService) DeleteAlbum(id int) error {
	return s.repo.Album.DeleteAlbum(id)
}

func (s *AlbumService) RenameAlbum(id int, newName string) ( error) {
	return s.repo.Album.RenameAlbum(id, newName)
}
