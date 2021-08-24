package service

import (
	"github.com/Stetsyk/signy"
	"github.com/Stetsyk/signy/pkg/repository"
)

type Authorization interface {
	CreateUser(user signy.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Document interface {
	GetAll(userId int) ([]signy.Document, error)
}

type Service struct {
	Authorization
	Document
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Document:      NewDocumentService(repos.Document),
	}
}
