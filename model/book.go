package model

type Book struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
	IsbnNo  string  `json:"isbn_no"`
	Authors *Author `json:"authors"`
}

type Books struct {
	Books []*Book `json:"books"`
}
