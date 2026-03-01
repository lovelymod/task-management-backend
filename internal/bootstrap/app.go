package bootstrap

import (
	"github.com/gin-gonic/gin"
)

type Application struct {
	// MongoDB Collections
	Mc *MongoCollections
	// ENV Configuration
	Config *Config
}

func AppInit() *Application {
	gin.SetMode(gin.ReleaseMode)

	config := configInit()
	mc := mongoInit(config)

	app := Application{
		Mc:     mc,
		Config: config,
	}

	return &app
}
