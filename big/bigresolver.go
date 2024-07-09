package big

import (
	"context"

	"github.com/akafazov/gqlgen-client/graph"
	"github.com/akafazov/gqlgen-client/graph/model"
	libgraph "github.com/akafazov/gqlgen/graph"
	libmodel "github.com/akafazov/gqlgen/graph/model"
)

type BigResolver struct {
}

type BigMutationResolver interface {
	CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	CreateUser(ctx context.Context, input libmodel.NewUser) (*libmodel.User, error)
	CreateMeetup(ctx context.Context, input libmodel.NewMeetup) (*libmodel.Meetup, error)
	Mutation() graph.MutationResolver
}
type BigQueryResolver interface {
	Todos(ctx context.Context) ([]*model.Todo, error)
	Users(ctx context.Context) ([]*libmodel.User, error)
	Meetups(ctx context.Context) ([]*libmodel.Meetup, error)
	Query() graph.QueryResolver
}

type mutationResolver struct{ *BigResolver }
type queryResolver struct{ *BigResolver }

func (r *BigResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic("not implemented")
}

func (r *BigResolver) CreateMeetup(ctx context.Context, input libmodel.NewMeetup) (*libmodel.Meetup, error) {
	panic("not implemented")
}

func (r *BigResolver) CreateUser(ctx context.Context, input libmodel.NewUser) (*libmodel.User, error) {
	panic("not implemented")
}

func (r *BigResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := []*model.Todo{
		{ID: "1", Text: "todo1", Done: false},
	}
	return todos, nil
}

func (r *BigResolver) Users(ctx context.Context) ([]*libmodel.User, error) {
	panic("not implemented")
}

func (r *BigResolver) Meetups(ctx context.Context) ([]*libmodel.Meetup, error) {
	println("--- Meetups ---")
	return libgraph.QueryResolver.Meetups(r, ctx)
}

func (r *BigResolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

func (r *BigResolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}
