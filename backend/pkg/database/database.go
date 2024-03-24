package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/takwot/tech-strelka.git/pkg/models"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Album interface {
	CreateAlbum(album models.Album) (int, error)
	GetAllAlbum() ([]models.Album, error)
	GetAlbum(id int) (models.Album, error)
	DeleteAlbum(id int) error
	UpdateAlbum(albumID int, newPhotoIDs []int) error
	RenameAlbum(id int, newName string) (models.Album, error) 
}

type Repository struct {
	Auth
	Album
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
		Album: NewAlbumPostgres(db),
	}
}
