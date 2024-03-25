package models

type Photo struct {
	Id     int      `json:"id"`
	Author int      `json:"author_id"`
	Name   string   `json:"photo_name"`
	tags   []string `json:"tags"`
}
