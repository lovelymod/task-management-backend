package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lovelymod/task-management-backend/internal/entity"
)

type Handlers struct {
	AuthHandler entity.AuthHandler
}

func SetupRouter(r *gin.Engine, handlers *Handlers) {

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.AuthHandler.Register)
		}
	}

}
