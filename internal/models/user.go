package models

type User struct {
	ID           string `bson:"_id,omitempty"`
	Login        string `bson:"login"`
	PasswordHash string `bson:"password_hash"`
}
