package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ingvaar/indeks-api/internal/link"
)

type linkController struct {
	service link.Service
}

// CreateLink implements link.Controller.
func (linkController) CreateLink(c *gin.Context) {
	panic("unimplemented")
}

// DeleteLinkByID implements link.Controller.
func (linkController) DeleteLinkByID(c *gin.Context) {
	panic("unimplemented")
}

// GetLinkByID implements link.Controller.
func (linkController) GetLinkByID(c *gin.Context) {
	panic("unimplemented")
}

// GetLinks implements link.Controller.
func (linkController) GetLinks(c *gin.Context) {
	panic("unimplemented")
}

// ReplaceLinkByID implements link.Controller.
func (linkController) ReplaceLinkByID(c *gin.Context) {
	panic("unimplemented")
}

// UpdateLinkByID implements link.Controller.
func (linkController) UpdateLinkByID(c *gin.Context) {
	panic("unimplemented")
}

// NewController creates a new controller.
func NewController(service *link.Service) link.Controller {
	return linkController{
		service,
	}
}
