package repository

import (
	"github.com/ingvaar/indeks-api/internal/link"
	"go.mongodb.org/mongo-driver/mongo"
)

type linkRepository struct {
	mongo *mongo.Client
}

// NewRepository creates a new repository.
func NewRepository(mongo *mongo.Client) link.Repository {
	return linkRepository{
		mongo,
	}
}