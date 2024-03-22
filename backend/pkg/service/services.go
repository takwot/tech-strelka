package service

import (
	"github.com/takwot/tech-strelka.git/pkg/database"
	"github.com/takwot/tech-strelka.git/pkg/models"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
	ParseToken(token string) (int, error)
	GenerateToken(username string, password string) (string, error)
}

type Album interface {
	CreateAlbum(album models.Album) (int, error)
	GetAllAlbum() ([]models.Album, error)
}

type Service struct {
	Auth
	Album
}

func NewServices(repo *database.Repository) *Service {
	return &Service{
		Auth:  NewAuthService(*repo),
		Album: NewAlbumService(*repo),
	}
}
