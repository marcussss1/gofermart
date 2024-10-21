package order

import (
	"context"
	"gofermart/internal/models"
)

type repository interface {
	UploadOrder(ctx context.Context, userID string, orderNumber string) error
	GetUserOrders(ctx context.Context, userID string) ([]models.Order, error)
	Create(ctx context.Context, order *models.Order) error
}
