package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is a struct for user domain
type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

func (e *User) SetID() *User {
	e.ID = primitive.NewObjectID().Hex()
	return e
}
