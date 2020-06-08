package resolvers

import (
	"context"
	"fmt"
	"github.com/MisakiFx/Golang/gqlgen/graph/types"
	"strconv"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*types.Todo, error) {
	fmt.Println("Todos()")
	return []*types.Todo{
		&types.Todo{
			ID:     strconv.FormatInt(1, 10),
			Text:   "todo text",
			UserID: "9527",
			Done:   false,
		},
	}, nil
	//panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*types.Todo, error) {
	fmt.Println("Todo()")
	return &types.Todo{
		ID:     id,
		Text:   "todo text",
		UserID: "9527",
		Done:   false,
	}, nil
	//panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Todo2(ctx context.Context, id string) (*types.Todo, error) {
	fmt.Println("Todo()")
	return &types.Todo{
		ID:     id,
		Text:   "todo text",
		UserID: "9527",
		Done:   false,
	}, nil
	//panic(fmt.Errorf("not implemented"))
}
