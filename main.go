package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "graphql-golang/internal/pkg/db/mysql"
	"graphql-golang/schema"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	db.InitDB()
	db.Migrate()

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema.ProductSchema)
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/author", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema.AuthorSchema)
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema.BookSchema)
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
	// h1 := handler.New(&handler.Config{
	// 	Schema:   &schema.BookSchema,
	// 	Pretty:   true,
	// 	GraphiQL: true,
	// })

	// h := handler.New(&handler.Config{
	// 	Schema:   &schema.BookSchema,
	// 	Pretty:   true,
	// 	GraphiQL: true,
	// })

	// http.Handle("/graphql", h)
	// http.Handle("/graphiql", h1)
	// http.ListenAndServe(":8080", nil)
}
