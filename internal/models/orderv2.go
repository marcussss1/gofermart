package models

import "time"

type OrderV2 struct {
	ID         string        `bson:"_id,omitempty"`
	Status     OrderV2Status `bson:"status"`
	UploadedAt time.Time     `bson:"uploaded_at"`
}

type OrderV2Status string

const (
	StatusNew        OrderV2Status = "NEW"
	StatusProcessing OrderV2Status = "PROCESSING"
	StatusInvalid    OrderV2Status = "INVALID"
	StatusProcessed  OrderV2Status = "PROCESSED"
)
