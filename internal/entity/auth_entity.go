package entity

import (
	"context"

	"github.com/gin-gonic/gin"
)

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
	Login(req *LoginRequest) (*LoginResponse, error)
	RefreshToken(token string) (*RefreshTokenResponse, error)
	Logout(token string) error
}
type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	Logout(c *gin.Context)
}
