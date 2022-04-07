package model

type Book struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	IsbnNo   string  `json:"isbn_no"`
	AuthorID string  `json:"author_id"`
	Authors  *Author `json:"authors" gorm:"foreignkey:AuthorID"`
}

type Books struct {
	Books []*Book `json:"books"`
}
