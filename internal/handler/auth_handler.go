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

	c.JSON(http.StatusCreated, entity.Response{
		Data:      user,
		Message:   "user_created",
		IsSuccess: true,
	})
}

func (h *authHandler) Login(c *gin.Context) {
	var req entity.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	tokens, err := h.authUsecase.Login(&req, c.ClientIP(), c.Request.UserAgent())
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	c.SetCookie("refreshToken", tokens.RefreshToken, 60*60*24*7, "/", "localhost", false, true)
	c.JSON(http.StatusOK, entity.Response{
		Data:      tokens.AccessToken,
		Message:   "logged_in",
		IsSuccess: true,
	})
}

func (h *authHandler) RefreshToken(c *gin.Context) {
	token, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.Response{
			Message:   "unauthorized",
			IsSuccess: false,
		})
		return
	}

	tokens, err := h.authUsecase.RefreshToken(token, c.ClientIP(), c.Request.UserAgent())
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	c.SetCookie("refreshToken", tokens.RefreshToken, 60*60*24*7, "/", "localhost", false, true)
	c.JSON(http.StatusOK, entity.Response{
		Data:      tokens.AccessToken,
		Message:   "refreshed",
		IsSuccess: true,
	})
}

func (h *authHandler) Logout(c *gin.Context) {
	token, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.Response{
			Message:   "unauthorized",
			IsSuccess: false,
		})
		return
	}

	if err := h.authUsecase.Logout(token); err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	c.SetCookie("refreshToken", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, entity.Response{
		Message:   "logged_out",
		IsSuccess: true,
	})
}
