// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type Book struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Price   int     `json:"price"`
	IsbnNo  string  `json:"isbn_no"`
	Authors *Author `json:"authors"`
}

type Books struct {
	Books []*Book `json:"books"`
}

type NewAuthor struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type NewBook struct {
	Title  string     `json:"title"`
	Price  int        `json:"price"`
	IsbnNo string     `json:"isbn_no"`
	Author *NewAuthor `json:"author"`
}
