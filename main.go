package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "graphql-golang/internal/pkg/db/mysql"
	"graphql-golang/schema"

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
		result := executeQuery(r.URL.Query().Get("query"), schema.Productchema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
