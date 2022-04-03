package repository

import (
	"log"

	db "graphql-golang/internal/pkg/db/mysql"
	"graphql-golang/schema"
)

//CreateAuthor create's author
func CreateAuthor(author schema.Author) (int64, error) {

	stmt, err := db.Db.Prepare("INSERT INTO Authors(Name,Biography) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	res, err := stmt.Exec(author.Name, author.Biography)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	defer stmt.Close()
	log.Println("Row inserted!!")
	return id, nil
}

//CreateBook creates new book
func CreateBook(book schema.Book) (int64, error) {
	stmt, err := db.Db.Prepare("insert into Books(Title,Price,IsbnNo,AuthorID) VALUES(?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(book.Title, book.Price, book.IsbnNo, book.Authors.ID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

//GetBooksByID returns books by respective id
func GetBooksByID(id *string) (*schema.Book, error) {
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Books.Price,Books.IsbnNo,Authors.ID,Authors.Name,Authors.Biography from Books inner join Authors where Books.AuthorID = Authors.ID and Books.ID = ? ;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	var bookID, title, isbn_no, authorID, name, biography string
	var price int
	if rows.Next() {
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}
	}

	book := &schema.Book{
		ID:     bookID,
		Title:  title,
		Price:  price,
		IsbnNo: isbn_no,
		Authors: &schema.Author{
			ID:        authorID,
			Name:      name,
			Biography: biography,
		},
	}
	defer rows.Close()
	defer stmt.Close()
	return book, nil
}

//GetAllBooks returns all Books Data
func GetAllBooks() ([]*schema.Book, error) {
	var books []*schema.Book
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Books.Price,Books.IsbnNo,Authors.ID,Authors.Name,Authors.Biography from Books inner join Authors where Books.AuthorID = Authors.ID;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookID, title, isbn_no, authorID, name, biography string
		var price int
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}

		book := &schema.Book{
			ID:     bookID,
			Title:  title,
			Price:  price,
			IsbnNo: isbn_no,
			Authors: &schema.Author{
				ID:        authorID,
				Name:      name,
				Biography: biography,
			},
		}
		books = append(books, book)
	}

	return books, nil
}

//GetAllBooks returns all Books Data
func GetAllBooksByAuthorName(name string) ([]*schema.Book, error) {
	var books []*schema.Book
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Books.Price,Books.IsbnNo,Authors.ID,Authors.Name,Authors.Biography from Books inner join Authors where Authors.Name = ? and Books.AuthorID = Authors.ID;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookID, title, isbn_no, authorID, name, biography string
		var price int
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}

		book := &schema.Book{
			ID:     bookID,
			Title:  title,
			Price:  price,
			IsbnNo: isbn_no,
			Authors: &schema.Author{
				ID:        authorID,
				Name:      name,
				Biography: biography,
			},
		}
		books = append(books, book)
	}

	return books, nil
}
