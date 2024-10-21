package orderv2

import (
	"context"
	"github.com/google/uuid"
	"gofermart/internal/models"
	"time"
)

func (s *Service) UploadOrder(ctx context.Context) error {
	return s.repository.UploadOrder(ctx, &models.OrderV2{
		ID:         uuid.NewString(),
		Status:     models.StatusProcessing,
		UploadedAt: time.Now(),
	})
}
