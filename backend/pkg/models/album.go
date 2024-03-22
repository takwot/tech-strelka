package models


type Album struct {
	Id 			int 		`json:"id"`
	Name 		string 		`json:"name`
	Author 		int 		`json:"author"`
	Photos 		[]int 		`json:"photos"`
}