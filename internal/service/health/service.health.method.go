package health

import (
	"context"

	model "github.com/kecci/go-gql-microservice/internal/model/health"
)

func (s *Service) CheckHealth() (*model.Health, error) {
	return s.healthRepo.CheckHealth(context.Background())
}
