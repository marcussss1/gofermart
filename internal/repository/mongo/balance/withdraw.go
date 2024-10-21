package balance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gophermartLoyalty/internal/models"
	"gophermartLoyalty/internal/repository"
)

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
