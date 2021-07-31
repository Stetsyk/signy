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
}

type Repository struct {
	Authorization
	Document
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
	}
}
