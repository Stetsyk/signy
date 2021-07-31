package repository

import "gorm.io/gorm"

type Authorization interface {
}

type Document interface {
}

type Repository struct {
	Authorization
	Document
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}
