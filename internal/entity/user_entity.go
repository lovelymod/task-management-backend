package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID             bson.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName      string        `json:"firstName" bson:"firstName"`
	LastName       string        `json:"lastName" bson:"lastName"`
	DisplayName    string        `json:"displayName" bson:"displayName"`
	Email          string        `json:"email" bson:"email"`
	HashedPassword string        `json:"-" bson:"hashedPassword"`
	Phone          string        `json:"phone" bson:"phone"`
	Avatar         string        `json:"avatar" bson:"avatar"`
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt" bson:"updatedAt"`
}
