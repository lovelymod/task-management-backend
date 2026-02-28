package bootstrap

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MONGO_CONNECTION string

	HASH_COST string

	ACCESS_TOKEN_SECRET  string
	REFRESH_TOKEN_SECRET string
}

func configInit() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		MONGO_CONNECTION: os.Getenv("MONGO_CONNECTION"),

		HASH_COST: os.Getenv("HASH_COST"),

		ACCESS_TOKEN_SECRET:  os.Getenv("ACCESS_TOKEN_SECRET"),
		REFRESH_TOKEN_SECRET: os.Getenv("REFRESH_TOKEN_SECRET"),
	}

	return &config
}
