package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConvertHandlerToGinHandler is used to convert a http middleware to gin middleware.
func ConvertHandlerToGinHandler(mid func(http.Handler) http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		mid(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	}
}
