package employee

import (
	repo "iHR/repositories"
)

type EmployeeHandler struct {
	repo repo.EmployeeRepository
}

func NewEmployeeHandler(r repo.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{
		repo: r,
	}
}

type Cacheable interface {
	GetCacheKey() string
}
