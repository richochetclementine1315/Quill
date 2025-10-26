package main

import (
	"QuillBackend/database"
	"QuillBackend/routes"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file only in development (optional in production)
	_ = godotenv.Load() // Ignore error if .env doesn't exist

	database.Connect()
	port := os.Getenv("PORT")
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10MB max body size (for image uploads)
	})

	// Rate limiting middleware - prevent abuse
	app.Use(limiter.New(limiter.Config{
		Max:        100,             // 100 requests
		Expiration: 1 * time.Minute, // per minute per IP
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{
				"message": "Too many requests. Please try again later.",
			})
		},
	}))

	// Compression middleware - reduce response size by ~70%
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // Fast compression
	}))

	// CORS middleware - allow your Vercel frontend and enable CORS for static files
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://quill-ten.vercel.app",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		ExposeHeaders:    "Content-Length, Content-Type",
	}))

	routes.Setup(app)
	app.Listen(":" + port)

}
