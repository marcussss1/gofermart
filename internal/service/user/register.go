package user

import (
	"gophermartLoyalty/internal/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func (s *Service) Register(user *models.User) error {
	// Проверка, существует ли пользователь с таким логином
	existingUser, err := s.repo.GetByLogin(user.Login)
	if err == nil && existingUser != nil {
		return ErrUserAlreadyExists
	}

	// Генерация UUID
	user.ID = uuid.New().String()

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	// Создание пользователя
	return s.repo.Create(user)
}
