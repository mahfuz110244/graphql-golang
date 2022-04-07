package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "graphql-golang/internal/pkg/db/mysql"
	"graphql-golang/schema"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
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

const defaultPort = "8080"

func main() {
	port := defaultPort
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
		result := executeQuery(r.URL.Query().Get("query"), schema.Schema)
		json.NewEncoder(w).Encode(result)
	})

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
