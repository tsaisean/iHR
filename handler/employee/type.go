package employee

import (
	"github.com/redis/go-redis/v9"
	repo "iHR/db/repositories"
)

type EmployeeHandler struct {
	repo  repo.EmployeeRepository
	cache *redis.Client
}

func NewEmployeeHandler(r repo.EmployeeRepository, redis *redis.Client) *EmployeeHandler {
	return &EmployeeHandler{
		repo:  r,
		cache: redis,
	}
}

type Cacheable interface {
	GetCacheKey() string
}
