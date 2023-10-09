package internal

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/hlog"
)

func accessLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeStart := time.Now().UTC()

		c.Next()

		hlog.FromRequest(c.Request).Info().
			Str("method", c.Request.Method).
			Stringer("url", c.Request.URL).
			Int("status", c.Writer.Status()).
			Int("size", c.Writer.Size()).
			Dur("duration", time.Now().UTC().Sub(timeStart)).
			Msg("")
	}
}
