package models

import "time"

type Order struct {
	ID         string    `bson:"_id,omitempty"`
	Number     string    `bson:"number"`
	UserID     string    `bson:"user_id"`
	Status     string    `bson:"status"`
	Accrual    float64   `bson:"accrual"`
	UploadedAt time.Time `bson:"uploaded_at"`
}

//const (
//	StatusNew        = "NEW"
//	StatusProcessing = "PROCESSING"
//	StatusInvalid    = "INVALID"
//	StatusProcessed  = "PROCESSED"
//)
