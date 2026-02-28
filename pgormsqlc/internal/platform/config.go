package platform

import (
	"log"
	"os"
	// "time"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// Mode represents the application mode (dev or prod)
type Mode string

// Constants for Mode enum
const (
	Dev  Mode = "dev"
	Prod Mode = "prod"
)

// Config holds all application configurations
type Config struct {
	AppPort     string
	Mode        Mode
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
	DatabaseURL string
}

// LoadConfig loads configurations from .env file or defaults
func LoadConfig() *Config {
	// Load .env file (ignore error if file doesn't exist)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	appPort := getEnv("APP_PORT", ":7080")
	appMode := Mode(getEnv("APP_MODE", string(Dev)))
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "1234")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "testdb")

	config := &Config{
		AppPort:     appPort,
		Mode:        appMode, // Default to "dev"
		DbHost:      dbHost,
		DbPort:      dbPort,
		DbUser:      dbUser,
		DbPassword:  dbPassword,
		DbName:      dbName,
		DatabaseURL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName),
	}

	return config
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ConnectDatabase(config *Config, ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
