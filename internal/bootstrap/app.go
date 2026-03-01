package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Application struct {
	// MongoDB Client
	Client *mongo.Client
	// MongoDB Collections
	Mc *MongoCollections
	// ENV Configuration
	Config *Config
}

func AppInit() *Application {
	gin.SetMode(gin.ReleaseMode)

	config := configInit()
	client, mc := mongoInit(config)

	app := Application{
		Client: client,
		Mc:     mc,
		Config: config,
	}

	return &app
}
