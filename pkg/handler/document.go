package handler

import (
	"net/http"

	"github.com/Stetsyk/signy"
	"github.com/gin-gonic/gin"
)

func (h *Handler) showAllUsers(c *gin.Context) {

}

type getAllDocumentsResponse struct {
	Documents []signy.Document `json:"documents"`
}

func (h *Handler) showAllDocuments(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	documents, err := h.services.Document.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDocumentsResponse{
		Documents: documents,
	})
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
