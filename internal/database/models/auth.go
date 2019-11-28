package models

// Credentials is a user registration credentials
type Credentials struct {
	Login    string `bson:"_id"`
	PassHash string `bson:"password_hash"`
}
