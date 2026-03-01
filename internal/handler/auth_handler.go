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
		c.JSON(http.StatusBadRequest, entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	user, err := h.authUsecase.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Data:      user,
		Message:   "user_created",
		IsSuccess: true,
	})
}
