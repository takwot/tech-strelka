package database

import "github.com/jmoiron/sqlx"

type Auth interface {
	CreateUser() (int, error)
}

type Repository struct {
	Auth
}


func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}