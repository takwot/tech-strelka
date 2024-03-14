package database

import "github.com/jmoiron/sqlx"

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB)  *AuthPostgres{
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser() (int, error){
	// Creating user...


	return 0, nil
}