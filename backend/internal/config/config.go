package config

import (
    "fmt"
    "os"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
    Port       string
}

func New() *Config {
    return &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "garage_user"),
        DBPassword: getEnv("DB_PASSWORD", "garage_password"),
        DBName:     getEnv("DB_NAME", "garage_management"),
        JWTSecret:  getEnv("JWT_SECRET", "your-secret-key-here"),
        Port:       getEnv("PORT", "8080"),
    }
}

func (c *Config) DatabaseURL() string {
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}