package resolver

//go:generate mockgen -source=resolver.go -package=resolver -destination=resolver.mock_test.go

import (
	model "github.com/kecci/go-gql-microservice/internal/model/health"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type healthServiceInterface interface {
	CheckHealth() (*model.Health, error)
}

// Resolver struct
type Resolver struct {
	healthService healthServiceInterface
}

// NewResolver constructor
func NewResolver(
	healthService healthServiceInterface,
) *Resolver {
	return &Resolver{
		healthService: healthService,
	}
}
