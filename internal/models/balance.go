package models

import "time"

type Balance struct {
	UserID    string  `bson:"user_id"`
	Current   float64 `bson:"current"`
	Withdrawn float64 `bson:"withdrawn"`
}

type Withdrawal struct {
	ID          string    `bson:"_id,omitempty"`
	UserID      string    `bson:"user_id"`
	Order       string    `bson:"order"`
	Sum         float64   `bson:"sum"`
	ProcessedAt time.Time `bson:"processed_at"`
}
