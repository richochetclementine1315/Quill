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

	// Protected routes (require authentication)
	// Apply middleware only to routes defined after this line
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allposts", controller.AllPosts)
	app.Get("/api/allposts/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/uniquepost", controller.UniquePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
}
