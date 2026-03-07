package entity

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
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

type RegisterRequest struct {
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) ([]User, error)
	GetRefreshToken(ctx context.Context, token string) (*RefreshToken, error)
	CreateUser(ctx context.Context, registerUser *User) (*User, error)
	CreateRefreshToken(ctx context.Context, refreshToken *RefreshToken) error
	RevokeRefreshToken(ctx context.Context, token string) error
}
type AuthUsecase interface {
	Register(req *RegisterRequest) (*User, error)
	Login(req *LoginRequest, clientIP string, userAgent string) (*LoginResponse, error)
	RefreshToken(token string, clientIP string, userAgent string) (*RefreshTokenResponse, error)
	Logout(token string) error
}
type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	Logout(c *gin.Context)
}
