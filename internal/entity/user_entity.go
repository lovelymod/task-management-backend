package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID  `json:"id" bson:"_id"`
	FirstName      string              `json:"firstName"`
	LastName       string              `json:"lastName"`
	DisplayName    string              `json:"displayName"`
	Email          string              `json:"email"`
	Username       string              `json:"username"`
	HashedPassword string              `json:"-" gorm:"not null"`
	Phone          string              `json:"phone"`
	Avatar         string              `json:"avatar"`
	CreatedAt      primitive.Timestamp `json:"createdAt"`
	UpdatedAt      primitive.Timestamp `json:"updatedAt"`
}
