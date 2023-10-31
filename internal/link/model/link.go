package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Link represents a link in the database.
type Link struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Path      string             `bson:"path"`
	Active    bool               `bson:"active"`
	CreatedAt time.Time          `bson:"created_at"`
}
