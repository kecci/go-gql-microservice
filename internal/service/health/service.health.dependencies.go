package health

//go:generate mockgen -source=service.health.dependencies.go -package=health -destination=service.health.dependencies_mock_test.go

import (
	"context"

	model "github.com/kecci/go-gql-microservice/internal/model/health"
)

type healthRepo interface {
	CheckHealth(ctx context.Context) (*model.Health, error)
}
