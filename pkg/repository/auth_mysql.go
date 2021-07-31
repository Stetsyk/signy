package repository

import (
	"fmt"

	"github.com/Stetsyk/signy"
	"gorm.io/gorm"
)

const (
	userTable      = "users"
	documentTable  = "documents"
	signatureTable = "signatures"
)

type AuthMysql struct {
	db *gorm.DB
}

func NewAuthMysql(db *gorm.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user signy.User) (int, error) {
	if r.db == nil {
		fmt.Println("oleksii lol")
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}
