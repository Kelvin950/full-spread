package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func NewConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}
	return &Config{}
}

func (c Config) GetKey(key string) string {

	return os.Getenv(key)
}
