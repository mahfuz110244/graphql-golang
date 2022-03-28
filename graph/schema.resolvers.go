package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"fmt"
	"graphql-golang/graph/generated"
	"graphql-golang/graph/model"
	"math/big"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	fmt.Println(randomInt)
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", randomInt),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	fmt.Println(randomInt)
	book := &model.Book{
		ID:     fmt.Sprintf("B%d", randomInt),
		Title:  input.Title,
		Price:  input.Price,
		IsbnNo: input.IsbnNo,
		Authors: &model.Author{
			ID:        input.AuthorID,
			Name:      "author " + input.AuthorID,
			Biography: "biography " + input.AuthorID,
		},
	}
	r.books = append(r.books, book)
	return book, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.books, nil
}

func (r *queryResolver) Authors(ctx context.Context, name string) ([]*model.Book, error) {
	var books []*model.Book
	for _, book := range r.books {
		if book.Authors.Name == name {
			books = append(books, book)
		}
	}
	return books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
