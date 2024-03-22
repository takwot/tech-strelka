package models

type User struct {
	Id       int                   `json:"id"`
	Name     string                `json:"name"`
	Password string                `json:"password"`
	Email    string                `json:"email"`
	Avatar   bool 					`json:"avatar"`
}
