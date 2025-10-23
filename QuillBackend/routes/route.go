package routes

import (
	"QuillBackend/controller"
	"QuillBackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Public routes (no authentication required)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/allposts", controller.AllPosts)       // Public - anyone can view posts
	app.Get("/api/allposts/:id", controller.DetailPost) // Public - anyone can view post details
	app.Static("/api/uploads", "./uploads")             // Public - serve uploaded images

	// Protected routes (require authentication)
	// Apply middleware only to routes defined after this line
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controller.CreatePost)             // Protected - need login to create
	app.Put("/api/updatepost/:id", controller.UpdatePost)    // Protected - need login to update
	app.Get("/api/uniquepost", controller.UniquePost)        // Protected - get user's own posts
	app.Delete("/api/deletepost/:id", controller.DeletePost) // Protected - need login to delete
	app.Post("/api/upload-image", controller.Upload)         // Protected - need login to upload
}
