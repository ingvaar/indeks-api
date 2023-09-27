package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ingvaar/indeks-api/internal/helper"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	s.router.Use(accessLoggerMiddleware())
}

func (s Server) initRoutes() {
	s.router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	s.router.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})
}
