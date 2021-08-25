package service

import (
	"github.com/Stetsyk/signy"
	"github.com/Stetsyk/signy/pkg/repository"
)

type DocumentService struct {
	repo repository.Document
}

func NewDocumentService(repo repository.Document) *DocumentService {
	return &DocumentService{repo: repo}
}
func (d *DocumentService) GetStatus(userId int, documentId int) (string, error) {
	status, err := d.repo.GetStatus(userId, documentId)
	if err != nil {
		return status, err
	}
	return status, nil
}

func (d *DocumentService) GetOwn(userId int) ([]signy.Document, error) {
	documents, err := d.repo.GetOwn(userId)
	if err != nil {
		return documents, err
	}
	return documents, nil
}

func (d *DocumentService) GetNeedToSign(userId int) ([]signy.Signature, error) {
	signatures, err := d.repo.GetNeedToSign(userId)
	if err != nil {
		return signatures, err
	}
	return signatures, nil
}

func (d *DocumentService) AddDocument(document signy.Document, userNeedToSign []int) error {
	return d.repo.AddDocument(document, userNeedToSign)
}
