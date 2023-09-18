package internal

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ingvaar/indeks/internal/helper"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func (s Server) initMiddleware() {
	s.router.Use(gin.Recovery())

	s.router.Use(
		helper.ConvertHandlerToGinHandler(hlog.NewHandler(log.Logger)),
	)
	s.router.Use(
		helper.ConvertHandlerToGinHandler(
			hlog.RequestIDHandler("req-id", "Request-Id"),
		),
	)
	s.router.Use(
		helper.ConvertHandlerToGinHandler(
			hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
				hlog.FromRequest(r).Info().
					Str("method", r.Method).
					Stringer("url", r.URL).
					Int("status", status).
					Int("size", size).
					Dur("duration", duration).
					Msg("")
			}),
		),
	)
}

func (s Server) initRoutes() {
	s.router.POST("/link", func(c *gin.Context) {
		c.AbortWithStatus(501)
	})
	s.router.GET("/link", func(c *gin.Context) {
		c.AbortWithStatus(501)
	})
}
