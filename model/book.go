package model

type Book struct {
	ID          int    `json:"id"`
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
