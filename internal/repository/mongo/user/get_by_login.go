package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gophermartLoyalty/internal/models"
	"gophermartLoyalty/internal/repository"
)

func (r *Repository) GetByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"login": login}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
