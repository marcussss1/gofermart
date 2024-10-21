package order

import (
	"context"
	"gofermart/internal/models"
)

type service interface {
	UploadOrder(ctx context.Context, userID string, orderNumber string) error
	GetUserOrders(ctx context.Context, userID string) ([]models.Order, error)
}
