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

var queryTypeAuthor = graphql.NewObject(
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
			   http://localhost:8080/author?query={list{id,name,biography}}
			*/
			"list": &graphql.Field{
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
		},
	})

var mutationTypeAuthor = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new author item
		http://localhost:8080/author?query=mutation+_{create(name:"John",biography:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)"){id,name,biography}}
		*/
		"create": &graphql.Field{
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
	},
})

var AuthorSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryTypeAuthor,
		Mutation: mutationTypeAuthor,
	},
)
