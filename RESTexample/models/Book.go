package models

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

