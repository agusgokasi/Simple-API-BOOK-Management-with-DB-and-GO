package model

type Book struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	Desc   string `json:"desc" db:"description"`
}
