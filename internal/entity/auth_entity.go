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

type AuthRepository interface {
	Register(ctx context.Context, registerUser *User) (*User, error)
}
type AuthUsecase interface {
	Register(req *RegisterRequest) (*User, error)
}
type AuthHandler interface {
	Register(c *gin.Context)
}
