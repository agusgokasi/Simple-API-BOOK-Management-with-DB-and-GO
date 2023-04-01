package model

type Book struct {
	ID     int64  `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	Desc   string `json:"desc" db:"description"`
}
