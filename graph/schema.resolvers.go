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

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	// fmt.Println(randomInt)
	book := &model.Book{
		ID:     fmt.Sprintf("B%d", randomInt),
		Title:  input.Title,
		Price:  input.Price,
		IsbnNo: input.IsbnNo,
		Authors: &model.Author{
			ID:        input.Author.ID,
			Name:      input.Author.Name,
			Biography: input.Author.Biography,
		},
	}
	r.books = append(r.books, book)
	return book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.books, nil
}

func (r *queryResolver) Authors(ctx context.Context, name string) (*model.Books, error) {
	var books []*model.Book
	for _, bk := range r.books {
		if bk.Authors.Name == name {
			books = append(books, bk)
		}
	}
	return &model.Books{Books: books}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
