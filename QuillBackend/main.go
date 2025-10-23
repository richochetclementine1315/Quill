package main

import (
	"QuillBackend/database"
	"QuillBackend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file only in development (optional in production)
	_ = godotenv.Load() // Ignore error if .env doesn't exist

	database.Connect()
	port := os.Getenv("PORT")
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://quill-ten.vercel.app",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Cookie",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		ExposeHeaders:    "Set-Cookie",
	}))

	// Serve static files from uploads directory
	app.Static("/uploads", "./uploads")

	routes.Setup(app)
	app.Listen(":" + port)

}
