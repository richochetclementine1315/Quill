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
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
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
	)
}
