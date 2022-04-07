package schema

import (
	"graphql-golang/repository"
	"strconv"

	"graphql-golang/model"

	"github.com/graphql-go/graphql"
)

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"biography": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"isbn_no": &graphql.Field{
				Type: graphql.String,
			},
			"authors": &graphql.Field{
				Type: authorType,
			},
		},
	},
)

var booksType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Books",
		Fields: graphql.Fields{
			"books": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get book list",
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single author by id
			   http://localhost:8080/author?query={author(id:1){name,biography}}
			*/
			"author": &graphql.Field{
				Type:        authorType,
				Description: "Get author by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					idStr := strconv.Itoa(id)
					author, err := repository.GetAuthorByID(&idStr)
					if err != nil {
						return nil, err
					} else {
						return author, nil
					}
				},
			},
			/* Get (read) author list
			   http://localhost:8080/author_list?query={list{id,name,biography}}
			*/
			"author_list": &graphql.Field{
				Type:        graphql.NewList(authorType),
				Description: "Get author list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					authors, err := repository.GetAllAuthors()
					if err != nil {
						return nil, err
					} else {
						return authors, err
					}
				},
			},

			/* Get (read) single book by id
			   http://localhost:8080/book?query={book(id:1){id,title,price,isbn_no,author{name,biography}}}
			*/
			"book": &graphql.Field{
				Type:        bookType,
				Description: "Get book by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					idStr := strconv.Itoa(id)
					book, err := repository.GetBooksByID(&idStr)
					if err != nil {
						return nil, err
					} else {
						return book, nil
					}
				},
			},
			/* Get (read) book list
			   http://localhost:8080/book?query={list{id,title,price,isbn_no,author{id,name,biography}}
			*/
			"books": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get book list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					books, err := repository.GetAllBooks()
					if err != nil {
						return nil, err
					} else {
						return books, err
					}
				},
			},

			/* Get (read) book list for authors
			   http://localhost:8080/authors?query={authors(name:"John"){books{id,title,price,isbn_no,author{id,name,biography}}
			*/
			"authors": &graphql.Field{
				Type:        booksType,
				Description: "Get book list for authors",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name := params.Args["name"].(string)
					books, err := repository.GetAllBooksByAuthorName(name)
					if err != nil {
						return nil, err
					} else {
						booksData := &model.Books{
							Books: books,
						}
						return booksData, err
						// return books, nil
					}
				},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new author item
		http://localhost:8080/author?query=mutation+_{createAuthor(name:"John",biography:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)"){id,name,biography}}
		*/
		"createAuthor": &graphql.Field{
			Type:        authorType,
			Description: "Create new author",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"biography": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var author model.Author
				author.Name = params.Args["name"].(string)
				author.Biography = params.Args["biography"].(string)
				id, err := repository.CreateAuthor(author)
				if err != nil {
					return nil, err
				} else {
					return &model.Author{ID: strconv.FormatInt(id, 10), Name: author.Name, Biography: author.Biography}, nil
				}
			},
		},

		/* Create new book item
		http://localhost:8080/book?query=mutation+_{createBook(title:"Book 1",price:1000,isbn_no:"6678557878798",author_id:"1"){id,title,price,isbn_no,author{id,name,biography}}}
		*/
		"createBook": &graphql.Field{
			Type:        bookType,
			Description: "Create new book",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"isbn_no": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"author_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var book model.Book
				book.Title = params.Args["title"].(string)
				book.IsbnNo = params.Args["isbn_no"].(string)
				book.Price = params.Args["price"].(float64)
				authorID := params.Args["author_id"].(string)
				book.Authors = &model.Author{
					ID: authorID,
				}
				id, err := repository.CreateBook(book)
				if err != nil {
					return nil, err
				}
				idStr := strconv.Itoa(int(id))
				createdBook, _ := repository.GetBooksByID(&idStr)
				// return &model.Book{ID: createdBook.ID, Title: createdBook.Title, IsbnNo: createdBook.IsbnNo, Price: createdBook.Price, Authors: &model.Author{
				// 	ID:        createdBook.Authors.ID,
				// 	Name:      createdBook.Authors.Name,
				// 	Biography: createdBook.Authors.Biography,
				// }}, nil
				return createdBook, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
