package database

import (
	"QuillBackend/models"
	"log"
	"os"

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
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
