package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/takwot/tech-strelka.git/pkg/models"
)



type AlbumPostgres struct {
	db *sqlx.DB
}


func NewAlbumPostgres(db *sqlx.DB) *AlbumPostgres {
	return &AlbumPostgres{db: db}
}


func (r *AlbumPostgres) CreateAlbum(album models.Album) (int, error) {
	
	var id int

	query := fmt.Sprintf("INTERT INTO %s (name, author) VALUES ($1, $2) RETURNING id", albumTable)
	
	row := r.db.QueryRow(query, album.Name, album.Author)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AlbumPostgres) GetAllAlbum() ([]models.Album, error) {
	var albums []models.Album

	query := fmt.Sprintf("SELECT * FROM %s", albumTable)

	row := r.db.QueryRow(query)

	if err := row.Scan(&albums); err != nil {
		return nil, err
	}
	
	return albums, nil
}