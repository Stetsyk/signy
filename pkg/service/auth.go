package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Stetsyk/signy"
	"github.com/Stetsyk/signy/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "sdfasdnkjnvfadsffdsfdfggfvsfdsaf"
	tokenTTL   = 12 * time.Hour
	signingKey = "asdfadsfdasf"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user signy.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}