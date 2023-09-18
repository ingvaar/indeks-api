package internal

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo is the struct containing everything related to mongodb.
type Mongo struct {
	client *mongo.Client
}

// NewMongo creates a new Mongo instance.
func NewMongo(ctx context.Context, config Config) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(config.Timeout))
	defer cancel()

	clientOpts := options.Client().ApplyURI(config.MongoURI)
	if err := clientOpts.Validate(); err != nil {
		return nil, err
	}

	log.Info().Msg("Connecting to Mongo")
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Info().Msg("Connected to Mongo")
	return &Mongo{client}, nil
}

// Disconnect the database.
func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
