package balance

import (
	"context"
	"gofermart/internal/models"
)

func (s *Service) GetWithdrawals(ctx context.Context, userID string) ([]models.Withdrawal, error) {
	return s.repository.GetWithdrawals(ctx, userID)
}
