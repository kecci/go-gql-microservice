package health

type Service struct {
	healthRepo healthRepo
}

func New(repoHealth healthRepo) *Service {
	return &Service{
		healthRepo: repoHealth,
	}
}
