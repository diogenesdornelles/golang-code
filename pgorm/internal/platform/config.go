package platform

// import (
// 	"log"
// 	"os"
// 	// "time"

// 	"fmt"

// 	"github.com/joho/godotenv"

// )

// // Mode represents the application mode (dev or prod)
// type Mode string

// // Constants for Mode enum
// const (
// 	Dev  Mode = "dev"
// 	Prod Mode = "prod"
// )

// // Config holds all application configurations
// type Config struct {
// 	AppPort    string
// 	Mode       Mode
// 	DbHost     string
// 	DbPort     string
// 	DbUser     string
// 	DbPassword string
// 	DbName     string
// }

// // LoadConfig loads configurations from .env file or defaults
// func LoadConfig() *Config {
// 	// Load .env file (ignore error if file doesn't exist)
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found, using defaults")
// 	}

// 	config := &Config{
// 		AppPort:    getEnv("APP_PORT", ":7080"),
// 		Mode:       Mode(getEnv("APP_MODE", string(Dev))), // Default to "dev"
// 		DbHost:     getEnv("DB_HOST", "localhost"),
// 		DbPort:     getEnv("DB_PORT", "5432"),
// 		DbUser:     getEnv("DB_USER", "postgres"),
// 		DbPassword: getEnv("DB_PASSWORD", "1234"),
// 		DbName:     getEnv("DB_NAME", "testdb"),
// 	}

// 	return config
// }

// // getEnv gets an environment variable or returns a default value
// func getEnv(key, defaultValue string) string {
// 	if value := os.Getenv(key); value != "" {
// 		return value
// 	}
// 	return defaultValue
// }

// func ConnectDatabase(config *Config) (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// // https://gorm.io/docs/gorm_config.html
// // gorm.Config is a struct that holds configuration options for GORM Ver linha 68

// /*
// type Config struct {
// 	SkipDefaultTransaction                   bool // SkipDefaultTransaction disables default transactions for all operations
// 	NamingStrategy                           schema.Namer // NamingStrategy allows you to customize the naming strategy for tables, columns, etc.
// 	Logger                                   logger.Interface // Logger allows you to set a custom logger for GORM
// 	NowFunc                                  func() time.Time // NowFunc allows you to set a custom function to get the current time
// 	DryRun                                   bool // DryRun enables dry run mode, which generates SQL without executing it
// 	PrepareStmt                              bool // PrepareStmt enables prepared statement mode, which can improve performance for repeated queries
// 	DisableNestedTransaction                 bool // DisableNestedTransaction disables nested transactions, which can be useful for databases that don't support them
// 	AllowGlobalUpdate                        bool // AllowGlobalUpdate allows global updates without specifying a WHERE clause, use with caution
// 	DisableAutomaticPing                     bool // DisableAutomaticPing disables automatic pinging of the database when opening a connection
// 	DisableForeignKeyConstraintWhenMigrating bool // DisableForeignKeyConstraintWhenMigrating disables foreign key constraints when auto-migrating, which can be useful for databases that don't support them
// }
// */
