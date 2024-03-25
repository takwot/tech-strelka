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
	GetAlbum(id string) (models.Album, error)
	DeleteAlbum(id int) error
	UpdateAlbum(albumID int, newPhotoIDs []int) error
	RenameAlbum(id int, newName string) error
}

type Photo interface {
	CreatePhoto(photo models.Photo) (int, error)
	DeletePhoto(id int) error
	GetPhoto(id int) (models.Photo, error)
}

type Repository struct {
	Auth
	Album
	Photo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:  NewAuthPostgres(db),
		Photo: NewPhotoPostgres(db),
		Album: NewAlbumPostgres(db),
	}
}
