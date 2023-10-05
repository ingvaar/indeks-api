package internal

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

// Server holds everything related to the API server.
type Server struct {
	router   *gin.Engine
	database Database
	srv      *http.Server
}

// NewServer creates a new server.
func NewServer(router *gin.Engine, database Database, config Config) Server {
	return Server{
		router:   router,
		database: database,
		srv: &http.Server{
			Addr:              ":" + strconv.Itoa(int(config.Port)),
			Handler:           router,
			ReadHeaderTimeout: time.Second * time.Duration(config.Timeout),
		},
	}
}

// Start the server.
func (s Server) Start() error {
	eg := new(errgroup.Group)

	eg.Go(func() error {
		log.Info().Str("port", s.srv.Addr).Msgf("🚀 Server listening on %s", s.srv.Addr)
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// Stop the server.
func (s Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info().Msg("❌ Shutting down the server gracefully...")

	// Stopping the http server
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	// Disconnect from the database
	if err := s.database.Disconnect(ctx); err != nil {
		return err
	}

	log.Info().Msg("👍 Server stopped")
	return nil
}
