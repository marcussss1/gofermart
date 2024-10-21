package balance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gophermartLoyalty/internal/models"
)

func (r *Repository) GetWithdrawals(userID int) ([]models.Withdrawal, error) {
	cursor, err := r.withdrawalCollection.Find(
		context.Background(),
		bson.M{"user_id": userID},
		options.Find().SetSort(bson.D{{Key: "processed_at", Value: -1}}),
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var withdrawals []models.Withdrawal
	if err = cursor.All(context.Background(), &withdrawals); err != nil {
		return nil, err
	}
	return withdrawals, nil
}
