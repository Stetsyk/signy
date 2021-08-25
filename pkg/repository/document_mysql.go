package repository

import (
	"errors"

	"github.com/Stetsyk/signy"
	"gorm.io/gorm"
)

type DocumentMysql struct {
	db *gorm.DB
}

func NewDocumentMysql(db *gorm.DB) *DocumentMysql {
	return &DocumentMysql{db: db}
}

func (r *DocumentMysql) GetStatus(userId int, documentId int) (string, error) {
	var status string
	var document signy.Document
	result := r.db.Where("id = ?", documentId).First(&document)
	if result.Error != nil {
		return status, result.Error
	}
	if result.RowsAffected == 0 {
		return status, errors.New("There's no such document")
	}

	var signatures []signy.Signature
	result = r.db.Where("document_id = ?", documentId).Find(&signatures)
	if result.RowsAffected == 0 {
		status = "There is no signatures"
	} else {
		cntRejected := 0
		cntApproved := 0
		cntUnsigned := 0
		for _, signature := range signatures {
			if signature.Status == "rejected" {
				cntRejected++
			} else if signature.Status == "approved" {
				cntApproved++
			} else if signature.Status == "Unsigned" {
				cntUnsigned++
			}
			if cntRejected > 0 {
				status = "rejected"
			} else if cntApproved == int(result.RowsAffected) {
				status = "approved"
			} else if cntUnsigned == int(result.RowsAffected) {
				status = "open"
			} else {
				status = "in_review"
			}
		}

	}
	return status, nil
}

func (r *DocumentMysql) GetOwn(userId int) ([]signy.Document, error) {
	var documents []signy.Document
	result := r.db.Where("owner_id = ?", userId).Find(&documents)
	if result.Error != nil {
		return documents, result.Error
	}
	return documents, nil
}

func (r *DocumentMysql) GetNeedToSign(userId int) ([]signy.Signature, error) {
	var signatures []signy.Signature
	result := r.db.Where("user_id = ? AND status = ?", userId, "Unsigned").Find(&signatures)
	if result.Error != nil {
		return signatures, result.Error
	}
	return signatures, nil
}

func (r *DocumentMysql) AddDocument(document signy.Document, userNeedToSign []int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&document).Error; err != nil {
			return err
		}
		for _, userId := range userNeedToSign {
			var signature signy.Signature
			signature.DocumentId = int(document.Id)
			signature.UserId = userId
			signature.Status = "Unsigned"
			if err := tx.Create(&signature).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
