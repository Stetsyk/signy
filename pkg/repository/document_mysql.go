package repository

import (
	"github.com/Stetsyk/signy"
	"gorm.io/gorm"
)

type DocumentMysql struct {
	db *gorm.DB
}

func NewDocumentMysql(db *gorm.DB) *DocumentMysql {
	return &DocumentMysql{db: db}
}

func (r *DocumentMysql) GetAll(userId int) ([]signy.Document, error) {
	var documents []signy.Document
	result := r.db.Where("owner_id = ?", userId).Find(&documents)
	if result.Error != nil {
		return documents, result.Error
	}
	return documents, nil
}
