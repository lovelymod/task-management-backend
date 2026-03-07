package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID             bson.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName      string        `json:"firstName" bson:"first_name"`
	LastName       string        `json:"lastName" bson:"last_name"`
	DisplayName    string        `json:"displayName" bson:"display_name"`
	Email          string        `json:"email" bson:"email"`
	HashedPassword string        `json:"-" bson:"hashed_password"`
	Phone          string        `json:"phone" bson:"phone"`
	Avatar         string        `json:"avatar" bson:"avatar"`
	CreatedAt      time.Time     `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updatedAt" bson:"updated_at"`
}
