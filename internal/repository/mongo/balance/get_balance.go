package balance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gophermartLoyalty/internal/models"
	"gophermartLoyalty/internal/repository"
)

func (r *Repository) GetBalance(userID int) (*models.Balance, error) {
	var balance models.Balance
	err := r.balanceCollection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&balance)
	if err == mongo.ErrNoDocuments {
		return nil, repository.ErrBalanceNotFound
	}
	return &balance, err
}
