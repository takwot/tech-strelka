package models

type Album struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Author int    `json:"author" db:"author"`
	Photos []int  `json:"photos" db:"photos"`
}
