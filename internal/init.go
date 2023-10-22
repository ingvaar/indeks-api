package internal

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Config holds the app configuration.
type Config struct {
	Port     uint
	Timeout  uint
	MongoURI string
	DevMode  bool
}

// Database defines the Database struct.
type Database interface {
	Disconnect(context.Context) error
	GetClient() any
}

// New creates a new server with provided config.
func New(config Config) (Server, error) {
	ctx := context.Background()
	gin.SetMode(gin.ReleaseMode)
	if config.DevMode {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.New()
	database, err := NewMongo(ctx, config)
	if err != nil {
		return Server{}, err
	}

	srv := NewServer(router, database, config)
	srv.initMiddleware()
	srv.initRoutes()

	return srv, nil
}
