package user

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gophermartLoyalty/internal/models"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("users"),
	}
}

func (r *Repository) Create(user *models.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}
