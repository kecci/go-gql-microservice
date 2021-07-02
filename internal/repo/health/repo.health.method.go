package health

import (
	"context"

	model "github.com/kecci/go-gql-microservice/internal/model/health"
)

// CheckDB health database
func (r *Repo) CheckHealth(ctx context.Context) (*model.Health, error) {
	return &model.Health{
		Message: "SERVED",
	}, nil
}
