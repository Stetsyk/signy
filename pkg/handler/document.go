package handler

import (
	"net/http"

	"github.com/Stetsyk/signy"
	"github.com/gin-gonic/gin"
)

func (h *Handler) showAllUsers(c *gin.Context) {

}

type getOwnDocumentsResponse struct {
	Documents []signy.Document `json:"documents"`
}

func (h *Handler) showOwnedDocuments(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	documents, err := h.services.Document.GetOwn(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOwnDocumentsResponse{
		Documents: documents,
	})
}

func (h *Handler) showNeedToSignDocuments(c *gin.Context) {

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
