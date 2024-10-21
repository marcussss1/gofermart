package balance

import (
	"context"
	"github.com/google/uuid"
	"gofermart/internal/models"
	"gofermart/internal/repository"
	"gofermart/internal/utils"
	"time"
)

func (s *Service) Withdraw(ctx context.Context, userID string, orderNumber string, amount float64) error {
	if !utils.ValidateLuhn(orderNumber) {
		return repository.ErrInvalidOrderNumber
	}

	balance, err := s.repository.GetBalance(ctx, userID)
	if err != nil {
		return err
	}

	if balance.Current < amount {
		return repository.ErrInsufficientFunds
	}

	withdrawal := &models.Withdrawal{
		ID:          uuid.New().String(),
		UserID:      userID,
		Order:       orderNumber,
		Sum:         amount,
		ProcessedAt: time.Now(),
	}

	return s.repository.Withdraw(ctx, withdrawal)
}
