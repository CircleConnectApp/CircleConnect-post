package config

import (
	"os"
)

type Config struct {
	MongoURI    string
	DBName      string
	Environment string
	JWTSecret   string
}

func LoadConfig() Config {
	return Config{
		MongoURI:    getEnvOrDefault("MONGO_URI", "mongodb://localhost:27017"),
		DBName:      getEnvOrDefault("DB_NAME", "circle_connect_posts"),
		Environment: getEnvOrDefault("ENVIRONMENT", "development"),
		JWTSecret:   getEnvOrDefault("JWT_SECRET", "your-secret-key"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
