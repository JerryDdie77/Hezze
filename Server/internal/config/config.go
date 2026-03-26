package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBPort     string
	DBSSLMode  string

	JWTSecret  string
}

func GetConfig() (*Config, error) {
	_ = godotenv.Load()

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return &Config{}, fmt.Errorf("DB_NAME is required")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return &Config{}, fmt.Errorf("DB_HOST is required")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return &Config{}, fmt.Errorf("DB_USER is required")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return &Config{}, fmt.Errorf("DB_PASSWORD is required")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return &Config{}, fmt.Errorf("DB_PORT is required")
	}

	dbSSLMode := os.Getenv("DB_SSLMODE")
	if dbSSLMode == "" {
		return &Config{}, fmt.Errorf("DB_SSLMODE is required")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return &Config{}, fmt.Errorf("JWT_SECRET is required")
	}

	return &Config{
		DBName:     dbName,
		DBHost:     dbHost,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBPort:     dbPort,
		DBSSLMode:  dbSSLMode,
		JWTSecret:  jwtSecret,
	}, nil

}
