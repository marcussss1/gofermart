package user

import (
	"context"
	"gofermart/internal/models"
)

type service interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, login, password string) (*models.User, error)
}
