package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"github.com/lovelymod/task-management-backend/utils"
)

func AuthMiddleware(config *bootstrap.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		at := c.GetHeader("Authorization")
		splitAT := strings.Split(at, " ")

		if len(splitAT) != 2 || splitAT[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entity.Response{
				Message:   errors.New("invalid_token").Error(),
				IsSuccess: false,
			})
			return
		}

		if splitAT[1] == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entity.Response{
				Message:   errors.New("token_not_provided").Error(),
				IsSuccess: false,
			})
			return
		}

		claims, err := utils.ParseAccessToken(splitAT[1], config.ACCESS_TOKEN_SECRET)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entity.Response{
				Message:   err.Error(),
				IsSuccess: false,
			})
			return
		}

		c.Set("userID", claims.Subject)
		c.Next()
	}
}
