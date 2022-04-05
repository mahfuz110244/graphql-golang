package model

// Author contains information about one author
type Author struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
