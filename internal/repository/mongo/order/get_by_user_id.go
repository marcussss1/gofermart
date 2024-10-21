package order

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gophermartLoyalty/internal/models"
)

func (r *Repository) GetByUserID(userID int) ([]models.Order, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID}, options.Find().SetSort(bson.D{{Key: "uploaded_at", Value: -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []models.Order
	if err = cursor.All(context.Background(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
