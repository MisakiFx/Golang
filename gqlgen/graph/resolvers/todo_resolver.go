package resolvers

import (
	"context"
	"fmt"
	"github.com/MisakiFx/Golang/gqlgen/graph/model"
	"github.com/MisakiFx/Golang/gqlgen/graph/types"
)

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *types.Todo) (*model.User, error) {
	fmt.Println("User()")
	return &model.User{
		ID:   obj.UserID,
		Name: "Misaki",
	}, nil
	//panic(fmt.Errorf("not implemented"))
}
