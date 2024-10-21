package balance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gophermartLoyalty/internal/models"
	"gophermartLoyalty/internal/repository"
)

type Repository struct {
	balanceCollection    *mongo.Collection
	withdrawalCollection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		balanceCollection:    db.Collection("balances"),
		withdrawalCollection: db.Collection("withdrawals"),
	}
}

func (r *Repository) GetBalance(userID int) (*models.Balance, error) {
	var balance models.Balance
	err := r.balanceCollection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&balance)
	if err == mongo.ErrNoDocuments {
		return nil, repository.ErrBalanceNotFound
	}
	return &balance, err
}

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

func (r *Repository) Withdraw(withdrawal *models.Withdrawal) error {
	session, err := r.balanceCollection.Database().Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	_, err = session.WithTransaction(context.Background(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Уменьшаем баланс
		result, err := r.balanceCollection.UpdateOne(
			sessCtx,
			bson.M{"user_id": withdrawal.UserID, "current": bson.M{"$gte": withdrawal.Sum}},
			bson.M{
				"$inc": bson.M{
					"current":   -withdrawal.Sum,
					"withdrawn": withdrawal.Sum,
				},
			},
		)
		if err != nil {
			return nil, err
		}
		if result.ModifiedCount == 0 {
			return nil, repository.ErrInsufficientFunds
		}

		// Добавляем запись о выводе средств
		_, err = r.withdrawalCollection.InsertOne(sessCtx, withdrawal)
		return nil, err
	})

	return err
}

func (r *Repository) GetWithdrawals(userID int) ([]models.Withdrawal, error) {
	cursor, err := r.withdrawalCollection.Find(
		context.Background(),
		bson.M{"user_id": userID},
		options.Find().SetSort(bson.D{{"processed_at", -1}}),
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
