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

	query := fmt.Sprintf("INSERT INTO %s (name, author, photos) VALUES ($1, $2, $3) RETURNING id", albumTable)

	row := r.db.QueryRow(query, album.Name, album.Author, album.Photos)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AlbumPostgres) GetAlbum(id int) (models.Album, error) {
	var album models.Album
	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1", albumTable)
	err := r.db.Get(&album, query, id)

	return album, err
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

func (r *AlbumPostgres) DeleteAlbum(id int) error {
	query := fmt.Sprintf("DELETE  FROM %s WHERE id=$1", albumTable)

	_, err := r.db.Exec(query, id)

	return err
}

func (r *AlbumPostgres) UpdateAlbum(albumID int, newPhotoIDs []int) error {
	query := fmt.Sprintf("UPDATE %s SET photos = photos || $1 WHERE id = $2", albumTable)

	interfacePhotoIDs := make([]interface{}, len(newPhotoIDs))
	for i, id := range newPhotoIDs {
		interfacePhotoIDs[i] = id
	}

	_, err := r.db.Exec(query, interfacePhotoIDs, albumID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AlbumPostgres) RenameAlbum(id int, newName string) error {
	var album models.Album

	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", albumTable)

	row := r.db.QueryRow(query, newName, id)

	if err := row.Scan(&album); err != nil {
		return err
	}

	return nil
}
