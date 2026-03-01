package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"github.com/lovelymod/task-management-backend/internal/middleware"
)

type Handlers struct {
	AuthHandler entity.AuthHandler
}

func SetupRouter(r *gin.Engine, handlers *Handlers, config *bootstrap.Config) {

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.AuthHandler.Register)
			auth.POST("/login", handlers.AuthHandler.Login)
			auth.POST("/refresh-token", handlers.AuthHandler.RefreshToken)
		}
		private := api.Group("/")
		private.Use(middleware.AuthMiddleware(config))
		{
			private.POST("/auth/logout", handlers.AuthHandler.Logout)
		}

	}

}
