package balance

import (
	"context"
	"gofermart/internal/models"
)

func (s *Service) GetBalance(ctx context.Context, userID string) (*models.Balance, error) {
	return s.repository.GetBalance(ctx, userID)
}
