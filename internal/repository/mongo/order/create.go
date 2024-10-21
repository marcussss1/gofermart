package order

import (
	"context"

	"gophermartLoyalty/internal/models"
)

func (r *Repository) Create(order *models.Order) error {
	_, err := r.collection.InsertOne(context.Background(), order)
	return err
}
