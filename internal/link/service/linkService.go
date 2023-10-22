package service

import "github.com/ingvaar/indeks-api/internal/link"

type linkService struct {
	repository link.Repository
}

// NewService creates a new service.
func NewService(repository *link.Repository) link.Service {
	return linkService{
		repository,
	}
}