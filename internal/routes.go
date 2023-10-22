package internal

import (
	"net/http"

	"github.com/ingvaar/indeks-api/internal/helper"
	linkController "github.com/ingvaar/indeks-api/internal/link/controller"
	linkRepository "github.com/ingvaar/indeks-api/internal/link/repository"
	linkService "github.com/ingvaar/indeks-api/internal/link/service"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
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
	// Link service
	linkRepository := linkRepository.NewRepository(s.database.GetClient().(*mongo.Client))
	linkService := linkService.NewService(&linkRepository)
	linkController := linkController.NewController(&linkService)

	s.router.GET("/link", linkController.GetLinks)
	s.router.GET("/link/:id", linkController.GetLinkByID)
	s.router.POST("/link", linkController.CreateLink)
	s.router.PATCH("/link/:id", linkController.UpdateLinkByID)
	s.router.PUT("/link/:id", linkController.ReplaceLinkByID)
	s.router.DELETE("/link/:id", linkController.DeleteLinkByID)
	// -- Link service

	s.router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	s.router.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})
}
