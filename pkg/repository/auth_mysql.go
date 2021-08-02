package repository

import (
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
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}
func (r *AuthMysql) GetUser(username, password string) (signy.User, error) {
	var user signy.User
	result := r.db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return signy.User{}, result.Error
	}
	return user, nil
}
