package balance

import (
	"context"
	"gofermart/internal/models"
)

type repository interface {
	GetBalance(ctx context.Context, userID string) (*models.Balance, error)
	Withdraw(ctx context.Context, userID string, orderNumber string, amount float64) error
	GetWithdrawals(ctx context.Context, userID string) ([]models.Withdrawal, error)
}
