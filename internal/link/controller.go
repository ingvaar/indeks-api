package link

import "github.com/gin-gonic/gin"

// Controller is the link controller interface.
type Controller interface {
	GetLinks(c *gin.Context)
	GetLinkByID(c *gin.Context)
	CreateLink(c *gin.Context)
	UpdateLinkByID(c *gin.Context)
	ReplaceLinkByID(c *gin.Context)
	DeleteLinkByID(c *gin.Context)
}
