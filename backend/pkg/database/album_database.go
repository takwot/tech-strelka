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

func (r *AlbumPostgres) DeleteAllAlbum(id int) error {
	query := fmt.Sprintf("DELETE  FROM %s WHERE id=$1", albumTable)

	_, err := r.db.Exec(query)

	return err
}

func (r *AlbumPostgres) UpdateAlbum(id int, update_id []int) (models.Album, error) {
	var album models.Album

	prev_list := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", albumTable)
	list = prev_list + update_id

	row := r.db.QueryRow(list, id)

	if err := row.Scan(&album); err != nil {
		return album, err
	}

	return album, nil
}

func (r *AlbumPostgres) RenameAlbum(id int, newName string) (models.Album, error) {
	var album models.Album

	query := fmt.Sprintf("UPDATE %s SET name = $2 WHERE id = $1", albumTable)

	row := r.db.QueryRow(query, update_id[0], id)

	if err := row.Scan(&album); err != nil {
		return album, err
	}

	return album, nil
}
