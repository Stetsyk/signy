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

func (d *DocumentService) GetOwn(userId int) ([]signy.Document, error) {
	documents, err := d.repo.GetOwn(userId)
	if err != nil {
		return documents, err
	}
	return documents, nil
}
