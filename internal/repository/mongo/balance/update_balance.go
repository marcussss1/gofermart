package balance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) UpdateBalance(userID int, amount float64) error {
	opts := options.Update().SetUpsert(true)
	_, err := r.balanceCollection.UpdateOne(
		context.Background(),
		bson.M{"user_id": userID},
		bson.M{
			"$inc": bson.M{"current": amount},
			"$setOnInsert": bson.M{"withdrawn": 0},
		},
		opts,
	)
	return err
}
