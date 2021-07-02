package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kecci/go-gql-microservice/graph/generated"
	"github.com/kecci/go-gql-microservice/internal/model/health"
)

func (r *queryResolver) CheckHealth(ctx context.Context) (*health.Health, error) {
	return &health.Health{Message: "SUCCESS"}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
