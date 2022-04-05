package repository

import (
	"fmt"
	"log"

	db "graphql-golang/internal/pkg/db/mysql"
	"graphql-golang/model"
)

//CreateAuthor create's author
func CreateAuthor(author model.Author) (int64, error) {

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
func CreateBook(book model.Book) (int64, error) {
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
func GetBooksByID(id *string) (*model.Book, error) {
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Books.Price,Books.IsbnNo,Authors.ID,Authors.Name,Authors.Biography from Books inner join Authors where Books.AuthorID = Authors.ID and Books.ID = ? ;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	var bookID, title, isbn_no, authorID, name, biography string
	var price float64
	if rows.Next() {
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}
	}
	if bookID == "" {
		return nil, fmt.Errorf("Book not found")
	}
	book := &model.Book{
		ID:     bookID,
		Title:  title,
		Price:  price,
		IsbnNo: isbn_no,
		Authors: &model.Author{
			ID:        authorID,
			Name:      name,
			Biography: biography,
		},
	}
	defer rows.Close()
	defer stmt.Close()
	fmt.Println(book.Authors.ID)
	fmt.Println(book.Authors.Name)
	fmt.Println(book.Authors.Biography)
	return book, nil
}

//GetAllBooks returns all Books Data
func GetAllBooks() ([]*model.Book, error) {
	var books []*model.Book
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
		var price float64
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}

		book := &model.Book{
			ID:     bookID,
			Title:  title,
			Price:  price,
			IsbnNo: isbn_no,
			Authors: &model.Author{
				ID:        authorID,
				Name:      name,
				Biography: biography,
			},
		}
		books = append(books, book)
		fmt.Println(book.Authors.ID)
		fmt.Println(book.Authors.Name)
		fmt.Println(book.Authors.Biography)
	}
	fmt.Println(books)
	return books, nil
}

//GetAllBooks returns all Books Data
func GetAllBooksByAuthorName(name string) ([]*model.Book, error) {
	var books []*model.Book
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
		var price float64
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}

		book := &model.Book{
			ID:     bookID,
			Title:  title,
			Price:  price,
			IsbnNo: isbn_no,
			Authors: &model.Author{
				ID:        authorID,
				Name:      name,
				Biography: biography,
			},
		}
		books = append(books, book)
	}

	return books, nil
}

//GetAuthorByID return author with respective id
func GetAuthorByID(id *string) (*model.Author, error) {
	stmt, err := db.Db.Prepare("select * from Authors where id=?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	var author model.Author
	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Name, &author.Biography)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer rows.Close()

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	if author.ID == "" {
		return nil, fmt.Errorf("Author not found")
	}
	return &author, nil

}

//GetAllAuthors returns all authors
func GetAllAuthors() ([]*model.Author, error) {
	stmt, err := db.Db.Prepare("select * from Authors")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var authors []*model.Author
	for rows.Next() {
		var author model.Author
		rows.Scan(&author.ID, &author.Name, &author.Biography)
		authors = append(authors, &author)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	defer rows.Close()

	return authors, err
}
