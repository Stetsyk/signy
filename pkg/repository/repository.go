package repository

import (
	"github.com/Stetsyk/signy"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user signy.User) (int, error)
	GetUser(username, password string) (signy.User, error)
}

type Document interface {
	GetOwn(userId int) ([]signy.Document, error)
	GetNeedToSign(userId int) ([]signy.Signature, error)
	AddDocument(document signy.Document, userNeedToSign []int) error
	GetStatus(userId int, documentId int) (string, error)
}

type Repository struct {
	Authorization
	Document
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Document:      NewDocumentMysql(db),
	}
}
