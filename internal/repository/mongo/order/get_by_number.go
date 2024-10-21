package order

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gophermartLoyalty/internal/models"
	"gophermartLoyalty/internal/repository"
)

func (r *Repository) GetByNumber(number string) (*models.Order, error) {
	var order models.Order
	err := r.collection.FindOne(context.Background(), bson.M{"number": number}).Decode(&order)
	if err == mongo.ErrNoDocuments {
		return nil, repository.ErrOrderNotFound
	}
	return &order, err
}
