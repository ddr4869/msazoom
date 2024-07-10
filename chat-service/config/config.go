package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DbConfig
	Gin GinConfig
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

type GinConfig struct {
	Mode string
}

var (
	JwtSecretPassword string
	Issuer            = "msazoom"
)

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GinConfig := GinConfig{
		Mode: os.Getenv("GIN_MODE"),
	}

	DbConfig := DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	JwtSecretPassword = os.Getenv("JWT_SECRET_PASSWORD")

	return &Config{
		Gin: GinConfig,
		DB:  DbConfig,
	}
}
