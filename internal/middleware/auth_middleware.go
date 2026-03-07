package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"github.com/lovelymod/task-management-backend/utils"
)

func AuthMiddleware(config *bootstrap.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		splitToken := strings.Split(bearerToken, " ")

		if len(splitToken) != 2 || splitToken[1] == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entity.Response{
				Message:   entity.ErrAuthAccessTokenNotProvided.Error(),
				IsSuccess: false,
			})
			return
		}

		if splitToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entity.Response{
				Message:   entity.ErrAuthAccessTokenInvalid.Error(),
				IsSuccess: false,
			})
			return
		}

		claims, err := utils.ParseAccessToken(splitToken[1], config.ACCESS_TOKEN_SECRET)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entity.Response{
				Message:   err.Error(),
				IsSuccess: false,
			})
			return
		}

		c.Set("userId", claims.Subject)
		c.Next()
	}
}
