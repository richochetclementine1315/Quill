package routes

import (
	"QuillBackend/controller"
	"QuillBackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Health check endpoint (lightweight, no DB query)
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Static files - must be BEFORE auth middleware so images are publicly accessible
	app.Static("/api/uploads", "./uploads")

	// Public routes (no authentication required)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/allpost", controller.AllPosts)       // Public - anyone can view posts
	app.Get("/api/allpost/:id", controller.DetailPost) // Public - anyone can view post details

	// Protected routes (require authentication)
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controller.CreatePost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/uniquepost", controller.UniquePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/upload-image", controller.Upload)
}
