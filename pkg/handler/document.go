package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) showAllUsers(c *gin.Context) {

}

func (h *Handler) showAllDocuments(c *gin.Context) {

}

func (h *Handler) showMyDocuments(c *gin.Context) {

}

func (h *Handler) addDocument(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) userNeedToSign(c *gin.Context) {

}

func (h *Handler) documentStatus(c *gin.Context) {

}
