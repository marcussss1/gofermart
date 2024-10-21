package order

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gophermartLoyalty/internal/models"
	"gophermartLoyalty/internal/repository"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("orders"),
	}
}

func (r *Repository) Create(order *models.Order) error {
	_, err := r.collection.InsertOne(context.Background(), order)
	return err
}

func (r *Repository) GetByNumber(number string) (*models.Order, error) {
	var order models.Order
	err := r.collection.FindOne(context.Background(), bson.M{"number": number}).Decode(&order)
	if err == mongo.ErrNoDocuments {
		return nil, repository.ErrOrderNotFound
	}
	return &order, err
}

func (r *Repository) GetByUserID(userID int) ([]models.Order, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID}, options.Find().SetSort(bson.D{{"uploaded_at", -1}}))
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

func (r *Repository) UpdateStatus(orderID int, status string) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": orderID},
		bson.M{"$set": bson.M{"status": status}},
	)
	return err
}
