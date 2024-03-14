package service

import (
	"github.com/takwot/tech-strelka.git/pkg/database"
)

type Auth interface {
	CreateUser()(int, error)
}

type Service struct {
	Auth
}

func NewServices(repo *database.Repository) *Service{
	return &Service{
		Auth: NewAuthService(*repo),
	}
}