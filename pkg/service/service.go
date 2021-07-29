package service

import "github.com/Stetsyk/signy/pkg/repository"

type Authorization interface {
}

type Document interface {
}

type Service struct {
	Authorization
	Document
}

func NewService(repost *repository.Repository) *Service {
	return &Service{}
}
