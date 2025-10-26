package database

import (
	"QuillBackend/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Capitalize the function to make it portable to other packages
func Connect() {
	// Load .env file only in development (optional in production)
	_ = godotenv.Load() // Ignore error if .env doesn't exist

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database!:%v", err)
	} else {
		log.Println("Connected Successfully")
	}

	// Configure connection pooling for better scalability
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to configure database pool: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)                 // Maximum 25 open connections (free tier MySQL limit)
	sqlDB.SetMaxIdleConns(5)                  // Keep 5 idle connections ready
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Recycle connections every 5 minutes
	sqlDB.SetConnMaxIdleTime(2 * time.Minute) // Close idle connections after 2 minutes

	log.Println("Database connection pool configured: MaxOpen=25, MaxIdle=5")

	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
