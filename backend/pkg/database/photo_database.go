package database

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/takwot/tech-strelka.git/pkg/models"
)

type PhotoPostgres struct {
	db *sqlx.DB
}

func NewPhotoPostgres(db *sqlx.DB) *PhotoPostgres {
	return &PhotoPostgres{db: db}
}

func (r *PhotoPostgres) CreatePhoto(photo models.Photo, c *gin.Context) (int, error, string) {

	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, author) VALUES ($1, $2) RETURNING id", photoTable)
	tags := c.PostForm("tags")
	row := r.db.QueryRow(query, photo.Name, photo.Author)

	if err := row.Scan(&id); err != nil {
		return 0, err, tags
	}

	return id, nil, tags
}

func (r *PhotoPostgres) GetPhoto(id int, c *gin.Context) (models.Photo, error, string) {
	var photo models.Photo
	tags := c.PostForm("tags")
	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1", photoTable)
	err := r.db.Get(&photo, query, id)

	return photo, err, tags
}

func (r *AlbumPostgres) GetAllPhotos() ([]models.Photo, error) {
	var photos []models.Photo

	query := fmt.Sprintf("SELECT * FROM %s", photoTable)

	row := r.db.QueryRow(query)

	if err := row.Scan(&photos); err != nil {
		return nil, err
	}

	return photos, nil
}

func (r *PhotoPostgres) DeletePhoto(id int) error {
	query := fmt.Sprintf("DELETE  FROM %s WHERE id=$1", photoTable)

	_, err := r.db.Exec(query)

	return err
}
