package user

import (
	"context"
	"gofermart/internal/models"
)

type repository interface {
	Register(ctx context.Context, user *models.User) error
	GetByLogin(ctx context.Context, login, password string) (*models.User, error)
}
