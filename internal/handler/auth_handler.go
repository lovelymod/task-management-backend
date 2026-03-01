package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lovelymod/task-management-backend/internal/entity"
)

type authHandler struct {
	authUsecase entity.AuthUsecase
}

func NewAuthHandler(authUsecase entity.AuthUsecase) entity.AuthHandler {
	return &authHandler{authUsecase: authUsecase}
}

func (h *authHandler) Register(c *gin.Context) {
	var req entity.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authUsecase.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
