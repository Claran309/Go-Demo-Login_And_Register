package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// jwt
	JWTSecret      string
	JWTIssuer      string
	JWTExpireHours int

	// mysql
	DSN string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	return &Config{
		JWTSecret:      os.Getenv("JWT_SECRET"),
		JWTIssuer:      os.Getenv("JWT_ISSUER"),
		JWTExpireHours: getEnvInt("JWT_EXPIRATION_HOURS", 24),
		DSN:            os.Getenv("DSN"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}
