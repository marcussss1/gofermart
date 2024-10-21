package user

import (
	"context"
	"gofermart/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, login, password string) (*models.User, error) {
	user, err := s.repository.GetByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
