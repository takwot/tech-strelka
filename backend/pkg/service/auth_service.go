package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/takwot/tech-strelka.git/pkg/database"
)

type AuthService struct {
	repo database.Repository
}

const (
	salt = "hjqrhjqw124617ajfhajs"
)

func NewAuthService(repo database.Repository) *AuthService{
	return &AuthService{repo: repo}
}


func (s *AuthService) CreateUser() (int, error){
	
	// Logic....

	return 0, nil
}


func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}