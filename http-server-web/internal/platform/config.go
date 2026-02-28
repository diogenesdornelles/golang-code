package platform

import (
    "log"
    "os"

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
    AppPort string
    Mode    Mode
    // Add other configs here, e.g., DatabaseURL, StripeKey, etc.
}

// LoadConfig loads configurations from .env file or defaults
func LoadConfig() *Config {
    // Load .env file (ignore error if file doesn't exist)
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using defaults")
    }

    config := &Config{
        AppPort: getEnv("APP_PORT", ":7080"),
        Mode:    Mode(getEnv("APP_MODE", string(Dev))), // Default to "dev"
        // Add more: config.DatabaseURL = getEnv("DATABASE_URL", "default_value")
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