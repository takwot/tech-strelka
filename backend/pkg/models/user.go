package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
