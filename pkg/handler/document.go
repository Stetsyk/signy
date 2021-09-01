package handler

import (
	"net/http"
	"strconv"

	"github.com/Stetsyk/signy"
	"github.com/gin-gonic/gin"
)

type getOwnedDocumentsResponse struct {
	Documents []signy.Document `json:"documents"`
}

type getNeedToSignResponse struct {
	Signatures []signy.Signature `json:"signatures"`
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

	c.JSON(http.StatusOK, map[string]interface{}{
		"Documents": documents,
	})
}

func (h *Handler) showNeedToSignDocuments(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	signatures, err := h.services.Document.GetNeedToSign(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Signatures": signatures,
	})
}

const (
	UsersNeedToSignCtx = "UsersNeedToSign"
)

type AddDocumentDTO struct {
	Document        signy.Document `json:"Document"`
	UsersNeedToSign []int          `json:"UsersNeedToSign"`
}

func (h *Handler) addDocument(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var addDocument AddDocumentDTO
	if err := c.BindJSON(&addDocument); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	document := addDocument.Document
	usersNeedToSign := addDocument.UsersNeedToSign

	err = h.services.Document.AddDocument(document, usersNeedToSign)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"users":    usersNeedToSign,
		"document": document,
	})
}

type SignDocumentDTO struct {
	Signature string `json:"Signature"`
}

func (h *Handler) signDocument(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	documentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var signDocumentDTO SignDocumentDTO
	if err := c.BindJSON(&signDocumentDTO); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	signature := signDocumentDTO.Signature
	status, err := h.services.Document.SignDocument(userId, documentId, signature)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"DocumentId": documentId,
		"Status":     status,
	})
}

func (h *Handler) documentStatus(c *gin.Context) {
	documentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	status, err := h.services.Document.GetStatus(userId, documentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": status,
	})
}
