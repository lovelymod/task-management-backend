package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type RefreshToken struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Token     string        `json:"token" bson:"token"`
	TokenID   bson.ObjectID `json:"tokenId" bson:"tokenId,omitempty"`
	UserID    bson.ObjectID `json:"userId" bson:"userId,omitempty"`
	User      *User         `json:"user,omitempty" bson:"user,omitempty"`
	ClientIP  string        `json:"clientIp" bson:"clientIp"`
	UserAgent string        `json:"userAgent" bson:"userAgent"`
	IsRevoked bool          `json:"isRevoked" bson:"isRevoked"`
	ExpiresAt time.Time     `json:"expiresAt" bson:"expiresAt"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}
