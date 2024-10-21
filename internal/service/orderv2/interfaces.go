package orderv2

import (
	"context"
	"gofermart/internal/models"
)

type repository interface {
	UploadOrder(ctx context.Context, order *models.OrderV2) error
}
