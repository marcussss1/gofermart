package user

import (
	"context"
	"gophermartLoyalty/internal/models"
	"gophermartLoyalty/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

var (
	ErrUserAlreadyExists   = errors.New("пользователь уже существует")
	ErrInvalidCredentials  = errors.New("неверные учетные данные")
)

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{
		repository: repository,
	}
}
