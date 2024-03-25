package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/takwot/tech-strelka.git/pkg/models"
)

type PhotoPostgres struct {
	db *sqlx.DB
}

func NewPhotoPostgres(db *sqlx.DB) *PhotoPostgres {
	return &PhotoPostgres{db: db}
}

func (r *PhotoPostgres) CreatePhoto(photo models.Photo) (int, error) {

	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, author) VALUES ($1, $2) RETURNING id", photoTable)
	row := r.db.QueryRow(query, photo.Name, photo.Author)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PhotoPostgres) GetPhoto(id int) (models.Photo, error) {
	var photo models.Photo
	query := "SELECT * FROM " + photoTable + " WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&photo.Id, &photo.Name)

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *PhotoPostgres) DeletePhoto(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", photoTable)

	_, err := r.db.Exec(query)

	return err
}
