package order

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) UpdateStatus(orderID int, status string) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": orderID},
		bson.M{"$set": bson.M{"status": status}},
	)
	return err
}
