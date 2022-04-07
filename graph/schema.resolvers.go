package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"gopato.io/db"
	"gopato.io/graph/generated"
	"gopato.io/graph/model"
	"gopato.io/models/public"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTenant(ctx context.Context, input *model.NewTenant) (*model.Tenant, error) {
	fmt.Println(ctx)
	db := db.Db
	tenant := public.Tenant{Name: input.Name, Domain: input.Domain}
	fmt.Println(tenant)
	result := db.Create(&tenant)
	if result.Error != nil {
		fmt.Println(result.Error)
		return &model.Tenant{}, result.Error
	}
	return &model.Tenant{ID: strconv.Itoa(tenant.ID), Domain: tenant.Domain, Name: tenant.Name}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
