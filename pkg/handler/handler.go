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
		api.GET("/show-all-documents", h.showAllDocuments)
		api.GET("/show-my-documents", h.showMyDocuments)
		api.POST("/add-document", h.addDocument)

		document := api.Group("document/:id/")
		{
			document.GET("/users-need-to-sign", h.userNeedToSign)
			document.GET("/status", h.documentStatus)
		}
	}
	return router
}
