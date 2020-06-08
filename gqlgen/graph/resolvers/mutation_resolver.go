package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MisakiFx/Golang/gqlgen/graph/model"
	"github.com/MisakiFx/Golang/gqlgen/graph/types"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*types.Todo, error) {
	fmt.Println("CreateTodo()")
	return &types.Todo{
		ID:     strconv.FormatInt(10, 10),
		Text:   input.Text,
		UserID: input.UserID,
		Done:   false,
	}, nil
	//panic(fmt.Errorf("not implemented"))
}

type mutationResolver struct{ *Resolver }
