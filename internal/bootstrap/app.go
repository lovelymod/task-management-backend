package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Application struct {
	MongoDB *mongo.Client
	config  *Config
}

func AppInit() *Application {
	gin.SetMode(gin.ReleaseMode)

	config := configInit()
	mongoDB := mongoInit(config)

	application := Application{
		MongoDB: mongoDB,
		config:  config,
	}

	return &application
}
