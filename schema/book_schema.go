package schema

import (
	"fmt"
	"graphql-golang/repository"
	"strconv"

	"graphql-golang/model"

	"github.com/graphql-go/graphql"
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
			// "author": &graphql.Field{
			// 	Type: graphql.String,
			// },
			// "author": &graphql.Field{
			// 	Type: graphql.NewList(authorType),
			// },
			"author": &graphql.Field{
				Type: authorType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"biography": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
		},
	},
)

var queryTypeBook = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
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
					// "title": &graphql.ArgumentConfig{
					// 	Type: graphql.String,
					// },
					// "price": &graphql.ArgumentConfig{
					// 	Type: graphql.Float,
					// },
					// "isbn_no": &graphql.ArgumentConfig{
					// 	Type: graphql.String,
					// },
					// "author": &graphql.ArgumentConfig{
					// 	Type: graphql.NewList(authorType),
					// },
					// "author": &graphql.ArgumentConfig{
					// 	Type: authorType,
					// 	// Args: graphql.FieldConfigArgument{
					// 	// 	"id": &graphql.ArgumentConfig{
					// 	// 		Type: graphql.Int,
					// 	// 	},
					// 	// 	"name": &graphql.ArgumentConfig{
					// 	// 		Type: graphql.String,
					// 	// 	},
					// 	// 	"biography": &graphql.ArgumentConfig{
					// 	// 		Type: graphql.String,
					// 	// 	},
					// 	// },
					// },
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					idStr := strconv.Itoa(id)
					book, err := repository.GetBooksByID(&idStr)
					fmt.Println(book.Authors.ID)
					fmt.Println(book.Authors.Name)
					fmt.Println(book.Authors.Biography)
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
			"list": &graphql.Field{
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
		},
	})

var mutationTypeBook = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new book item
		http://localhost:8080/book?query=mutation+_{create(title:"Book 1",price:1000,isbn_no:"6678557878798",author:"1"){id,title,price,isbn_no,author{id,name,biography}}}
		*/
		"create": &graphql.Field{
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
				"author": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				// "author": &graphql.ArgumentConfig{
				// 	Type: graphql.String,
				// 	// Args: graphql.FieldConfigArgument{
				// 	// 	"id": &graphql.ArgumentConfig{
				// 	// 		Type: graphql.Int,
				// 	// 	},
				// 	// 	"name": &graphql.ArgumentConfig{
				// 	// 		Type: graphql.String,
				// 	// 	},
				// 	// 	"biography": &graphql.ArgumentConfig{
				// 	// 		Type: graphql.String,
				// 	// 	},
				// 	// },
				// },
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var book model.Book
				book.Title = params.Args["title"].(string)
				book.IsbnNo = params.Args["isbn_no"].(string)
				book.Price = params.Args["price"].(float64)
				authorID := params.Args["author"].(string)
				// authorIDStr := strconv.Itoa(authorID)
				book.Authors = &model.Author{
					ID: authorID,
				}
				id, err := repository.CreateBook(book)
				if err != nil {
					return nil, err
				}
				idStr := strconv.Itoa(int(id))
				createdBook, _ := repository.GetBooksByID(&idStr)
				return createdBook, nil
				// id, err := repository.CreateBook(book)
				// if err != nil {
				// 	return nil, err
				// } else {
				// 	return &model.Book{ID: strconv.FormatInt(id, 10), Title: book.Title, IsbnNo: book.IsbnNo, Price: book.Price, Authors: book.Authors}, nil
				// }
			},
		},
	},
})

var BookSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryTypeBook,
		Mutation: mutationTypeBook,
	},
)
