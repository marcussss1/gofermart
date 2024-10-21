package order

import (
	"context"
	"github.com/google/uuid"
	"gofermart/internal/models"
	"gofermart/internal/repository"
	"gofermart/internal/utils"
	"time"
)

func (s *Service) UploadOrder(ctx context.Context, userID string, orderNumber string) error {
	if !utils.ValidateLuhn(orderNumber) {
		return repository.ErrInvalidOrderNumber
	}

	existingOrder, err := s.repository.GetByNumber(orderNumber)
	if err != nil && err != repository.ErrOrderNotFound {
		return err
	}

	if existingOrder != nil {
		if existingOrder.UserID == userID {
			return repository.ErrOrderAlreadyUploaded
		}
		return repository.ErrOrderUploadedByAnotherUser
	}

	order := &models.Order{
		ID:         uuid.New().String(),
		Number:     orderNumber,
		UserID:     userID,
		Status:     models.StatusNew,
		UploadedAt: time.Now(),
	}

	return s.repository.Create(ctx, order)
}
