package handler

import (
	"github.com/Stetsyk/signy/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/show-all-users", h.showAllUsers)
		api.GET("/show-owned-documents", h.showOwnedDocuments)
		api.GET("/show-need-to-sign", h.showNeedToSignDocuments)
		api.POST("/add-document", h.addDocument)

		document := api.Group("document/:id/")
		{
			document.POST("/sign-document", h.signDocument)
			document.GET("/status", h.documentStatus)
		}
	}
	return router
}
